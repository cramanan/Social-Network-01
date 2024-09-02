package database

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"Social-Network-01/api/models"

	_ "github.com/mattn/go-sqlite3"
)

const TransactionTimeout = 3 * time.Second

var ErrConflict = errors.New("Conflict")

type SQLite3Store struct{ *sql.DB }

func NewSQLite3Store() (*SQLite3Store, error) {
	db, err := sql.Open("sqlite3", "db/db.sqlite3")
	if err != nil {
		return nil, err
	}
	return &SQLite3Store{db}, nil
}

func (store *SQLite3Store) RegisterUser(ctx context.Context) (user models.User, err error) {
	return
}

func (store *SQLite3Store) LogUser(ctx context.Context) (user models.User, err error) {
	return
}

func (store *SQLite3Store) GetUsers(ctx context.Context, limit, offset *int) (users []models.User, err error) {
	return
}

func (store *SQLite3Store) GetUsersById(ctx context.Context, id string) (user *models.User, err error) {
	return
}

func (store *SQLite3Store) CreatePost(ctx context.Context, req *models.PostRequest) (post models.Post, err error) {
	return
}

func (store *SQLite3Store) GetPosts(ctx context.Context, limit, offset *int) (posts []models.Post, err error) {
	return
}

func (store *SQLite3Store) GetPostByID(ctx context.Context, id string) (post *models.Post, err error) {
	return
}

func (store *SQLite3Store) CreateComment(ctx context.Context, req *models.CommentRequest) (comment models.Comment, err error) {
	return
}

func (store *SQLite3Store) GetCommentsOfID(ctx context.Context, id string, limit, offset *int) (comments []models.Comment, err error) {
	return
}

func (store *SQLite3Store) GetCategory(ctx context.Context, name string, limit, offset *int) (posts []models.Post, err error) {
	return
}
