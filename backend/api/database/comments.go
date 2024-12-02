package database

import (
	"context"
	"database/sql"
	"fmt"

	"Social-Network-01/api/types"
)

// GetComments retrieves all comments for a specific post from the database.
// - `postId`: The ID of the post for which comments are being retrieved.
// - `limit`: The maximum number of results to return (for pagination).
// - `offset`: The number of results to skip (for pagination).
// This method returns a slice of Comment objects or an error.
func (store *SQLite3Store) GetComments(ctx context.Context, postId string, limit, offset int) (comments []types.Comment, err error) {
    // Begin a read-only transaction.
    tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
    if err != nil {
        return nil, err
    }
    // Ensure the transaction is rolled back in case of an error.
    defer tx.Rollback()

    // Execute the SQL query to retrieve comments for the specified post.
    // Joins the `comments` table with the `users` table to include user details.
    rows, err := tx.QueryContext(ctx, `
	SELECT 
		c.*, 
		u.id, 
		u.nickname, 
		u.image_path
	FROM comments c JOIN users u
	ON c.user_id = u.id
	WHERE post_id = ?
	ORDER BY timestamp DESC
	LIMIT ? OFFSET ?;`,
        postId, limit, offset)
    if err != nil {
        return nil, err
    }
    // Ensure the rows are closed after processing.
    defer rows.Close()

    // Iterate through the result set and populate the comments slice.
    for rows.Next() {
        comment := types.Comment{}
        // Map the result set to the Comment struct, including user details.
        err := rows.Scan(
            &comment.UserId,
            &comment.PostId,
            &comment.Content,
            &comment.Image,
            &comment.Timestamp,
            &comment.UserId,      // User ID from the joined table.
            &comment.Username,    // User nickname.
            &comment.UserImage,   // User profile image path.
        )
        if err != nil {
            fmt.Println(err) // Log the error and skip to the next row.
            continue
        }

        comments = append(comments, comment)
    }

    // If no comments are found, return an empty slice.
    if comments == nil {
        comments = make([]types.Comment, 0)
    }

    return comments, nil
}

// CreateComment inserts a new comment into the database.
// - `comment`: The Comment object containing the comment details (user ID, post ID, content, image path, and timestamp).
// This method begins a transaction, inserts the comment into the `comments` table, and commits the transaction.
func (store *SQLite3Store) CreateComment(ctx context.Context, comment *types.Comment) (err error) {
    // Start a new transaction.
    tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
    if err != nil {
        return err
    }
    // Ensure the transaction is rolled back if an error occurs.
    defer tx.Rollback()

    // Execute the SQL query to insert the new comment into the `comments` table.
    _, err = tx.ExecContext(ctx, "INSERT INTO comments (user_id, post_id, content, image_path, timestamp) VALUES (?, ?, ?, ?, ?);",
        comment.UserId,
        comment.PostId,
        comment.Content,
        comment.Image,
        comment.Timestamp,
    )
    if err != nil {
        return err
    }

    // Commit the transaction.
    return tx.Commit()
}

