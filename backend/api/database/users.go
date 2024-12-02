package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"Social-Network-01/api/types"

	"github.com/gofrs/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Perform the action of registering one user in the database.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `req` is a RegisterRequest (see ./api/types/users.go) and is create from a form data after posting.
//
// This method return a user (see ./api/types/users.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) RegisterUser(ctx context.Context, req *types.RegisterRequest) (user types.User, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer tx.Rollback()

	var exists bool
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 FROM users WHERE email = ?
	);`, req.Email).Scan(&exists)
	if err != nil {
		return
	}

	if exists {
		return user, ErrConflict
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

	_, err = tx.ExecContext(ctx, `
		INSERT INTO users (id, nickname, email, password, first_name, last_name, date_of_birth, image_path, about_me, is_private, timestamp) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`,
		user.Id,
		user.Nickname,
		user.Email,
		crypt,
		user.FirstName,
		user.LastName,
		user.DateOfBirth,
		"https://upload.wikimedia.org/wikipedia/commons/2/2c/Default_pfp.svg",
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
// `ctx` is the context of the request. `req` is a LoginRequest (see ./api/types/users.go) and is create from a form data after posting.
//
// This method return a user (see ./api/types/users.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) LogUser(ctx context.Context, req *types.LoginRequest) (user types.User, err error) {
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
		&user.IsPrivate,
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
// This method return a user (see ./api/types/users.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetUser(ctx context.Context, userId string) (user *types.User, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	user = new(types.User)
	err = tx.QueryRowContext(ctx, `SELECT 
		id,
		nickname,
		email,
		first_name,
		last_name,
		date_of_birth,
		image_path,
		about_me,
		is_private,
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
		&user.IsPrivate,
		&user.Timestamp,
	)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// Delete a user from the database using its userId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding user in the database and is usualy find in the request pathvalue or
// in the sessions field of the API structure.
//
// This method return an SQL error or nil if there are none.
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

// Sort users from the database on the order of chat log of another user.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `id` is the corresponding user in the database and is usualy find in the request pathvalue.
//
// This method return an array of users (see ./api/types/users.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetMessagedUsers(ctx context.Context, userId string, limit, offset int) (users []*types.User, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	WITH contacted AS (
		SELECT DISTINCT m.recipient_id, u.nickname, u.image_path
		FROM chats m JOIN users u
		ON m.recipient_id = u.id
		WHERE m.sender_id = ?
		GROUP BY u.nickname
		ORDER BY MAX(m.timestamp) DESC, u.nickname
	), not_contacted AS (
		SELECT u.id, u.nickname, u.image_path
		FROM users u
		WHERE u.id NOT IN (SELECT recipient_id FROM contacted)
		AND u.id != ?
		ORDER BY u.nickname ASC
	)

	SELECT * FROM contacted
	UNION ALL
	SELECT * FROM not_contacted
	LIMIT ? OFFSET ?;`,
		userId,
		userId,
		limit,
		offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := new(types.User)
		err = rows.Scan(
			&user.Id,
			&user.Nickname,
			&user.ImagePath,
		)
		if err != nil {
			log.Println(err)
			continue
		}

		users = append(users, user)
	}

	return users, nil
}

// Retrieve all follower of a user from the database using his userId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding user in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be retrieve with the parseRequestLimitAndOffset method using the request.
//
// This method return an array of user (see ./api/types/users.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetProfileFollowers(ctx context.Context, userId string, limit, offset int) (users []types.User, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	SELECT u.id, u.nickname, u.image_path
	FROM follow_records f 
	JOIN users u 
	ON f.follower_id = u.id
	WHERE user_id = ? AND accepted = TRUE
	LIMIT ? OFFSET ?;`,
		userId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := types.User{}
		err = rows.Scan(
			&user.Id,
			&user.Nickname,
			&user.ImagePath,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if users == nil {
		users = make([]types.User, 0)
	}

	return users, nil
}

