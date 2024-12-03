package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"Social-Network-01/api/types"
)

const groupIdLength = 8

// GetGroup retrieves a group from the database by its ID.
// - `groupId`: The ID of the group to retrieve.
// Returns the Group object or an SQL error.
func (store *SQLite3Store) GetGroup(ctx context.Context, groupId string) (group *types.Group, err error) {
	// Start a read-only transaction.
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	// Ensure transaction rollback on error.
	defer tx.Rollback()

	// Initialize a new Group object.
	group = new(types.Group)

	// Query to retrieve the group details.
	err = tx.QueryRowContext(ctx, `
	SELECT * 
	FROM groups 
	WHERE id = ?`, groupId).Scan(
		&group.Id,
		&group.Name,
		&group.Owner,
		&group.Description,
		&group.Image,
		&group.Timestamp)
	if err != nil {
		return nil, err
	}

	return group, err
}

// NewGroup creates a new group in the database.
// - `group`: The group object containing the new group details.
// Returns an error if the operation fails.
func (store *SQLite3Store) NewGroup(ctx context.Context, group *types.Group) (err error) {
	// Start a transaction.
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	// Ensure transaction rollback on error.
	defer tx.Rollback()

	// Check if a group with the same name already exists.
	var exists bool
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 FROM groups WHERE name = ?
	);`, group.Name).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return ErrConflict // Group name conflict.
	}

	// Generate a unique ID and timestamp for the new group.
	group.Id = generateB64(groupIdLength)
	group.Timestamp = time.Now().UTC()

	// Insert the new group into the database.
	_, err = tx.ExecContext(ctx, `
	INSERT INTO groups (id, name, owner, description, timestamp)
	VALUES (?, ?, ?, ?, ?);`,
		group.Id,
		group.Name,
		group.Owner,
		group.Description,
		group.Timestamp)
	if err != nil {
		return
	}

	return tx.Commit()
}

// GetGroups retrieves a list of groups from the database with pagination.
// - `limit`: Maximum number of groups to retrieve.
// - `offset`: Offset for the pagination.
// Returns a slice of Group objects or an SQL error.
func (store *SQLite3Store) GetGroups(ctx context.Context, limit, offset int) (groups []*types.Group, err error) {
	// Start a transaction.
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	// Ensure transaction rollback on error.
	defer tx.Rollback()

	// Query to retrieve the groups with pagination.
	rows, err := tx.QueryContext(ctx, `
	SELECT * FROM groups
	LIMIT ? OFFSET ?;`,
		limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows and populate the groups slice.
	for rows.Next() {
		group := new(types.Group)
		err = rows.Scan(
			&group.Id,
			&group.Name,
			&group.Owner,
			&group.Description,
			&group.Image,
			&group.Timestamp)
		if err != nil {
			log.Println(err)
			continue // Skip rows with errors.
		}
		groups = append(groups, group)
	}

	// Return an empty slice if no groups are found.
	if groups == nil {
		groups = make([]*types.Group, 0)
	}

	return
}

func (store *SQLite3Store) AllowGroupInvite(ctx context.Context, hostId, guestId, groupId string) (boolean bool, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 FROM groups_record
		WHERE group_id = ? AND user_id = ?
	) AND NOT EXISTS (
		SELECT 1 FROM groups_record 
		WHERE group_id = ? AND user_id = ?
	);`,
		groupId, hostId,
		groupId, guestId,
	).Scan(&boolean)

	if err != nil {
		return false, err
	}

	return boolean, tx.Commit()
}

func (store *SQLite3Store) AllowGroupRequest(ctx context.Context, groupId, userId string) (boolean bool, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	err = tx.QueryRowContext(ctx, `
	SELECT NOT EXISTS (
		SELECT 1 FROM groups_record 
		WHERE group_id = ? AND user_id = ?
	) AND NOT EXISTS (
		SELECT 1 FROM groups
		WHERE id = ? AND owner = ?
	);`,
		groupId, userId,
		groupId, userId,
	).Scan(&boolean)

	if err != nil {
		return false, err
	}

	return boolean, tx.Commit()
}

// UserJoinGroup allows a user to join a group (via request or invite).
// - `userId`: The ID of the user joining the group.
// - `groupId`: The ID of the group to join.
// - `isRequest`: Indicates if this is a request to join (as opposed to an invite).
// Returns an SQL error if the operation fails.
func (store *SQLite3Store) UserJoinGroup(ctx context.Context, userId, groupId string, isRequest bool) (err error) {
	// Start a transaction.
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	// Ensure transaction rollback on error.
	defer tx.Rollback()

	// Check if the group and user exist in the database.
	exists := false
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 FROM groups WHERE id = ?
	) AND EXISTS (
		SELECT 1 FROM users WHERE id = ? 
	);`, groupId, userId).Scan(&exists)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("group or user does not exist")
	}

	// Insert a record into groups_record for the user.
	_, err = tx.ExecContext(ctx, `
	INSERT INTO groups_record (group_id, user_id, is_request, accepted)
	VALUES (?, ?, ?, FALSE);`, groupId, userId, isRequest)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (store *SQLite3Store) GetGroupInvites(ctx context.Context, userId string) (groups []types.Group, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	SELECT g.id, g.name, g.image
	FROM groups_record gr JOIN groups g
	ON gr.group_id = g.id
	WHERE gr.user_id = ? AND gr.is_request = FALSE AND gr.accepted = FALSE
	;`, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var group types.Group
		err = rows.Scan(
			&group.Id,
			&group.Name,
			&group.Image,
		)
		if err != nil {
			return nil, err
		}

		groups = append(groups, group)

	}

	if groups == nil {
		return make([]types.Group, 0), nil
	}

	return
}

