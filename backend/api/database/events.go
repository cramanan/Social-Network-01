package database

import (
	"Social-Network-01/api/types"
	"context"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

func (store *SQLite3Store) GetEvents(ctx context.Context, groupId string, limit, offset int) (events []types.Event, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	SELECT id, group_id, title, description, date
	FROM events
	WHERE group_id = ?
	LIMIT ? OFFSET ?;`,
		groupId,
		limit, offset)
	if err != nil {
		return nil, err
	}

	//TODO: This pattern was seen a lot: refactor using generics
	for rows.Next() {
		var event types.Event
		err = rows.Scan(
			&event.Id,
			&event.GroupId,
			&event.Title,
			&event.Description,
			&event.Date,
		)
		if err != nil {
			log.Println(err)
			continue
		}

		events = append(events, event)
	}

	if events == nil {
		return make([]types.Event, 0), nil
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return events, err
}

func (store *SQLite3Store) CreateEvent(ctx context.Context, event types.Event) (err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	groupExists := false
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT id FROM groups WHERE id = ? 
	);`, event.GroupId).Scan(&groupExists)
	if err != nil {
		return err
	}

	if !groupExists {
		return fmt.Errorf("group does not exists")
	}

	rawId, err := uuid.NewV4()
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx,
		"INSERT INTO events (id, group_id, title, description, date) VALUES (?, ?, ?, ?, ?)",
		rawId.String(),
		event.GroupId,
		event.Title,
		event.Description,
		event.Date,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}
