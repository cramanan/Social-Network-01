package database

import (
	"Social-Network-01/api/models"
	"context"
	"database/sql"

	"github.com/gofrs/uuid"
)

func (store *SQLite3Store) GetUserPosts(ctx context.Context, userId uuid.UUID) (posts []models.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	rows, err := tx.QueryContext(ctx, "SELECT * FROM posts WHERE id = ?;", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan rows into posts
	for rows.Next() {
		var post models.Post

		err = rows.Scan(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (store *SQLite3Store) GetGroupPosts(ctx context.Context, groupId uuid.UUID) (posts []models.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Commit()

	rows, err := tx.QueryContext(ctx, "SELECT * FROM posts WHERE group_id = ?;", groupId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan rows into posts
	for rows.Next() {
		var post models.Post

		err = rows.Scan(&post)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
