package database

import (
	"context"
	"database/sql"
	"fmt"

	"Social-Network-01/api/types"
)

// Retrieve all comments of one post from the database using its postId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `postId` is the corresponding post in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be retrieve with the parseRequestLimitAndOffset method using the request.
//
// This method return an array of comment (see ./api/types/comments.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetComments(ctx context.Context, postId string, limit, offset int) (comments []types.Comment, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, "SELECT * FROM comments WHERE parent_id = ? ORDER BY timestamp DESC LIMIT ? OFFSET ? ;", postId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		comment := types.Comment{}
		err := rows.Scan(&comment.UserId, &comment.PostId, &comment.Content, &comment.Image, &comment.Timestamp)
		if err != nil {
			fmt.Println(err)
			continue
		}

		comments = append(comments, comment)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (store *SQLite3Store) CreateComment(ctx context.Context, comment types.Comment) (err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, "INSERT INTO comments VALUES(?, ?, ?, ?, ?);",
		comment.UserId,
		comment.PostId,
		comment.Content,
		comment.Image,
		comment.Timestamp,
	)
	if err != nil {
		return err
	}

	return tx.Commit()
}
