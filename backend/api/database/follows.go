package database

import (
	"context"
	"database/sql"
)

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
	);`).Scan(&follows)

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

	query := "INSERT INTO follow_records VALUES(?, ?, FALSE);"
	var alreadyFollows bool

	err = tx.QueryRowContext(ctx, `
	SELECT EXISTS (
		SELECT * FROM follow_records 
		WHERE user_id = ? AND follower_id = ?
	);`, userId, followerId).Scan(&alreadyFollows)
	if err != nil {
		return err
	}

	if alreadyFollows {
		query = "DELETE FROM follow_records WHERE user_id = ? AND follower_id = ?;"
	}

	_, err = store.ExecContext(ctx, query, userId, followerId)
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

	_, err = store.ExecContext(ctx, `
	DELETE FROM follow_records
	WHERE user_id = ? AND follower_id = ?;`,
		userId, followerId)
	if err != nil {
		return err
	}

	return tx.Commit()
}
