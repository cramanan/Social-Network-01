package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"Social-Network-01/api/models"

	"github.com/gofrs/uuid"
)

// Create a new posts in the database.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `req` is the corresponding postRequest (see ./api/models/posts.go).
//
// This method return a Post (see ./api/models/posts.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) CreatePost(ctx context.Context, req models.Post) (err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	if req.GroupId == "" {
		req.GroupId = "00000000"
	}

	_, err = tx.ExecContext(ctx, "INSERT INTO posts VALUES(?, ?, ?, ?, ?);",
		id.String(),
		req.UserId,
		req.GroupId,
		req.Content,
		time.Now(),

		id.String(),
	)
	if err != nil {
		return err
	}

	for _, image := range req.Images {
		_, err = tx.ExecContext(ctx, `
		INSERT INTO posts_images VALUES(?, ?)`,
			id.String(),
			image,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// Retrieve a post from the database using its postId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `postId` is the corresponding id in the database and is usualy find in the request pathvalue.
//
// This method return a post (see ./api/models/posts.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetPost(ctx context.Context, postId string) (post *models.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	post = new(models.Post)

	err = tx.QueryRowContext(ctx, `
	SELECT p.*, u.nickname
	FROM posts p JOIN users u
	ON p.user_id = u.id
	WHERE p.id = ?;`, postId).Scan(
		&post.Id,
		&post.UserId,
		&post.GroupId,
		&post.Content,
		&post.Timestamp,
		&post.Username,
	)
	if err != nil {
		return nil, err
	}

	if post.Images == nil {
		post.Images = make([]string, 0)
	}

	return
}

func (store *SQLite3Store) GetGroupPosts(ctx context.Context, groupId string, limit, offset int) (posts []models.Post, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	rows, err := store.QueryContext(ctx, `
	SELECT p.*, u.nickname
	FROM posts p JOIN users u
	ON p.user_id = u.id
	WHERE group_id = ?
	ORDER BY timestamp DESC
	LIMIT ? OFFSET ?;`,
		groupId, limit, offset)
	if err != nil {
		return
	}

	for rows.Next() {
		post := models.Post{}
		err = rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Content, &post.Timestamp, &post.Username)
		if err != nil {
			log.Println(err)
			continue
		}

		post.Images = make([]string, 0) // TODO: restore image system
		posts = append(posts, post)
	}

	if posts == nil {
		posts = make([]models.Post, 0)
	}

	return posts, tx.Commit()
}

func (store *SQLite3Store) LikePost(ctx context.Context, userId, postId string) (err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	var exists bool
	err = store.QueryRowContext(ctx, `
	SELECT EXISTS(
		SELECT * 
		FROM likes_records 
		WHERE user_id = ? AND post_id = ?
	);`, userId, postId).Scan(&exists)
	if err != nil {
		return err
	}

	var query string
	if !exists {
		query = "INSERT INTO likes_records VALUES(?, ?);"
	} else {
		query = "DELETE FROM likes_records WHERE user_id = ? AND post_id = ?;"
	}

	_, err = store.ExecContext(ctx, query, userId, postId)
	if err != nil {
		return err
	}

	return tx.Commit()
}
