package database

import (
	"Social-Network-01/api/models"
	"context"
	"database/sql"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

func (store *SQLite3Store) RegisterUser(ctx context.Context, req *models.RegisterRequest) (user models.User, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	var exists bool
	err = tx.QueryRowContext(ctx, "SELECT EXISTS (SELECT 1 FROM users WHERE email = $1);", req.Email).Scan(&exists)
	if err != nil {
		return
	}

	if exists {
		return user, ErrConflict
	}

	tx, err = store.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	id, err := uuid.NewV4()
	if err != nil {
		return
	}

	crypt, err := bcrypt.GenerateFromPassword([]byte(req.Password), 11)
	if err != nil {
		return
	}

	user.Id = id
	user.Nickname = req.Nickname
	user.Email = req.Email
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.DateOfBirth, err = time.Parse("2006-05-01", req.DateOfBirth)
	user.Timestamp = time.Now().UTC()

	// Id          uuid.UUID
	// Nickname    string
	// Email       string
	// password    []byte
	// FirstName   string
	// LastName    string
	// DateOfBirth time.Time
	// ImagePath   *string
	// AboutMe     *string
	// Private     bool
	// Timestamp   time.Time

	_, err = tx.ExecContext(ctx, "INSERT INTO users VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
		user.Id,
		user.Nickname,
		user.Email,
		crypt,
		user.FirstName,
		user.LastName,
		user.DateOfBirth,
		user.Timestamp,
	)
	if err != nil {
		return
	}

	return user, tx.Commit()
}

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
