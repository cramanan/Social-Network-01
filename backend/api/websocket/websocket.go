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

func (m mxmap[K, V]) Entries() map[K]V {
	return m.entries
}

func (m mxmap[K, V]) Lookup(key K) (value V, ok bool) {
	m.Lock()
	value, ok = m.entries[key]
	m.Unlock()
	return value, ok
}

func (m *mxmap[K, V]) Add(key K, value V) V {
	m.Lock()
	m.entries[key] = value
	m.Unlock()
	return value
}

func (m *mxmap[K, V]) Remove(key K) {
	m.Lock()
	delete(m.entries, key)
	m.Unlock()
}

type MxConn struct {
	*sync.Mutex
	conn *websocket.Conn
}

func (conn *MxConn) WriteJSON(v any) (err error) {
	conn.Lock()
	err = conn.conn.WriteJSON(v)
	conn.Unlock()
	return err
}

func (conn *MxConn) ReadJSON(v any) (err error) {
	return conn.conn.ReadJSON(v)
}

type WebSocket struct {
	upgrader websocket.Upgrader

	Users     mxmap[string, *MxConn]
	Chatrooms mxmap[string, *ChatRoom]
}

func (ws *WebSocket) Upgrade(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*MxConn, error) {
	conn, err := ws.upgrader.Upgrade(w, r, responseHeader)
	return &MxConn{Mutex: new(sync.Mutex), conn: conn}, err
}

type ChatRoom = mxmap[string, *MxConn]

func NewWebSocket() WebSocket {
	return WebSocket{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		},
		Users:     mxmap[string, *MxConn]{Mutex: new(sync.Mutex), entries: make(map[string]*MxConn)},
		Chatrooms: mxmap[string, *ChatRoom]{Mutex: new(sync.Mutex), entries: make(map[string]*ChatRoom)},
	}
}

func NewChatRoom() *ChatRoom {
	return &mxmap[string, *MxConn]{Mutex: new(sync.Mutex), entries: make(map[string]*MxConn)}
}
