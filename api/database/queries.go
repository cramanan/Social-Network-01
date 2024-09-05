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

func (store *SQLite3Store) GetGroupPosts(ctx context.Context, groupId string, limit, offset int) (posts []models.Post, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, "SELECT * FROM posts WHERE group_id = ? LIMIT ? OFFSET ? ORDER BY timestamp DESC;", groupId, limit, offset)
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

func (store *SQLite3Store) GetComments(ctx context.Context, postId string, limit, offset int) (comments []models.Comments, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, "SELECT * FROM comments WHERE parent_id = ? LIMIT ? OFFSET ? ORDER BY timestamp DESC;", postId, limit, offset)
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
