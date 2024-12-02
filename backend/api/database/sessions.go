package database

import (
	"Social-Network-01/api/types"
	"errors"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	session_timeout = time.Hour
	cookie_name     = "SESSION-ID"
)

type SessionStore struct {
	mx       sync.RWMutex
	sessions map[string]*Session
}

// generateB64 generates a random base-64 ID of specified length `n`.
// The ID is composed of alphanumeric characters and "+-".
// Note: Uses `math/rand`, which is not cryptographically secure. 
//       Consider using `crypto/rand` for improved security if required.
func generateB64(n int) string {
    var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890+-")
    id := make([]rune, n)
    for i := range id {
        id[i] = letters[rand.Intn(len(letters))]
    }
    return string(id)
}

// NewSessionStore initializes and returns a new SessionStore instance.
// It creates an empty session map and starts a background goroutine 
// to handle session expiration through `timeoutCycle`.
func NewSessionStore() *SessionStore {
    store := new(SessionStore)
    store.sessions = make(map[string]*Session)
    store.timeoutCycle() // Starts periodic cleanup of expired sessions.
    return store
}

// timeoutCycle periodically checks and removes expired sessions from the store.
// Runs in a separate goroutine with a ticker based on `session_timeout` duration.
// Note: This approach might not scale efficiently for large session stores 
//       as it iterates over all sessions every cycle.
func (store *SessionStore) timeoutCycle() {
    go func() {
        for range time.NewTicker(session_timeout).C {
            store.mx.Lock() // Ensures thread-safe access to the session map.
            for key, sess := range store.sessions {
                if sess.Expires.Before(time.Now()) { // Checks if the session is expired.
                    log.Println("Deleted", key) // Logs the session ID being removed.
                    delete(store.sessions, key) // Removes expired session.
                }
            }
            store.mx.Unlock()
        }
    }()
}

// Session structure holds the session's unique ID, associated user details, 
// and its expiration time.
type Session struct {
    ID      string       // Unique identifier for the session.
    User    types.User   // Contains information about the associated user.
    Expires time.Time    // The time when the session expires.
}

// NewSession creates a new session, sets a cookie in the HTTP response, 
// and stores the session in the session map.
// Note: The cookie has `HttpOnly` set to `false`, allowing client-side access.
//       This can be a security risk if not required.
func (store *SessionStore) NewSession(w http.ResponseWriter, r *http.Request) *Session {
    session := new(Session)
    session.ID = generateB64(5) // Generates a unique 5-character session ID.
    cookie := http.Cookie{
        Name:     cookie_name,                       // Name of the cookie.
        Value:    session.ID,                       // Session ID as cookie value.
        Expires:  time.Now().Add(session_timeout),  // Expiration time.
        Path:     "/",                              // Cookie is valid for all paths.
        SameSite: http.SameSiteLaxMode,             // Limits cross-site access.
        HttpOnly: false,                            // Allows client-side access.
    }
    session.Expires = cookie.Expires
    http.SetCookie(w, &cookie) // Sets the cookie in the HTTP response.
    store.mx.Lock()
    store.sessions[session.ID] = session // Stores the session in the map.
    store.mx.Unlock()
    return session
}

// GetSession retrieves a session associated with the request's cookie.
// Returns an error if no valid session is found.
// Note: If the cookie is missing or invalid, the method returns an error.
func (store *SessionStore) GetSession(r *http.Request) (s *Session, err error) {
    cookie, err := r.Cookie(cookie_name) // Retrieves the session cookie.
    if err != nil {
        return nil, err // No cookie or invalid cookie.
    }

    s, ok := store.sessions[cookie.Value] // Looks up the session by ID.
    if !ok {
        return nil, errors.New("no session found") // Session not found.
    }
    return s, nil
}

// EndSession removes the session associated with the request and invalidates its cookie.
// The session is deleted from the session map, and the cookie is updated with an expired timestamp.
// Note: `HttpOnly` is set to `false` in the cookie, which may expose it to client-side scripts.
func (store *SessionStore) EndSession(w http.ResponseWriter, r *http.Request) error {
    session, err := store.GetSession(r) // Retrieves the session for the request.
    if err != nil {
        return err // Returns an error if no session is found.
    }

    http.SetCookie(w, &http.Cookie{ // Invalidates the session cookie.
        Name:     cookie_name,
        Value:    session.ID,
        Expires:  time.Now(), // Sets expiration time to the past.
        Path:     "/",
        SameSite: http.SameSiteNoneMode,
        HttpOnly: false, // Allows client-side access.
    })

    store.mx.Lock()
    delete(store.sessions, session.ID) // Deletes the session from the map.
    store.mx.Unlock()
    return nil
}
