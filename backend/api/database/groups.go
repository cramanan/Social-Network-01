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

// Retrieve the group from the database using its name.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `groupName` is the corresponding name in the database and is usualy find in the request pathvalue.
//
// This method return a Group (see ./api/types/groups.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetGroup(ctx context.Context, groupId string) (group *types.Group, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	group = new(types.Group)
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

// Create a new group in the database.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `group` is the corresponding group values (name...).
//
// This method return a Group (see ./api/types/groups.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) NewGroup(ctx context.Context, group *types.Group) (err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	var exists bool
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 FROM groups WHERE name = ?
	);`, group.Name).Scan(&exists)
	if err != nil {
		return err
	}
	if exists {
		return ErrConflict
	}

	group.Id = generateB64(groupIdLength)
	group.Timestamp = time.Now().UTC()

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

func (store *SQLite3Store) GetGroups(ctx context.Context, limit, offset int) (groups []*types.Group, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	SELECT * FROM groups
	LIMIT ? OFFSET ?;`,
		limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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
			continue
		}

		groups = append(groups, group)
	}

	if groups == nil {
		groups = make([]*types.Group, 0)
	}

	return
}

func (store *SQLite3Store) UserInGroup(ctx context.Context, groupId, userId string) (inGroup bool, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

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

func (store *SQLite3Store) UserJoinGroup(ctx context.Context, userId, groupId string, isRequest bool) (err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

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
		return fmt.Errorf("group or user does not exists")
	}

	_, err = tx.ExecContext(ctx, `
	INSERT INTO groups_record (group_id, user_id, is_request, accepted)
	VALUES (?, ?, ?, FALSE);`, groupId, userId, isRequest)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (store *SQLite3Store) GetGroupInvites(ctx context.Context, userId string) (groupInvites []types.Group, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	SELECT g.id, g.name
	FROM groups_record gr JOIN groups g
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
		)
		if err != nil {
			return nil, err
		}

	}

	if groupInvites == nil {
		return make([]types.Group, 0), nil
	}

	return
}

func (store *SQLite3Store) GetGroupRequests(ctx context.Context, userId string) (groupInvites []string, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	WITH owned_groups AS (
		SELECT gr.user_id
		FROM groups g JOIN groups_record gr
		ON g.id = gr.group_id
		WHERE g.owner = ?
	)

	SELECT og.user_id
	FROM owned_groups og JOIN users u
	ON og.user_id = u.id;
	`, userId)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id string
		err = rows.Scan(
			&id,
		)
		if err != nil {
			return nil, err
		}

	}

	if groupInvites == nil {
		return make([]string, 0), nil
	}

	return
}
