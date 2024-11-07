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
	timer    *time.Ticker
}

// generateB64 generate a random base 64 id of n length
func generateB64(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890+-")
	id := make([]rune, n)
	for i := range id {
		id[i] = letters[rand.Intn(len(letters))]
	}
	return string(id)
}

// NewSQLite3Store opens a connection with api/database.sqlite and returns a pointer to a SQLite3Store
func NewSessionStore() *SessionStore {
	store := new(SessionStore)
	store.sessions = make(map[string]*Session)
	store.timeoutCycle()
	return store
}

// The timeoutCycle function initialize the expiration ticker and
// asynchronously handle ticks to delete the Sessions
func (store *SessionStore) timeoutCycle() {
	store.timer = time.NewTicker(session_timeout)
	go func() {
		for range store.timer.C {
			store.mx.Lock()
			for key, sess := range store.sessions {
				if sess.Expires.Before(time.Now()) {
					log.Println("Deleted", key)
					delete(store.sessions, key)
				}
			}
			store.mx.Unlock()
		}
	}()
}

// The Session type holds the User informations and it's expiration time
type Session struct {
	ID      string
	User    types.User
	Expires time.Time
}

// The NewSession method initialize a Session and set the writer's cookie value to the Sessions ID and map[key]
func (store *SessionStore) NewSession(w http.ResponseWriter, r *http.Request) *Session {
	session := new(Session)
	session.ID = generateB64(5)
	cookie := http.Cookie{
		Name:     cookie_name,
		Value:    session.ID,
		Expires:  time.Now().Add(session_timeout),
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		HttpOnly: false,
	}
	session.Expires = cookie.Expires
	http.SetCookie(w, &cookie)
	store.mx.Lock()
	store.sessions[session.ID] = session
	store.mx.Unlock()
	return session
}

// The GetSession method retrieve the session of a request.
func (store *SessionStore) GetSession(r *http.Request) (s *Session, err error) {
	cookie, err := r.Cookie(cookie_name)
	if err != nil {
		return nil, err
	}

	s, ok := store.sessions[cookie.Value]
	if !ok {
		return nil, errors.New("no session found")
	}
	return s, nil
}

// The EndSession method removes the request's Session from the Session map
func (store *SessionStore) EndSession(w http.ResponseWriter, r *http.Request) error {
	session, err := store.GetSession(r)
	if err != nil {
		return err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     cookie_name,
		Value:    session.ID,
		Expires:  time.Now(),
		Path:     "/",
		SameSite: http.SameSiteNoneMode,
		HttpOnly: false,
	})

	store.mx.Lock()
	delete(store.sessions, session.ID)
	store.mx.Unlock()
	return nil
}
