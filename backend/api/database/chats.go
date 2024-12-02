package database

import (
	"context"
	"database/sql"
	"log"

	"Social-Network-01/api/types"
)

// StoreChat stores a single chat message in the database.
// - `ctx`: The context of the operation (can be used for timeouts, cancellation, etc.).
// - `chat`: The ServerChat object containing the chat details (sender ID, recipient ID, content, and timestamp).
// This method begins a transaction, inserts the chat data into the `chats` table, and commits the transaction.
func (store *SQLite3Store) StoreChat(ctx context.Context, chat types.ServerChat) (err error) {
    // Start a new transaction in read-write mode.
    tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
    if err != nil {
        return err
    }
    // Ensure the transaction is rolled back if an error occurs.
    defer tx.Rollback()

    // Execute the SQL query to insert the chat into the `chats` table.
    _, err = tx.ExecContext(ctx, `INSERT INTO chats (sender_id, recipient_id, content, timestamp) VALUES (?, ?, ?, ?);`,
        chat.SenderId,
        chat.RecipientId,
        chat.Content,
        chat.Timestamp,
    )
    if err != nil {
        return err
    }

    // Commit the transaction.
    return tx.Commit()
}

// GetChats retrieves all chat messages exchanged between two users.
// - `senderId`: The ID of the sender.
// - `recipientId`: The ID of the recipient.
// - `limit`: The maximum number of results to return.
// - `offset`: The number of results to skip for pagination.
// This method returns a slice of ServerChat objects or an error.
func (store *SQLite3Store) GetChats(ctx context.Context, senderId, recipientId string, limit, offset int) (chats []types.ServerChat, err error) {
    // Begin a read-only transaction.
    tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
    if err != nil {
        return nil, err
    }
    // Ensure the transaction is rolled back in case of an error.
    defer tx.Rollback()

    // Execute the SQL query to retrieve chat messages between the specified users.
    rows, err := tx.QueryContext(ctx, `
	SELECT *
	FROM chats
	WHERE (sender_id = ? AND recipient_id = ?)
	OR (recipient_id = ? AND sender_id = ?)
	ORDER BY timestamp
	LIMIT ? OFFSET ?;`,
        senderId, recipientId,
        senderId, recipientId,
        limit, offset)
    if err != nil {
        return nil, err
    }
    // Ensure the rows are closed after processing.
    defer rows.Close()

    // Iterate through the result set and populate the chats slice.
    for rows.Next() {
        chat := types.ServerChat{}
        err := rows.Scan(&chat.SenderId, &chat.RecipientId, &chat.Content, &chat.Timestamp)
        if err != nil {
            log.Println(err) // Log the error and continue with the next row.
            continue
        }

        chats = append(chats, chat)
    }

    // Check for errors that occurred during iteration.
    err = rows.Err()
    if err != nil {
        return nil, err
    }

    // If no chats were found, return an empty slice.
    if chats == nil {
        chats = make([]types.ServerChat, 0)
    }

    return chats, nil
}

// GetChatsFromGroup retrieves chat messages from a group by group ID.
// - `groupId`: The ID of the group.
// - `limit`: The maximum number of results to return.
// - `offset`: The number of results to skip for pagination.
// This method returns a slice of ServerChat objects or an error.
func (store *SQLite3Store) GetChatsFromGroup(ctx context.Context, groupId string, limit, offset int) (chats []types.ServerChat, err error) {
    // Begin a read-only transaction.
    tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
    if err != nil {
        return nil, err
    }
    // Ensure the transaction is rolled back in case of an error.
    defer tx.Rollback()

    // Execute the SQL query to retrieve chat messages from the specified group.
    rows, err := tx.QueryContext(ctx, `
	SELECT * 
	FROM group_chats 
	WHERE group_id = ?
	LIMIT ? OFFSET ?;`,
        groupId,
        limit, offset)
    if err != nil {
        return nil, err
    }
    // Ensure the rows are closed after processing.
    defer rows.Close()

    // Iterate through the result set and populate the chats slice.
    for rows.Next() {
        var chat types.ServerChat
        err = rows.Scan(
            &chat.SenderId,
            &chat.RecipientId,
            &chat.Content,
            &chat.Timestamp,
        )
        if err != nil {
            continue // Skip rows with errors.
        }

        chats = append(chats, chat)
    }

    // If no chats were found, return an empty slice.
    if chats == nil {
        chats = make([]types.ServerChat, 0)
    }

    // Commit the transaction.
    return chats, tx.Commit()
}

// StoreGroupChat stores a chat message in a group.
// - `chat`: The ServerChat object containing the chat details (sender ID, group ID, content, and timestamp).
// This method begins a transaction, inserts the chat data into the `group_chats` table, and commits the transaction.
func (store *SQLite3Store) StoreGroupChat(ctx context.Context, chat types.ServerChat) (err error) {
    // Start a new transaction.
    tx, err := store.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    // Ensure the transaction is rolled back if an error occurs.
    defer tx.Rollback()

    // Execute the SQL query to insert the chat into the `group_chats` table.
    _, err = tx.ExecContext(ctx,
        "INSERT INTO group_chats (sender_id, group_id, content, timestamp) VALUES (?, ?, ?, ?)",
        chat.SenderId,
        chat.RecipientId,
        chat.Content,
        chat.Timestamp,
    )
    if err != nil {
        return err
    }

    // Commit the transaction.
    err = tx.Commit()

    return
}
