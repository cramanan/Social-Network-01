package websocket

import (
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

type WebSocket struct {
	websocket.Upgrader

	Users     mxmap[string, *websocket.Conn]
	Chatrooms mxmap[string, *ChatRoom]
}

type ChatRoom = mxmap[string, *websocket.Conn]

func NewWebSocket() WebSocket {
	return WebSocket{
		Upgrader:  websocket.Upgrader{},
		Users:     mxmap[string, *websocket.Conn]{Mutex: new(sync.Mutex), entries: make(map[string]*websocket.Conn)},
		Chatrooms: mxmap[string, *ChatRoom]{Mutex: new(sync.Mutex), entries: make(map[string]*ChatRoom)},
	}
}

func NewChatRoom() *ChatRoom {
	return &mxmap[string, *websocket.Conn]{Mutex: new(sync.Mutex), entries: make(map[string]*websocket.Conn)}
}
