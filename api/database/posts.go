package database

import (
	"Social-Network-01/api/models"
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/gofrs/uuid"
)

func (store *SQLite3Store) CreatePost(ctx context.Context, req models.PostRequest) (group *models.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	group = new(models.Post)

	id, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	marshalCategories, err := json.Marshal(req.Categories)
	if err != nil {
		return nil, err
	}

	marshalImages, err := json.Marshal(req.Images)
	if err != nil {
		return nil, err
	}

	tx.ExecContext(ctx,
		`INSERT INTO posts VALUES(?, ?, ?, ?, ?, ?, ?);`,
		id.String(),
		req.UserId,
		req.GroupId,
		marshalCategories,
		req.Content,
		marshalImages,
		time.Now(),
	)

	return group, tx.Commit()
}
