package database

import (
	"context"
	"database/sql"
	"log"

	"Social-Network-01/api/types"
)

func (store *SQLite3Store) StoreChat(ctx context.Context, chat types.ServerChat) (err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `INSERT INTO chats (sender_id, recipient_id, content, timestamp) VALUES (?, ?, ?, ?);`,
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
func (store *SQLite3Store) GetChats(ctx context.Context, senderId, recipientId string, limit, offset int) (chats []types.ServerChat, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

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
	defer rows.Close()

	for rows.Next() {
		chat := types.ServerChat{}
		err := rows.Scan(&chat.SenderId, &chat.RecipientId, &chat.Content, &chat.Timestamp)
		if err != nil {
			log.Println(err)
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

func (store *SQLite3Store) GetChatsFromGroup(ctx context.Context, groupId string, limit, offset int) (chats []types.ServerChat, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

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
	defer rows.Close()

	for rows.Next() {
		var chat types.ServerChat
		err = rows.Scan(
			&chat.SenderId,
			&chat.RecipientId,
			&chat.Content,
			&chat.Timestamp,
		)
		if err != nil {
			continue
		}

		chats = append(chats, chat)
	}

	if chats == nil {
		chats = make([]types.ServerChat, 0)
	}

	return chats, tx.Commit()
}

func (store *SQLite3Store) StoreGroupChat(ctx context.Context, chat types.ServerChat) (err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

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

	err = tx.Commit()

	return
}
