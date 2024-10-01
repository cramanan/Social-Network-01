package database

import (
	"Social-Network-01/api/models"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
)

func (store *SQLite3Store) CreatePost(ctx context.Context, req *models.PostRequest) (err error) {
	tx, err := store.BeginTx(ctx, nil)
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
		INSERT INTO posts VALUES(?, ?, ?, ?, ?, ?);
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

func (store *SQLite3Store) GetPost(ctx context.Context, userId, postId string) (post *models.Post, err error) {
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
