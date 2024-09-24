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

func (store *SQLite3Store) CreatePost(ctx context.Context, req *models.PostRequest) (group *models.Post, err error) {
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
	marshalImages, err := json.Marshal(req.Images)
	if err != nil {
		return nil, err
	}

	tx.ExecContext(ctx,
		`INSERT INTO posts VALUES(?, ?, ?, ?, ?, ?);`,
		id.String(),
		req.UserId,
		req.GroupId,
		req.Content,
		marshalImages,
		time.Now(),
	)

	return group, tx.Commit()
}

func (store *SQLite3Store) GetPost(ctx context.Context, postId string) (post *models.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	post = new(models.Post)

	err = tx.QueryRowContext(ctx,
		"SELECT * FROM posts WHERE id = ?;",
	).Scan(
		post.Id,
		post.UserId,
		post.GroupId,
		post.Content,
		post.Timestamp,
	)
	return post, nil
}

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
