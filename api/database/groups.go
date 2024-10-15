package database

import (
	"context"
	"database/sql"
	"time"

	"Social-Network-01/api/models"
)

// Retrieve the group from the database using its name.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `groupName` is the corresponding name in the database and is usualy find in the request pathvalue.
//
// This method return a Group (see ./api/models/groups.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetGroup(ctx context.Context, groupName string) (group *models.Group, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(ctx, `SELECT * FROM groups WHERE name = ?`, groupName)
	group = new(models.Group)
	err = row.Scan(&group.Name, &group.Description, &group.Timestamp)
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
// This method return a Group (see ./api/models/groups.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) NewGroup(ctx context.Context, group *models.Group) (newgroup models.Group, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	var exists bool
	err = tx.QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM groups WHERE name = $1);", group.Name).Scan(&exists)
	if err != nil {
		return
	}
	if exists {
		return *group, ErrConflict
	}

	newgroup.Name = group.Name
	newgroup.Description = group.Description
	newgroup.Timestamp = time.Now().UTC()

	_, err = store.ExecContext(ctx,
		`INSERT INTO groups VALUES (?,?,?,?)`,
		newgroup.Name,
		newgroup.Description,
		newgroup.Timestamp)
	if err != nil {
		return
	}

	return newgroup, tx.Commit()
}
