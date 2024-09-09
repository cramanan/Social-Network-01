package database

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"Social-Network-01/api/models"

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

// Recover all user datas from one user from the database using its userId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding user in the database and is usualy find in the request pathvalue or
// in the sessions field of the API structure.
//
// This function return a user (see ./api/models/users.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetUser(ctx context.Context, userId string) (user *models.User, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, "SELECT * FROM users WHERE id = ?;", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

// Recover all posts from one user from the database using its userId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding user in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be recover with the parseRequestLimitAndOffset function using the request.
//
// This function return an array of post (see ./api/models/posts.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetAllPostsFromOneUser(ctx context.Context, userId string, limit, offset int) (posts []*models.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, "SELECT * FROM posts WHERE user_id = ? LIMIT ? OFFSET ?;", userId, limit, offset)
	for rows.Next() {
		post := new(models.Post)
		err = rows.Scan(
			&post.Id,
			&post.UserId,
			&post.GroupId,
			&post.Categories,
			&post.ImagePath,
			&post.Timestamp)
		posts = append(posts, post)

	}
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Recover all posts from one group from the database using its groupId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `groupId` is the corresponding group in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be recover with the parseRequestLimitAndOffset function using the request.
//
// This function return an array of post (see ./api/models/posts.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetGroupPosts(ctx context.Context, groupId string, limit, offset int) (posts []models.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, "SELECT * FROM posts WHERE group_id = ? ORDER BY timestamp DESC LIMIT ? OFFSET ? ORDER BY timestamp DESC;", groupId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Categories, &post.Content, &post.ImagePath, &post.Timestamp)
		if err != nil {
			fmt.Println(err)
			continue
		}

		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Recover all comments from one post from the database using its postId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `postId` is the corresponding post in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be recover with the parseRequestLimitAndOffset function using the request.
//
// This function return an array of comment (see ./api/models/comments.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetComments(ctx context.Context, postId string, limit, offset int) (comments []models.Comments, err error) {
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
		comment := models.Comments{}
		err := rows.Scan(&comment.Id, &comment.UserId, &comment.ParentId, &comment.Content, &comment.ImgPath, &comment.TimeStamp)
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

// Recover all posts from ones likes from the database using his userId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding user in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be recover with the parseRequestLimitAndOffset function using the request.
//
// This function return an array of post (see ./api/models/posts.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetPostsLike(ctx context.Context, userId string, limit, offset int) (posts []models.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
    WITH liked_posts AS (
        SELECT DISTINCT post_id
        FROM likes_records
        WHERE user_id = ?
    )
    SELECT p.*
    FROM posts p
    LEFT JOIN liked_posts lp ON p.id = lp.postid
    ORDER BY p.timestamp DESC;
`, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Categories, &post.Content, &post.ImagePath, &post.Timestamp)
		if err != nil {
			fmt.Println(err)
			continue
		}

		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return posts, nil
}

// Recover all follower from a user from the database using his userId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding user in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be recover with the parseRequestLimitAndOffset function using the request.
//
// This function return an array of user (see ./api/models/users.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetFollowersOfUser(ctx context.Context, userId string, limit, offset int) (users []models.User, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := store.QueryContext(ctx,
		`SELECT 
			u.id,
			u.nickname,
			u.email,
			u.first_name,
			u.last_name,
			u.date_of_birth,
			u.image_path,
			u.about_me,
			u.private,
			u.timestamp
		FROM follow_records f 
		JOIN users u 
		ON f.user_id = u.id
		WHERE user_id = ?
		LIMIT ? OFFSET ?;`, userId, limit, offset)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := models.User{}
		err = rows.Scan(
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
		users = append(users, user)
	}

	return users, nil
}

func (store *SQLite3Store) FollowUser(ctx context.Context, userId, followerId string) error {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = store.ExecContext(ctx, `INSERT INTO follow_records VALUES(?, ?);`, userId, followerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (store *SQLite3Store) Follows(ctx context.Context, userId, followerId string) (follows bool, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	return follows, tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM likes_records WHERE user_id = ? and follower_id = ?)").Scan(follows)
}

// Recover all chats beetween 2 users from the database using their userIds.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `user1Id` and `user2Id` are the corresponding users in the database and are usualy find in the
// request pathvalue and in the sessions field of the API structure. `limit` and `offset` can be recover with the parseRequestLimitAndOffset
// function using the request.
//
// This function return an array of chat (see ./api/models/chat.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetChats(ctx context.Context, user1Id, user2Id string, limit, offset int) (chats []models.Chat, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, "SELECT * FROM chats WHERE (sender_id = ? AND recipient_id = ?) OR (recipient_id = ? AND sender_id = ?) ORDER BY timestamp DESC LIMIT ? OFFSET ? ;", user1Id, user2Id, user1Id, user2Id, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		chat := models.Chat{}
		err := rows.Scan(&chat.ID, &chat.SenderId, &chat.RecipientId, &chat.Content, &chat.ImgPath, &chat.Timestamp)
		if err != nil {
			fmt.Println(err)
			continue
		}

		chats = append(chats, chat)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return chats, nil
}

// Recover all posts from a user's follows from the database using his userId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding user in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be recover with the parseRequestLimitAndOffset function using the request.
//
// This function return an array of posts (see ./api/models/posts.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetFollowsPosts(ctx context.Context, userId string, limit, offset int) (posts []models.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, "SELECT * FROM [...] ORDER BY timestamp DESC LIMIT ? OFFSET ? ORDER BY timestamp DESC;", userId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Categories, &post.Content, &post.ImagePath, &post.Timestamp)
		if err != nil {
			fmt.Println(err)
			continue
		}

		posts = append(posts, post)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return posts, nil
}
