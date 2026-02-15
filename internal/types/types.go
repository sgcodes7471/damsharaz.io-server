package types

import(
	"github.com/gorilla/websocket"
)

type Room_Object struct {
	RoomId string 
	Token string
	Den *websocket.Conn
	Ongoing bool
}

