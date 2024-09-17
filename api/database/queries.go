package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"Social-Network-01/api/models"
)

// Retrieve all posts of one user from the database using its userId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding user in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be retrieve with the parseRequestLimitAndOffset method using the request.
//
// This method return an array of post (see ./api/models/posts.go) or usualy an SQL error (one is nil when the other isn't).
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
			&post.ImagePath,
			&post.Timestamp)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

// Retrieve all posts of one group from the database using its groupId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `groupId` is the corresponding group in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be retrieve with the parseRequestLimitAndOffset method using the request.
//
// This method return an array of post (see ./api/models/posts.go) or usualy an SQL error (one is nil when the other isn't).
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
		err := rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Content, &post.ImagePath, &post.Timestamp)
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

// Retrieve all comments of one post from the database using its postId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `postId` is the corresponding post in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be retrieve with the parseRequestLimitAndOffset method using the request.
//
// This method return an array of comment (see ./api/models/comments.go) or usualy an SQL error (one is nil when the other isn't).
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

// Retrieve all posts of ones likes from the database using his userId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding user in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be retrieve with the parseRequestLimitAndOffset method using the request.
//
// This method return an array of post (see ./api/models/posts.go) or usualy an SQL error (one is nil when the other isn't).
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
		err := rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Content, &post.ImagePath, &post.Timestamp)
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

// Retrieve all follower of a user from the database using his userId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding user in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be retrieve with the parseRequestLimitAndOffset method using the request.
//
// This method return an array of user (see ./api/models/users.go) or usualy an SQL error (one is nil when the other isn't).
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

// Indicate if a user follow another or not in the database.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding followed user in the database and is usualy find in the request pathvalue.
// `followerId` is the corresponding following user in the database and is usualy find in the sessions field of the API structure.
//
// This method return a boolean corresponding at the state of the follow and/or usualy an SQL error.
func (store *SQLite3Store) Follows(ctx context.Context, userId, followerId string) (follows bool, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	return follows, tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM follow_records WHERE user_id = ? and follower_id = ?)").Scan(&follows)
}

// Retrieve all chats beetween 2 users from the database using their userIds.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `user1Id` and `user2Id` are the corresponding users in the database and are usualy find in the
// request pathvalue and in the sessions field of the API structure. `limit` and `offset` can be retrieve with the parseRequestLimitAndOffset
// method using the request.
//
// This method return an array of chat (see ./api/models/chat.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetChats(ctx context.Context, user1Id, user2Id string, limit, offset int) (chats []models.Chat, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx,
		`SELECT * 
		FROM chats 
		WHERE (sender_id = ? AND recipient_id = ?)
		OR (recipient_id = ? AND sender_id = ?)
		ORDER BY timestamp DESC 
		LIMIT ? OFFSET ? ;`, user1Id, user2Id, user1Id, user2Id, limit, offset)
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

// Retrieve all posts of a user's follows from the database using his userId.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding user in the database and is usualy find in the request pathvalue.
// `limit` and `offset` can be retrieve with the parseRequestLimitAndOffset method using the request.
//
// This method return an array of posts (see ./api/models/posts.go) or usualy an SQL error (one is nil when the other isn't).
func (store *SQLite3Store) GetFollowsPosts(ctx context.Context, userId string, limit, offset int) (posts []models.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx,
		`SELECT * 
		FROM posts p 
		JOIN follow_records f 
		ON p.userid = f.user_id 
		WHERE f.follower_id = ? 
		ORDER BY timestamp DESC 
		LIMIT ? OFFSET ?;`, userId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		post := models.Post{}
		err := rows.Scan(&post.Id, &post.UserId, &post.GroupId, &post.Content, &post.ImagePath, &post.Timestamp)
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

