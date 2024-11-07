package database

import (
	"context"
	"database/sql"
	"fmt"

	"Social-Network-01/api/types"
)

func (store *SQLite3Store) StoreChat(ctx context.Context, chat types.ServerChat) (err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `INSERT INTO chats VALUES(?, ?, ?, ?)`,
		chat.SenderId,
		chat.RecipientId,
		chat.Content,
		chat.Timestamp,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// Retrieve all chats beetween 2 users from the database using their userIds.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `user1Id` and `user2Id` are the corresponding users in the database and are usualy find in the
// request pathvalue and in the sessions field of the API structure. `limit` and `offset` can be retrieve with the parseRequestLimitAndOffset
// method using the request.
//
// This method return an array of chat (see ./api/types/chat.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetChats(ctx context.Context, user1Id, user2Id string, limit, offset int) (chats []types.ServerChat, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx,
		`SELECT * 
		FROM chats 
		WHERE (sender_id = ? AND recipient_id = ?)
		OR (recipient_id = ? AND sender_id = ?)
		ORDER BY timestamp 
		LIMIT ? OFFSET ? ;`, user1Id, user2Id, user1Id, user2Id, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		chat := types.ServerChat{}
		err := rows.Scan(&chat.SenderId, &chat.RecipientId, &chat.Content, &chat.Timestamp)
		if err != nil {
			fmt.Println(err)
			continue
		}

		chats = append(chats, chat)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	if chats == nil {
		chats = make([]types.ServerChat, 0)
	}

	return chats, nil
}
