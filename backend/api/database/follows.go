package database

import (
	"Social-Network-01/api/types"
	"context"
	"database/sql"
)

// RequestedFollow checks if a follow request exists between two users.
// - `userId`: The ID of the user being followed.
// - `followerId`: The ID of the user sending the follow request.
// Returns a boolean indicating the existence of the follow request and/or an SQL error.
func (store *SQLite3Store) RequestedFollow(ctx context.Context, userId, followerId string) (follows bool, err error) {
    // Start a read-only transaction.
    tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
    if err != nil {
        return false, err
    }
    // Ensure transaction rollback on error.
    defer tx.Rollback()

    // Query to check if a follow request exists.
    err = tx.QueryRowContext(ctx, `
	SELECT EXISTS(
		SELECT 1 FROM follow_records 
		WHERE user_id = ? AND follower_id = ?
	);`, userId, followerId).Scan(&follows)

    // Return the result or an error.
    if err != nil {
        return false, err
    }
    return follows, err
}

// Follows checks if a user follows another (accepted follow request).
// - `userId`: The ID of the user being followed.
// - `followerId`: The ID of the user following.
// Returns a boolean indicating the follow status and/or an SQL error.
func (store *SQLite3Store) Follows(ctx context.Context, userId, followerId string) (follows bool, err error) {
    // Start a read-only transaction.
    tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
    if err != nil {
        return false, err
    }
    // Ensure transaction rollback on error.
    defer tx.Rollback()

    // Query to check if the follow is accepted.
    err = tx.QueryRowContext(ctx, `
	SELECT EXISTS(
		SELECT 1 FROM follow_records 
		WHERE user_id = ? AND follower_id = ? AND accepted = TRUE
	);`, userId, followerId).Scan(&follows)

    // Return the result or an error.
    if err != nil {
        return false, err
    }
    return follows, err
}

// SendFriendRequest creates a follow request from one user to another.
// - `userId`: The ID of the user being followed.
// - `followerId`: The ID of the user sending the follow request.
// Automatically accepts the follow request if the target user is not private.
// Returns an SQL error if any issues occur.
func (store *SQLite3Store) SendFriendRequest(ctx context.Context, userId, followerId string) error {
    // Start a transaction.
    tx, err := store.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    // Ensure transaction rollback on error.
    defer tx.Rollback()

    // Check if both users exist in the database.
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

    // Insert a follow request into the follow_records table.
    _, err = tx.ExecContext(ctx, `
	INSERT INTO follow_records (user_id, follower_id, accepted) 
	VALUES (?, ?, (
		SELECT NOT is_private FROM users WHERE id = ?
	));`, userId, followerId, userId)
    if err != nil {
        return err
    }
    // Commit the transaction.
    return tx.Commit()
}

// AcceptFriendRequest marks a follow request as accepted.
// - `userId`: The ID of the user being followed.
// - `followerId`: The ID of the user who sent the follow request.
// Returns an SQL error if any issues occur.
func (store *SQLite3Store) AcceptFriendRequest(ctx context.Context, userId, followerId string) error {
    // Start a transaction.
    tx, err := store.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    // Ensure transaction rollback on error.
    defer tx.Rollback()

    // Update the follow_records table to mark the follow request as accepted.
    _, err = tx.ExecContext(ctx, `
	UPDATE follow_records 
	SET accepted = TRUE
	WHERE user_id = ? and follower_id = ?;`,
        userId, followerId)
    if err != nil {
        return err
    }

    // Commit the transaction.
    return tx.Commit()
}

// UnfollowUser removes a follow relationship between two users.
// - `userId`: The ID of the user being unfollowed.
// - `followerId`: The ID of the user unfollowing.
// Returns an SQL error if any issues occur.
func (store *SQLite3Store) UnfollowUser(ctx context.Context, userId, followerId string) error {
    // Start a transaction.
    tx, err := store.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    // Ensure transaction rollback on error.
    defer tx.Rollback()

    // Delete the follow relationship from the follow_records table.
    _, err = tx.ExecContext(ctx, `
	DELETE FROM follow_records
	WHERE user_id = ? AND follower_id = ?;`,
        userId, followerId)
    if err != nil {
        return err
    }

    // Commit the transaction.
    return tx.Commit()
}

// GetFriendRequests retrieves all pending follow requests for a user.
// - `userId`: The ID of the user receiving the follow requests.
// Returns a slice of User objects representing the followers and/or an SQL error.
func (store *SQLite3Store) GetFriendRequests(ctx context.Context, userId string) (users []types.User, err error) {
    // Start a read-only transaction.
    tx, err := store.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
    if err != nil {
        return nil, err
    }
    // Ensure transaction rollback on error.
    defer tx.Rollback()

    // Query to fetch pending follow requests.
    rows, err := tx.QueryContext(ctx, `
	SELECT u.id, u.nickname, u.image_path 
	FROM follow_records f JOIN users u
	ON f.follower_id = u.id
	WHERE f.user_id = ? AND f.accepted = FALSE;`,
        userId)
    if err != nil {
        return nil, err
    }
    // Ensure the rows are closed after processing.
    defer rows.Close()

    // Iterate through the result set and populate the users slice.
    for rows.Next() {
        var user types.User
        err = rows.Scan(&user.Id, &user.Nickname, &user.ImagePath)
        if err != nil {
            continue // Skip rows with errors.
        }

        users = append(users, user)
    }

    // Return an empty slice if no follow requests are found.
    if users == nil {
        users = make([]types.User, 0)
    }

    return users, err
}
