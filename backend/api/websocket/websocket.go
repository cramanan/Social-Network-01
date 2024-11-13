package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type SocketConn struct {
	*websocket.Conn
	sync.Mutex
}

func (conn *SocketConn) WriteJSON(v any) (err error) {
	conn.Lock()
	err = conn.Conn.WriteJSON(v)
	conn.Unlock()
	return err
}

type WebSocket struct {
	sync.Mutex

	Upgrader websocket.Upgrader
	Users    map[string]*SocketConn
}

func (socket *WebSocket) Add(id string, conn *SocketConn) {
	socket.Lock()
	socket.Users[id] = conn
	socket.Unlock()
}

func (socket *WebSocket) Remove(id string) {
	socket.Lock()
	delete(socket.Users, id)
	socket.Unlock()
}
