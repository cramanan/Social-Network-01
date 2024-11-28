package api

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/mail"

	"Social-Network-01/api/database"
	"Social-Network-01/api/types"
)

/// Register handles the registration of a new user. It expects user details in the request body,
// performs validation, and stores the user data in the database.
func (server *API) Register(writer http.ResponseWriter, request *http.Request) error {
    // Ensure the request method is POST for user registration
    if request.Method != http.MethodPost {
        return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
    }

    // Parse the registration request from the request body
    registerReq := new(types.RegisterRequest)
    err := json.NewDecoder(request.Body).Decode(registerReq)
    if err != nil {
        return writeJSON(writer, http.StatusUnprocessableEntity, HTTPerror(http.StatusUnprocessableEntity))
    }

    // Check if all required fields are provided in the registration request
    if registerReq.Nickname == "" || registerReq.Email == "" || registerReq.Password == "" ||
        registerReq.FirstName == "" || registerReq.LastName == "" || registerReq.DateOfBirth == "" {
        return writeJSON(writer, http.StatusUnauthorized, HTTPerror(http.StatusBadRequest, "All fields are required"))
    }

    // Validate the email address format
    if _, err = mail.ParseAddress(registerReq.Email); err != nil {
        return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusBadRequest, "Invalid Email address"))
    }

    // Attempt to register the user in the storage/database
    user, err := server.Storage.RegisterUser(request.Context(), registerReq)
    if errors.Is(err, database.ErrConflict) {
        return writeJSON(writer, http.StatusConflict, HTTPerror(http.StatusConflict, "Email address is already taken"))
    }
    if err != nil {
        return err
    }

    // Create a new session for the registered user
    session := server.Sessions.NewSession(writer, request)
    session.User = user

    // Respond with the created user data
    return writeJSON(writer, http.StatusCreated, user)
}

// Login handles the user login process. It validates the user's credentials and creates a session if valid.
func (server *API) Login(writer http.ResponseWriter, request *http.Request) (err error) {
    // Ensure the request method is POST for login
    if request.Method != http.MethodPost {
        return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
    }

    // Parse the login request
    loginReq := new(types.LoginRequest)
    if err := json.NewDecoder(request.Body).Decode(loginReq); err != nil {
        return writeJSON(writer, http.StatusUnprocessableEntity, HTTPerror(http.StatusUnprocessableEntity))
    }

    // Check if email and password are provided
    if loginReq.Email == "" || loginReq.Password == "" {
        return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusUnauthorized, "Email and password are required"))
    }

    // Validate the email format
    if _, err = mail.ParseAddress(loginReq.Email); err != nil {
        return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusBadRequest, "Invalid Email address"))
    }

    // Attempt to log the user in by verifying credentials
    user, err := server.Storage.LogUser(request.Context(), loginReq)
    if err != nil {
        return writeJSON(writer, http.StatusBadRequest, HTTPerror(http.StatusBadRequest, "Invalid Email or Password"))
    }

    // Create a new session for the logged-in user
    session := server.Sessions.NewSession(writer, request)
    session.User = user

    // Respond with the logged-in user's data
    return writeJSON(writer, http.StatusOK, user)
}

// User handles GET requests to retrieve user data by user ID. It can also handle private user visibility based on follows.
func (server *API) User(writer http.ResponseWriter, request *http.Request) (err error) {
    switch request.Method {
    case http.MethodGet:
        // Retrieve the user ID from the request path
        userId := request.PathValue("userid")
        user, err := server.Storage.GetUser(request.Context(), userId)
        if err == sql.ErrNoRows {
            return writeJSON(writer, http.StatusNotFound, HTTPerror(http.StatusNotFound, "User not found"))
        }
        if err != nil {
            return err
        }

        // Get the session to check if the current user is logged in
        sess, err := server.Sessions.GetSession(request)
        if err != nil {
            return err
        }

        // If the user is private and not followed by the current user, hide personal information
        if !user.IsPrivate || sess.User.Id == userId {
            return writeJSON(writer, http.StatusOK, user)
        }

        // Check if the current user follows the requested user
        follows, err := server.Storage.Follows(request.Context(), userId, sess.User.Id)
        if !follows || err != nil {
            // Hide sensitive data if the current user does not follow the user
            user.Email = ""
            user.FirstName = ""
            user.LastName = ""
            user.AboutMe = nil
            return writeJSON(writer, http.StatusUnauthorized, user)
        }

        return writeJSON(writer, http.StatusOK, user)

    default:
        return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
    }
}

