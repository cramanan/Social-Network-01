package database

import (
	"Social-Network-01/api/types"
	"context"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

func (store *SQLite3Store) RegisterUserToEvent(ctx context.Context, userId, eventId string) (err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	rowsExists := false
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 
		FROM users 
		WHERE id = ?
	) AND EXISTS (
		SELECT 1 
		FROM events 
		WHERE id = ?
	);`, userId, eventId).Scan(&rowsExists)
	if err != nil {
		return err
	}

	if !rowsExists {
		return fmt.Errorf("user or event does not exist")
	}

	var alreadyGoing bool
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 FROM events_records 
		WHERE user_id = ? AND event_id = ?
		);`, userId, eventId).Scan(&alreadyGoing)
	if err != nil {
		return err
	}

	query := "INSERT INTO events_records (event_id, user_id) VALUES(?, ?);"
	if alreadyGoing {
		query = "DELETE FROM events_records WHERE event_id = ? AND user_id = ?;"
	}

	_, err = tx.ExecContext(ctx, query, eventId, userId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (store *SQLite3Store) GetEvents(ctx context.Context, userId, groupId string, limit, offset int) (events []types.Event, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	SELECT e.*, CASE WHEN er.event_id IS NOT NULL THEN TRUE ELSE FALSE END AS events_records
	FROM events e LEFT JOIN events_records er
	ON e.id = er.event_id AND er.user_id = ?
	WHERE e.group_id = ?
	LIMIT ? OFFSET ?;`,
		userId, groupId,
		limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var event types.Event
		err = rows.Scan(
			&event.Id,
			&event.GroupId,
			&event.Title,
			&event.Description,
			&event.Date,
			&event.Going,
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
