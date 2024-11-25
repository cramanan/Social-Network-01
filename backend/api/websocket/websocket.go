package websocket

import (
	"sync"

	"github.com/gorilla/websocket"
)

type SocketConn struct {
	*websocket.Conn
	sync.Mutex
}

type Userlist map[string]*SocketConn

func (conn *SocketConn) WriteJSON(v any) (err error) {
	conn.Lock()
	err = conn.Conn.WriteJSON(v)
	conn.Unlock()
	return err
}

type WebSocket struct {
	sync.Mutex

	Upgrader  websocket.Upgrader
	Users     Userlist
	ChatRooms map[string]ChatRoom
}

func (socket *WebSocket) AddUser(id string, conn *SocketConn) {
	socket.Lock()
	socket.Users[id] = conn
	socket.Unlock()
}

func (socket *WebSocket) RemoveUser(id string) {
	socket.Lock()
	delete(socket.Users, id)
	socket.Unlock()
}

type ChatRoom struct {
	*sync.Mutex
	Users Userlist
}
