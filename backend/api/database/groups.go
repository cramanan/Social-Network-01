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

	row := tx.QueryRowContext(ctx, `SELECT * FROM groups WHERE id = ?`, groupId)
	group = new(types.Group)
	err = row.Scan(&group.Id, &group.Name, &group.Description, &group.Image, &group.Timestamp)
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

	_, err = tx.ExecContext(ctx,
		`INSERT INTO groups (id, name, description, timestamp) VALUES (?, ?, ?, ?);`,
		group.Id,
		group.Name,
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
	) AND SELECT NOT EXISTS (
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

func (store *SQLite3Store) InviteUserIntoGroup(ctx context.Context, userId, groupId string) (err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	exists := false
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 FROM groups WHERE id = ?
	) AND SELECT EXISTS (
		SELECT 1 FROM users WHERE id = ? 
	);`).Scan(&exists)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("group or user does not exists")
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO groups_record (group_id, user_id, accepted) VALUES (?, ?, FALSE);")
	if err != nil {
		return err
	}

	return tx.Commit()
}