type groupRequest struct {
	GroupId    string `json:"groupId"`
	GroupName  string `json:"groupName"`
	GroupImage string `json:"groupImage"`
	UserId     string `json:"userId"`
	UserName   string `json:"userName"`
	UserImage  string `json:"userImage"`
}

func (store *SQLite3Store) GetGroupRequests(ctx context.Context, userId string) (requests []groupRequest, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	WITH owned_groups AS (
		SELECT gr.*, g.name, g.image
		FROM groups g JOIN groups_record gr
		ON g.id = gr.group_id
		WHERE g.owner = ?
	)

	SELECT og.group_id, og.name,og.image , u.id, u.nickname, u.image_path
	FROM owned_groups og JOIN users u
	ON og.user_id = u.id
	WHERE og.is_request = TRUE AND og.accepted = FALSE;
	`, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var request groupRequest
		err = rows.Scan(
			&request.GroupId,
			&request.GroupName,
			&request.GroupImage,
			&request.UserId,
			&request.UserName,
			&request.UserImage,
		)
		if err != nil {
			return nil, err
		}

		requests = append(requests, request)
	}

	if requests == nil {
		return make([]groupRequest, 0), nil
	}

	return
}

func (store *SQLite3Store) AcceptGroupInvite(ctx context.Context, userId, groupId string) error {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	exists := false
	err = tx.QueryRowContext(ctx, `SELECT EXISTS (
		SELECT 1 from groups WHERE id = ?
	) AND EXISTS (
		SELECT 1 FROM users WHERE id = ? 
	);`, groupId, userId).Scan(&exists)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
	UPDATE groups_record
	SET accepted = TRUE
	WHERE group_id = ? AND user_id = ?;
	`, groupId, userId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (store *SQLite3Store) DeclineGroupInvite(ctx context.Context, userId, groupId string) error {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	exists := false
	err = tx.QueryRowContext(ctx, `SELECT EXISTS (
		SELECT 1 from groups WHERE id = ?
	) AND EXISTS (
		SELECT 1 FROM users WHERE id = ? 
	);`, groupId, userId).Scan(&exists)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
	DELETE FROM groups_record
	WHERE group_id = ? AND user_id = ?;
	`, groupId, userId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// AllowGroupInvite checks if a host can invite a guest to a group.
// - `hostId`: The ID of the user issuing the invitation.
// - `guestId`: The ID of the user being invited.
// - `groupId`: The ID of the group.
// Returns a boolean indicating if the invitation is allowed and/or an SQL error.
func (store *SQLite3Store) UserInGroup(ctx context.Context, groupId, userId string) (inGroup bool, err error) {
	// Start a transaction.
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	// Ensure transaction rollback on error.
	defer tx.Rollback()

	// Query to check if the host is a member of the group and the guest is not.
	exists := false
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 FROM groups
		WHERE id = ?
	) AND EXISTS (
		SELECT 1 FROM users
		WHERE id = ?
	);`, groupId, userId).Scan(&exists)
	if err != nil {
		return false, err
	}
	if !exists {
		return false, fmt.Errorf("error: group or user does not exists")
	}

	// Query to check if the user is neither a member nor the group owner.
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 FROM groups
		WHERE owner = ?
	) OR EXISTS (
		SELECT 1 FROM groups_record
		WHERE group_id = ? AND user_id = ? AND accepted = TRUE
	);`, userId, groupId, userId).Scan(&inGroup)
	if err != nil {
		return false, err
	}
	return inGroup, tx.Commit()
}

func (store *SQLite3Store) GetGroupMembers(ctx context.Context, groupId string, limit, offset int) (users []types.User, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Ensure transaction rollback on error.
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	SELECT u.nickname
	FROM groups g JOIN users u
	ON g.owner = u.id
	WHERE g.id = ?

	UNION

	SELECT u.nickname
	FROM groups_record gr JOIN users u
	ON gr.user_id = u.id
	WHERE gr.group_id = ? AND gr.accepted = TRUE
	LIMIT ? OFFSET ?;
	
	`, groupId, groupId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user types.User
		err = rows.Scan(
			&user.Nickname,
		)
		if err != nil {
			log.Println(err)
			continue
		}

		users = append(users, user)
	}

	return users, tx.Commit()
}
