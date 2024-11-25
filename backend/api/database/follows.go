package database

import (
	"Social-Network-01/api/types"
	"context"
	"database/sql"
)

// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding followed user in the database and is usualy find in the request pathvalue.
// `followerId` is the corresponding following user in the database and is usualy find in the sessions field of the API structure.
//
// This method return a boolean corresponding at the state of the follow and/or usualy an SQL error.
func (store *SQLite3Store) RequestedFollow(ctx context.Context, userId, followerId string) (follows bool, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return false, err
	}
	defer tx.Rollback()

	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS(
		SELECT 1 FROM follow_records 
		WHERE user_id = ? AND follower_id = ?
	);`, userId, followerId).Scan(&follows)

	if err != nil {
		return false, err
	}

	return follows, err
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

	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS(
		SELECT 1 FROM follow_records 
		WHERE user_id = ? AND follower_id = ? AND accepted = TRUE
	);`, userId, followerId).Scan(&follows)

	if err != nil {
		return false, err
	}

	return follows, err
}

// Perform the action of following one from another in the database using their userids.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding followed user in the database and is usualy find in the request pathvalue.
// `followerId` is the corresponding following user in the database and is usualy find in the sessions field of the API structure.
//
// This method return an SQL error or nil if there are none.
func (store *SQLite3Store) SendFriendRequest(ctx context.Context, userId, followerId string) error {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	rowsExists := false
	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT 1 
		FROM users 
		WHERE id = ?
	) AND EXISTS (
		SELECT 1 
		FROM users 
		WHERE id = ?
	);`, userId, followerId).Scan(&rowsExists)
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(ctx, `
	INSERT INTO follow_records (user_id, follower_id, accepted) 
	VALUES (?, ?, (
		SELECT NOT is_private FROM users WHERE id = ?
	));`, userId, followerId, userId)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (store *SQLite3Store) AcceptFriendRequest(ctx context.Context, userId, followerId string) error {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `
	UPDATE follow_records 
	SET accepted = TRUE 
	WHERE user_id = ? and follower_id = ?;`,
		userId, followerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

// Perform the action of unfollowing one from another in the database using their userids.
//
// `store` is find in the API structure and is the SQLite3 DB.
// `ctx` is the context of the request. `userId` is the corresponding followed user in the database and is usualy find in the request pathvalue.
// `followerId` is the corresponding following user in the database and is usualy find in the sessions field of the API structure.
//
// This method return an SQL error or nil if there are none.
func (store *SQLite3Store) UnfollowUser(ctx context.Context, userId, followerId string) error {
	tx, err := store.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.ExecContext(ctx, `
	DELETE FROM follow_records
	WHERE user_id = ? AND follower_id = ?;`,
		userId, followerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (store *SQLite3Store) GetFriendRequests(ctx context.Context, userId string) (users []types.User, err error) {
	tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	rows, err := tx.QueryContext(ctx, `
	SELECT u.nickname, u.image_path 
	FROM follow_records f JOIN users u
	ON f.follower_id = u.id
	WHERE f.user_id = ? AND f.accepted = FALSE;`,
		userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user types.User
		err = rows.Scan(&user.Nickname, &user.ImagePath)
		if err != nil {
			continue
		}

		users = append(users, user)
	}

	if users == nil {
		users = make([]types.User, 0)
	}

	return users, err
}
