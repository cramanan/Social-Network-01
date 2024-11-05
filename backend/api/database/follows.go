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

	return follows, tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM follow_records WHERE user_id = ? and follower_id = ?)").Scan(&follows)
}