package database

import (
	"context"
	"database/sql"
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
	err = row.Scan(&group.Id, &group.Name, &group.Description, &group.Timestamp)
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
func (store *SQLite3Store) NewGroup(ctx context.Context, group *types.Group) (newgroup *types.Group, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	var exists bool
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 FROM groups WHERE name = $1
	);`, group.Name).Scan(&exists)
	if err != nil {
		return
	}
	if exists {
		return nil, ErrConflict
	}

	newgroup = new(types.Group)
	newgroup.Id = generateB64(groupIdLength)
	newgroup.Name = group.Name
	newgroup.Description = group.Description
	newgroup.Timestamp = time.Now().UTC()

	_, err = store.ExecContext(ctx,
		`INSERT INTO groups VALUES (?,?,?,?)`,
		newgroup.Id,
		newgroup.Name,
		newgroup.Description,
		newgroup.Timestamp)
	if err != nil {
		return
	}

	return newgroup, tx.Commit()
}

func (store *SQLite3Store) GetGroups(ctx context.Context, limit, offset int) (groups []*types.Group, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	rows, err := store.QueryContext(ctx, `
	SELECT * FROM groups
	LIMIT ? OFFSET ?;`,
		limit, offset)

	for rows.Next() {
		group := new(types.Group)
		err = rows.Scan(&group.Id, &group.Name, &group.Description, &group.Timestamp)
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
