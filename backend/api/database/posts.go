package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"Social-Network-01/api/types"

	"github.com/gofrs/uuid"
)

// Create a new posts in the database.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `req` is the corresponding postRequest (see ./api/types/posts.go).
//
// This method return a Post (see ./api/types/posts.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) CreatePost(ctx context.Context, req types.Post) (err error) {
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

	_, err = tx.ExecContext(ctx, "INSERT INTO posts (id, user_id, group_id, content, timestamp) VALUES (?, ?, ?, ?, ?);",
		id.String(),
		req.UserId,
		req.GroupId,
		req.Content,
		time.Now(),
	)
	if err != nil {
		return err
	}

	stmt, err := tx.PrepareContext(ctx, "INSERT INTO posts_images (post_id, path) VALUES (?, ?);")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, image := range req.Images {
		_, err = stmt.Exec(id.String(), image)
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
// This method return a post (see ./api/types/posts.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetPost(ctx context.Context, postId string) (post *types.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	post = new(types.Post)

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

func (store *SQLite3Store) GetGroupPosts(ctx context.Context, groupId string, limit, offset int) (posts []types.Post, err error) {
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

	stmt, err := tx.PrepareContext(ctx, `
	SELECT path
	FROM posts_images
	WHERE post_id = ?;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post types.Post
		err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.GroupId,
			&post.Content,
			&post.Timestamp,
			&post.Username)
		if err != nil {
			log.Println(err)
			continue
		}

		images, err := stmt.QueryContext(ctx, post.Id)
		if err != nil {
			return nil, err
		}

		for images.Next() {
			var path string
			err = images.Scan(&path)
			if err != nil {
				return nil, err
			}
			post.Images = append(post.Images, path)
		}

		if post.Images == nil {
			post.Images = make([]string, 0)
		}

		posts = append(posts, post)
	}

	if posts == nil {
		posts = make([]types.Post, 0)
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

	var query string = "INSERT INTO likes_records (user_id, post_id) VALUES (?, ?);"
	if exists {
		query = "DELETE FROM likes_records WHERE user_id = ? AND post_id = ?;"
	}

	_, err = store.ExecContext(ctx, query, userId, postId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (store *SQLite3Store) GetUserPosts(ctx context.Context, userId string, limit, offset int) (posts []types.Post, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	rows, err := tx.Query(`
	SELECT p.*, u.nickname
	FROM posts p JOIN users u
	ON p.user_id = u.id
	WHERE user_id = ?
	LIMIT ? OFFSET ?;`,
		userId, limit, offset)
	if err != nil {
		return nil, err
	}

	stmt, err := tx.PrepareContext(ctx, `
	SELECT path
	FROM posts_images
	WHERE post_id = ?;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var post types.Post

		err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.GroupId,
			&post.Content,
			&post.Timestamp,
			&post.Username)
		if err != nil {
			log.Println(err)
			continue
		}

		images, err := stmt.QueryContext(ctx, post.Id)
		if err != nil {
			return nil, err
		}

		for images.Next() {
			var path string
			err = images.Scan(&path)
			if err != nil {
				return nil, err
			}
			post.Images = append(post.Images, path)
		}

		if post.Images == nil {
			post.Images = make([]string, 0)
		}

		posts = append(posts, post)
	}

	if posts == nil {
		posts = make([]types.Post, 0)
	}

	return posts, nil
}
