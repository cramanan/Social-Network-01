package websocket

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// mxmap is a concurrent-safe map implementation that associates keys of type K
// with values of type V. It embeds a sync.Mutex to ensure that all operations
// on the map are thread-safe, preventing race conditions when accessed by
// multiple goroutines simultaneously.
//
// The map is initialized with a specified key-value type and provides a
// simple interface for common map operations while ensuring safe concurrent
// access.
type mxmap[K comparable, V any] struct {
	*sync.Mutex
	entries map[K]V
}

// Entries returns all entries in the mxmap as a plain map.
func (m mxmap[K, V]) Entries() map[K]V {
	return m.entries
}

// Lookup retrieves the value associated with the given key from the mxmap.
// It returns the value and a boolean indicating whether the key was found.
func (m mxmap[K, V]) Lookup(key K) (value V, ok bool) {
	// Locking the map to ensure thread-safety.
	m.Lock()                
	value, ok = m.entries[key]

	// Unlocking after the operation is complete.
	m.Unlock()              
	return value, ok
}

// Add adds a new key-value pair to the mxmap and returns the value that was added.
func (m *mxmap[K, V]) Add(key K, value V) V {
	// Locking to prevent concurrent modification.
	m.Lock()                

	// Adding the entry to the map.
	m.entries[key] = value  

	// Unlocking after the operation.
	m.Unlock()              
	return value
}

// Remove deletes the entry associated with the given key from the mxmap.
func (m *mxmap[K, V]) Remove(key K) {
	// Locking to ensure thread-safety.
	m.Lock()          

	// Removing the entry from the map.      
	delete(m.entries, key) 
	
	// Unlocking after the operation.
	m.Unlock()              
}

// MxConn represents a websocket connection with a Mutex for thread-safe operations.
type MxConn struct {
	 // Embedding a Mutex for thread-safe operations.
	*sync.Mutex      

	 // The actual WebSocket connection.
	conn *websocket.Conn
}

// WriteJSON sends a JSON-encoded message over the websocket connection.
func (conn *MxConn) WriteJSON(v any) (err error) {
	// Locking to ensure thread-safe access to the connection.
	conn.Lock()               

	// Writing the JSON message to the connection.
	err = conn.conn.WriteJSON(v) 

	// Unlocking after the operation.
	conn.Unlock()             
	return err
}

// ReadJSON reads a JSON-encoded message from the websocket connection.
func (conn *MxConn) ReadJSON(v any) (err error) {
	// Directly reading the JSON message from the connection.
	return conn.conn.ReadJSON(v) 
}

// WebSocket represents a WebSocket server that manages users and chatrooms.
type WebSocket struct {
	// WebSocket upgrader to handle connections.
	upgrader websocket.Upgrader 

	// A map of connected users.
	Users     mxmap[string, *MxConn] 

	// A map of active chatrooms.
	Chatrooms mxmap[string, *ChatRoom] 
}

// Upgrade handles the HTTP upgrade for the WebSocket connection and returns an MxConn instance.
func (ws *WebSocket) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*MxConn, error) {
	// Upgrading the HTTP connection to WebSocket.
	conn, err := ws.upgrader.Upgrade(w, r, responseHeader) 

	// Returning the MxConn with a Mutex for thread-safe operations.
	return &MxConn{Mutex: new(sync.Mutex), conn: conn}, err 
}

// ChatRoom represents a chatroom, which is a map of WebSocket connections.
// A chatroom is just a map of user connections.
type ChatRoom = mxmap[string, *MxConn] 

// NewWebSocket initializes and returns a new WebSocket server instance with default settings.
func NewWebSocket() WebSocket {
	return WebSocket{
		upgrader: websocket.Upgrader{
			// Allow all origins by default.
			CheckOrigin: func(r *http.Request) bool { return true }, 
		},
		Users:     mxmap[string, *MxConn]{Mutex: new(sync.Mutex), entries: make(map[string]*MxConn)},
		Chatrooms: mxmap[string, *ChatRoom]{Mutex: new(sync.Mutex), entries: make(map[string]*ChatRoom)},
	}
}

// NewChatRoom creates and returns a new, empty chatroom.
func NewChatRoom() *ChatRoom {
	return &mxmap[string, *MxConn]{Mutex: new(sync.Mutex), entries: make(map[string]*MxConn)}
}