// Profile allows the logged-in user to view or update their profile information.
func (server *API) Profile(writer http.ResponseWriter, request *http.Request) (err error) {
    // Retrieve the session of the logged-in user
    sess, err := server.Sessions.GetSession(request)
    if err != nil {
        return err
    }

    switch request.Method {
    case http.MethodGet:
        // Return the user's profile data (session data)
        return writeJSON(writer, http.StatusOK, sess.User)

    case http.MethodPatch:
        // Handle profile updates (e.g., change name, image)
        err = request.ParseMultipartForm(5 * (1 << 20))
        if err != nil {
            return err
        }

        // Create a User object to update profile fields
        user := types.User{}
        data := request.MultipartForm.Value["data"]
        if len(data) != 1 {
            return fmt.Errorf("invalid number of datas")
        }

        // Unmarshal the data into the user object
        err = json.Unmarshal([]byte(data[0]), &user)
        if err != nil {
            return err
        }

        // Handle file uploads (images) for the profile picture
        imgs, err := MultiPartFiles(request)
        if err != nil {
            return err
        }

        // If a new image is provided, update the profile image path
        if len(imgs) > 0 {
            user.ImagePath = imgs[0]
        } else {
            user.ImagePath = ""
        }

        // Update the user profile in the database
        modified, err := server.Storage.UpdateUser(request.Context(), sess.User.Id, user)
        if err != nil {
            return err
        }

        // Update the session with the modified user data
        sess.User = *modified

        return writeJSON(writer, http.StatusOK, modified)

    case http.MethodDelete:
        // Ensure the request is for the logged-in user to delete their account
        if sess.User.Id != request.PathValue("userid") {
            return writeJSON(writer, http.StatusUnauthorized, HTTPerror(http.StatusUnauthorized, "You are not authorized to perform this action."))
        }

        // Delete the user from the database
        err = server.Storage.DeleteUser(request.Context(), sess.User.Id)
        if err != nil {
            return err
        }

        return writeJSON(writer, http.StatusNoContent, "")

    default:
        return writeJSON(writer, http.StatusMethodNotAllowed, HTTPerror(http.StatusMethodNotAllowed))
    }
}

// GetUserStats retrieves and returns statistics for a given user.
func (server *API) GetUserStats(writer http.ResponseWriter, request *http.Request) error {
    stats, err := server.Storage.GetUserStats(request.Context(), request.PathValue("userid"))
    if err != nil {
        return err
    }
    return writeJSON(writer, http.StatusOK, stats)
}

// GetOnlineUsers returns a list of online users who the logged-in user has messaged recently.
func (server *API) GetOnlineUsers(writer http.ResponseWriter, request *http.Request) error {
    // Get the session for the logged-in user
    sess, err := server.Sessions.GetSession(request)
    if err != nil {
        return err
    }

    // Parse pagination parameters
    limit, offset := parseRequestLimitAndOffset(request)

    // Get a list of users the logged-in user has messaged
    users, err := server.Storage.GetMessagedUsers(request.Context(), sess.User.Id, limit, offset)
    if err != nil {
        return err
    }

    // Check which users are currently online
    onlineUsers := make([]types.OnlineUser, len(users))
    for idx, user := range users {
        onlineUsers[idx] = types.OnlineUser{User: user}
        _, onlineUsers[idx].Online = server.WebSocket.Users.Lookup(user.Id)
    }

    return writeJSON(writer, http.StatusOK, onlineUsers)
}

// GetUserFriendList retrieves the list of friends for the logged-in user, including their online status.
func (server *API) GetUserFriendList(writer http.ResponseWriter, request *http.Request) (err error) {
    sess, err := server.Sessions.GetSession(request)
    if err != nil {
        return err
    }

    // Parse pagination parameters
    limit, offset := parseRequestLimitAndOffset(request)

    // Get the user's friend list
    users, err := server.Storage.GetUserFriendList(context.TODO(), sess.User.Id, limit, offset)
    if err != nil {
        return err
    }

    // Check which friends are currently online
    onlineUsers := make([]types.OnlineUser, len(users))
    for idx, user := range users {
        onlineUsers[idx] = types.OnlineUser{User: user}
        _, onlineUsers[idx].Online = server.WebSocket.Users.Lookup(user.Id)
    }

    return writeJSON(writer, http.StatusOK, onlineUsers)
}
