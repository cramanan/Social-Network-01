package database

import (
	"Social-Network-01/api/models"
	"context"
	"database/sql"
	"encoding/json"
	"time"
)

func (store *SQLite3Store) GetGroup(ctx context.Context, groupName string) (group *models.Group, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	row := tx.QueryRowContext(ctx, `SELECT * FROM group WHERE name = ?`, groupName)
	group = new(models.Group)
	var users []byte
	err = row.Scan(&group.Name, &users)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(users, &group.UsersIds)
	if err != nil {
		return nil, err
	}

	return group, err
}

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
	newgroup.UsersIds = group.UsersIds
	newgroup.Timestamp = time.Now().UTC()

	_, err = store.ExecContext(ctx,
		`INSERT INTO groups VALUES (?,?,?,?)`,
		newgroup.Name,
		newgroup.Description,
		newgroup.UsersIds,
		newgroup.Timestamp)
	if err != nil {
		return
	}

	return newgroup, tx.Commit()
}
