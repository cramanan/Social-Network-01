package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
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
func (store *SQLite3Store) CreatePost(ctx context.Context, req *models.PostRequest) (err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	id, err := uuid.NewV4()
	if err != nil {
		return err
	}

	marshalImages, err := json.Marshal(req.Images)
	if err != nil {
		return err
	}

	marshalSelectedUsers, err := json.Marshal(req.SelectedUsers)
	if err != nil {
		return err
	}

	if req.Status != models.ENUM_ALMOST_PRIVATE {
		marshalSelectedUsers = nil
	} else {
		var exists bool
		for _, userid := range req.SelectedUsers {
			err = tx.QueryRowContext(ctx, `SELECT EXISTS(
				SELECT 1 FROM users WHERE id = ?
			);`, userid).Scan(&exists)
			if err != nil {
				return err
			}

			if !exists {
				return fmt.Errorf("user with id:%s do not exist", userid)
			}
		}
	}

	_, err = tx.ExecContext(ctx, `
		INSERT INTO posts VALUES(?, ?, COALESCE(?, "Global"), ?, ?, ?);
		INSERT INTO posts_status VALUES(?, ?, ?);`,

		id.String(),
		req.UserId,
		req.GroupName,
		req.Content,
		marshalImages,
		time.Now(),

		id.String(),
		req.Status,
		marshalSelectedUsers,
	)
	if err != nil {
		return err
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

	var status_enum int
	var unmarshalImages, unmarshalUsers []byte

	err = tx.QueryRowContext(ctx, `
	SELECT 
		posts.*, ps.status_enum, ps.users
	FROM posts JOIN posts_status
	ON posts.id = posts_status.post_id
	WHERE id = ?;`, postId).Scan(
		&post.Id,
		&post.UserId,
		&post.GroupName,
		&post.Content,
		&unmarshalImages,
		&post.Timestamp,

		&status_enum,
		&unmarshalUsers,
	)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(unmarshalImages, &post.Images)
	if post.Images == nil {
		post.Images = make([]string, 0)
	}

	return
}
