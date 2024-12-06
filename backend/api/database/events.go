package database

import (
	"Social-Network-01/api/types"
	"context"
	"fmt"
	"log"

	"github.com/gofrs/uuid"
)

// RegisterUserToEvent registers or unregisters a user for an event.
// - `userId`: The ID of the user.
// - `eventId`: The ID of the event.
// If the user is not already registered, they are added to the event. If they are already registered, they are removed.
func (store *SQLite3Store) RegisterUserToEvent(ctx context.Context, userId, eventId string) (err error) {
	// Begin a transaction.
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	// Ensure the transaction is rolled back if an error occurs.
	defer tx.Rollback()

	// Check if the user and event exist in the database.
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

	// Return an error if the user or event does not exist.
	if !rowsExists {
		return fmt.Errorf("user or event does not exist")
	}

	// Check if the user is already registered for the event.
	var alreadyGoing bool
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 FROM events_records 
		WHERE user_id = ? AND event_id = ?
		);`, userId, eventId).Scan(&alreadyGoing)
	if err != nil {
		return err
	}

	// Prepare the query: insert to register, delete to unregister.
	query := "INSERT INTO events_records (event_id, user_id) VALUES(?, ?);"
	if alreadyGoing {
		query = "DELETE FROM events_records WHERE event_id = ? AND user_id = ?;"
	}

	// Execute the query.
	_, err = tx.ExecContext(ctx, query, eventId, userId)
	if err != nil {
		return err
	}

	// Commit the transaction.
	return tx.Commit()
}

// GetEvents retrieves events for a specific group.
// - `userId`: The ID of the user (to check registration status).
// - `groupId`: The ID of the group for which events are being retrieved.
// - `limit`: The maximum number of results to return (for pagination).
// - `offset`: The number of results to skip (for pagination).
// This method returns a slice of Event objects or an error.
func (store *SQLite3Store) GetEvents(ctx context.Context, userId, groupId string, limit, offset int) (events []types.Event, err error) {
	// Begin a transaction.
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Ensure the transaction is rolled back in case of an error.
	defer tx.Rollback()

	// Execute the SQL query to retrieve events and check if the user is registered for each.
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
	// Ensure the rows are closed after processing.
	defer rows.Close()

	// Iterate through the result set and populate the events slice.
	for rows.Next() {
		var event types.Event
		err = rows.Scan(
			&event.Id,
			&event.GroupId,
			&event.Title,
			&event.Description,
			&event.Date,
			&event.Going, // Boolean indicating if the user is registered.
		)
		if err != nil {
			log.Println(err) // Log errors and continue with the next row.
			continue
		}

		events = append(events, event)
	}

	// Return an empty slice if no events are found.
	if events == nil {
		return make([]types.Event, 0), nil
	}

	// Commit the transaction.
	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return events, err
}

// CreateEvent creates a new event in the database.
// - `event`: The Event object containing the event details (group ID, title, description, and date).
// This method checks if the group exists before creating the event.
func (store *SQLite3Store) CreateEvent(ctx context.Context, event types.Event) (value *types.Event, err error) {
	// Begin a transaction.
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	// Ensure the transaction is rolled back if an error occurs.
	defer tx.Rollback()

	// Check if the group exists in the database.
	groupExists := false
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT id FROM groups WHERE id = ? 
	);`, event.GroupId).Scan(&groupExists)
	if err != nil {
		return nil, err
	}

	// Return an error if the group does not exist.
	if !groupExists {
		return nil, fmt.Errorf("group does not exist")
	}

	// Generate a new UUID for the event.
	rawId, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	event.Id = rawId.String()

	// Insert the new event into the database.
	_, err = tx.ExecContext(ctx,
		"INSERT INTO events (id, group_id, title, description, date) VALUES (?, ?, ?, ?, ?)",
		event.Id,
		event.GroupId,
		event.Title,
		event.Description,
		event.Date,
	)
	if err != nil {
		return nil, err
	}

	// Commit the transaction.
	return &event, tx.Commit()
}
