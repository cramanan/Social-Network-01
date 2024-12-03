package database

import (
	"context"
	"database/sql"
	"log"
	"time"

	"Social-Network-01/api/types"

	"github.com/gofrs/uuid"
)

// CreatePost inserts a new post into the database, along with associated images.
// - `post`: The post object containing the post details (e.g., user ID, group ID, content, images).
// Returns an error if the operation fails.
func (store *SQLite3Store) CreatePost(ctx context.Context, post *types.Post) (err error) {
	// Start a read-only transaction.
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Generate a unique ID for the post.
	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	// TODO: Verify that the group exists, or ensure it can be nil.

	// Insert the post into the database.
	_, err = tx.ExecContext(ctx, `
	INSERT INTO posts (id, user_id, group_id, content, timestamp) 
	VALUES (?, ?, ?, ?, ?);`,
		id.String(),
		post.UserId,
		post.GroupId,
		post.Content,
		time.Now(),
	)
	if err != nil {
		return err
	}

	// Prepare the statement for inserting images.
	stmt, err := tx.PrepareContext(ctx, `
	INSERT INTO posts_images (post_id, path) 
	VALUES (?, ?);`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Insert each image associated with the post.
	for _, image := range post.Images {
		_, err = stmt.Exec(id.String(), image)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

// GetPost retrieves a post by its ID.
// - `postId`: The ID of the post to retrieve.
// Returns the Post object or an SQL error.
func (store *SQLite3Store) GetPost(ctx context.Context, postId string) (post *types.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	post = new(types.Post)

	// Retrieve the post details, including the username of the author.
	err = tx.QueryRowContext(ctx, `
	SELECT p.*, u.nickname, u.image_path
	FROM posts p 
	JOIN users u ON p.user_id = u.id
	WHERE p.id = ?;`, postId).Scan(
		&post.Id,
		&post.UserId,
		&post.GroupId,
		&post.Content,
		&post.Timestamp,
		&post.Username,
		&post.UserImage,
	)
	if err != nil {
		return nil, err
	}

	// Ensure the Images field is initialized.
	if post.Images == nil {
		post.Images = make([]string, 0)
	}

	return
}

// GetGroupPosts retrieves posts for a specific group, with pagination.
// - `groupId`: The ID of the group.
// - `limit`: The maximum number of posts to retrieve.
// - `offset`: The offset for pagination.
// Returns a slice of Post objects or an SQL error.
func (store *SQLite3Store) GetGroupPosts(ctx context.Context, groupId *string, limit, offset int) (posts []types.Post, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	var args []any
	var query string

	if groupId == nil {
		query = `
	SELECT p.*, u.nickname, u.image_path
	FROM posts p JOIN users u
	ON p.user_id = u.id
	WHERE group_id IS NULL
	ORDER BY timestamp DESC
	LIMIT ? OFFSET ?;`
		args = []any{limit, offset}
	} else {
		query = `
	SELECT p.*, u.nickname, u.image_path
	FROM posts p 
	JOIN users u ON p.user_id = u.id
	WHERE group_id = ?
	ORDER BY timestamp DESC
	LIMIT ? OFFSET ?;`
		args = []any{*groupId, limit, offset}
	}

	rows, err := tx.QueryContext(ctx, query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	// Prepare a statement to retrieve images for posts.
	stmt, err := tx.PrepareContext(ctx, `
	SELECT path
	FROM posts_images
	WHERE post_id = ?;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Process each post.
	for rows.Next() {
		var post types.Post
		err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.GroupId,
			&post.Content,
			&post.Timestamp,
			&post.Username,
			&post.UserImage)
		if err != nil {
			log.Println(err)
			continue // Skip posts with errors.
		}

		// Retrieve associated images for the post.
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

	// Ensure posts slice is initialized.
	if posts == nil {
		posts = make([]types.Post, 0)
	}

	return posts, tx.Commit()
}

// LikePost toggles a like for a post by a user.
// - `userId`: The ID of the user liking/unliking the post.
// - `postId`: The ID of the post to like/unlike.
// Returns an error if the operation fails.
func (store *SQLite3Store) LikePost(ctx context.Context, userId, postId string) (err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	// Check if the like record already exists.
	var exists bool
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS(
		SELECT * 
		FROM likes_records 
		WHERE user_id = ? AND post_id = ?
	);`, userId, postId).Scan(&exists)
	if err != nil {
		return err
	}

	// Toggle the like: Insert if not exists, delete otherwise.
	query := "INSERT INTO likes_records (user_id, post_id) VALUES (?, ?);"
	if exists {
		query = "DELETE FROM likes_records WHERE user_id = ? AND post_id = ?;"
	}

	_, err = tx.ExecContext(ctx, query, userId, postId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// GetUserPosts retrieves posts created by a specific user, with pagination.
// - `userId`: The ID of the user.
// - `limit`: The maximum number of posts to retrieve.
// - `offset`: The offset for pagination.
// Returns a slice of Post objects or an SQL error.
func (store *SQLite3Store) GetUserPosts(ctx context.Context, userId string, limit, offset int) (posts []types.Post, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	// Query posts by the user.
	rows, err := tx.Query(`
	SELECT p.*, u.nickname
	FROM posts p 
	JOIN users u ON p.user_id = u.id
	WHERE user_id = ?
	LIMIT ? OFFSET ?;`,
		userId, limit, offset)
	if err != nil {
		return nil, err
	}

	// Prepare a statement to retrieve images for posts.
	stmt, err := tx.PrepareContext(ctx, `
	SELECT path
	FROM posts_images
	WHERE post_id = ?;`)
	if err != nil {
		return nil, err
	}

	// Process each post.
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
			continue // Skip posts with errors.
		}

		// Retrieve associated images for the post.
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

	// Ensure posts slice is initialized.
	if posts == nil {
		posts = make([]types.Post, 0)
	}

	return posts, nil
}