func (store *SQLite3Store) GetProfileFollowing(ctx context.Context, userId string, limit, offset int) (users []types.User, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	SELECT u.id, u.nickname, u.image_path
	FROM follow_records f 
	JOIN users u 
	ON f.user_id = u.id
	WHERE follower_id = ? AND accepted = TRUE
	LIMIT ? OFFSET ?;`,
		userId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := types.User{}
		err = rows.Scan(
			&user.Id,
			&user.Nickname,
			&user.ImagePath,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if users == nil {
		users = make([]types.User, 0)
	}

	return users, nil
}

func (store *SQLite3Store) GetUserStats(ctx context.Context, userId string) (stats types.UserStats, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return
	}
	defer tx.Rollback()

	err = tx.QueryRowContext(ctx, `SELECT
		(SELECT COUNT(*) FROM follow_records WHERE user_id = ? AND accepted = TRUE) AS followers,
		(SELECT COUNT(*) FROM follow_records WHERE follower_id = ? AND accepted = TRUE) AS following,
		(SELECT COUNT(*) FROM posts WHERE user_id = ?) AS posts,
		(SELECT COUNT(*) FROM likes_records WHERE user_id = ?) AS likes;`,
		userId, userId, userId, userId).Scan(

		&stats.NumFollowers,
		&stats.NumFollowing,
		&stats.NumPosts,
		&stats.NumLikes)
	if err != nil {
		return
	}

	stats.Id = userId
	return stats, tx.Commit()
}

func (store *SQLite3Store) UpdateUser(ctx context.Context, id string, value types.User) (modified *types.User, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var original types.User
	// Retrieve the original user
	err = tx.QueryRow(`
	SELECT 
		id,
		nickname,
		email,
		first_name,
		last_name,
		date_of_birth,
		image_path,
		about_me,
		is_private,
		timestamp
	FROM users u
	WHERE id = ?;`, id).Scan(
		&original.Id,
		&original.Nickname,
		&original.Email,
		&original.FirstName,
		&original.LastName,
		&original.DateOfBirth,
		&original.ImagePath,
		&original.AboutMe,
		&original.IsPrivate,
		&original.Timestamp,
	)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user with id %s not found", id)
	}
	if err != nil {
		return nil, err
	}

	ValidString := func(value string) bool {
		return strings.TrimSpace(value) != ""
	}

	queryParts := make([]string, 0, 6)
	args := make([]any, 0, 6)

	// Collect updated fields
	if ValidString(value.Nickname) && value.Nickname != original.Nickname {
		queryParts = append(queryParts, "nickname = ?")
		args = append(args, value.Nickname)
		original.Nickname = value.Nickname
	}

	if ValidString(value.FirstName) && value.FirstName != original.FirstName {
		queryParts = append(queryParts, "first_name = ?")
		args = append(args, value.FirstName)
		original.FirstName = value.FirstName

	}

	if ValidString(value.LastName) && value.LastName != original.LastName {
		queryParts = append(queryParts, "last_name = ?")
		args = append(args, value.LastName)
		original.LastName = value.LastName

	}

	if ValidString(value.ImagePath) && value.ImagePath != original.ImagePath {
		queryParts = append(queryParts, "image_path = ?")
		args = append(args, value.ImagePath)
		original.ImagePath = value.ImagePath
	}

	if value.AboutMe != original.AboutMe {
		queryParts = append(queryParts, "about_me = ?")
		args = append(args, value.AboutMe)
		original.AboutMe = value.AboutMe

	}

	if value.IsPrivate != original.IsPrivate {
		queryParts = append(queryParts, "is_private = ?")
		args = append(args, value.IsPrivate)
		original.IsPrivate = value.IsPrivate
	}

	if len(queryParts) == 0 {
		return &original, nil // Return the original user if no changes were made
	}

	// Prepare the update query
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = ?", strings.Join(queryParts, ","))
	args = append(args, id)
	modified = &original

	// Execute the update
	if _, err = tx.ExecContext(ctx, query, args...); err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return modified, nil
}

func (store *SQLite3Store) GetUserFollowList(ctx context.Context, userId string, limit, offset int) (users []*types.User, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	SELECT 
		u.id,
		u.nickname,
		u.email,
		u.first_name,
		u.last_name,
		u.date_of_birth,
		u.image_path,
		u.about_me,
		u.is_private,
		u.timestamp
	FROM follow_records f 
	JOIN users u
	ON u.id = f.follower_id
	WHERE f.user_id = ?;`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := new(types.User)
		err = rows.Scan(
			&user.Id,
			&user.Nickname,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.DateOfBirth,
			&user.ImagePath,
			&user.AboutMe,
			&user.IsPrivate,
			&user.Timestamp,
		)
		if err != nil {
			continue
		}

		users = append(users, user)
	}
	return
}

func (store *SQLite3Store) GetUserGroups(ctx context.Context, userId string) (groups []types.Group, err error) {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	SELECT * 
	FROM groups_record 
	WHERE user_id = ?;`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var group types.Group

		err = rows.Scan(
			&group.Id,
			&group.Name,
			&group.Owner,
			&group.Description,
			&group.Image,
			&group.Timestamp,
		)
		if err != nil {
			log.Println(err)
			continue
		}

		groups = append(groups, group)
	}

	return groups, tx.Commit()
}
