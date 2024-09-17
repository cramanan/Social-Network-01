package database

import (
	"Social-Network-01/api/models"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Perform the action of registering one user in the database.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `req` is a RegisterRequest (see ./api/models/users.go) and is create from a form data after posting.
//
// This method return a user (see ./api/models/users.go) or usualy an SQL error (one is nil when the other isn't).
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

	user.Id = id.String()
	user.Nickname = req.Nickname
	user.Email = req.Email
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.DateOfBirth, err = time.Parse("2006-05-01", req.DateOfBirth)
	user.Timestamp = time.Now().UTC()

	_, err = tx.ExecContext(ctx, "INSERT INTO users VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);",
		user.Id,
		user.Nickname,
		user.Email,
		crypt,
		user.FirstName,
		user.LastName,
		user.DateOfBirth,
		nil,
		nil,
		false,
		user.Timestamp,
	)
	if err != nil {
		return
	}

	return user, tx.Commit()
}

// Perform the action of logging one user.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `req` is a LoginRequest (see ./api/models/users.go) and is create from a form data after posting.
//
// This method return a user (see ./api/models/users.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) LogUser(ctx context.Context, req *models.LoginRequest) (user models.User, err error) {
	row := store.QueryRowContext(ctx, "SELECT * FROM users WHERE email = ?;", req.Email)
	comp := []byte{}

	err = row.Scan(
		&user.Id,
		&user.Nickname,
		&user.Email,
		&comp,
		&user.FirstName,
		&user.LastName,
		&user.DateOfBirth,
		&user.ImagePath,
		&user.AboutMe,
		&user.Private,
		&user.Timestamp,
	)
	if err != nil {
		return
	}

	return user, bcrypt.CompareHashAndPassword(comp, []byte(req.Password))
}

// Retrieve all user datas of one user from the database using its userId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding user in the database and is usualy find in the request pathvalue or
// in the sessions field of the API structure.
//
// This method return a user (see ./api/models/users.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetUser(ctx context.Context, userId string) (user *models.User, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	user = new(models.User)
	err = store.QueryRowContext(ctx, `SELECT 
		id,
		nickname,
		email,
		first_name,
		last_name,
		date_of_birth,
		image_path,
		about_me,
		private,
		timestamp

		FROM users WHERE id = ?;`, userId).Scan(

		&user.Id,
		&user.Nickname,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.DateOfBirth,
		&user.ImagePath,
		&user.AboutMe,
		&user.Private,
		&user.Timestamp,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (store *SQLite3Store) DeleteUser(ctx context.Context, userId string) error {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("DELETE FROM users WHERE id = ?;", userId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (store *SQLite3Store) SortUsers(ctx context.Context, id string) (users []models.User, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true}) // Begin SQL Transaction (readonly)
	if err != nil {
		log.Println(err)
		return
	}
	defer tx.Rollback()
	rows, err := tx.QueryContext(ctx, `
	WITH contacted AS (
		SELECT DISTINCT m.recipientid, u.name
		FROM messages m JOIN users u
		ON m.recipientid = u.id
		WHERE m.senderid = ?
		GROUP BY u.name
		ORDER BY m.created DESC, u.name
	), not_contacted AS (
		SELECT u.id, u.name
		FROM users u
		WHERE u.id NOT IN (SELECT recipientid FROM contacted)
		AND u.id != ?
		ORDER BY u.name ASC
	)

	SELECT * FROM contacted
	UNION ALL
	SELECT * FROM not_contacted;`, id, id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.Id, &user.Nickname)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}

	return users, tx.Commit()
}
