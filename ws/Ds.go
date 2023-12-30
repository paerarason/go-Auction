package ws

import (
	"github.com/gorilla/websocket"
)

type Server struct{
	subscription Subscription
}
type Subscription map[string]Client

type Client map[string] *websocket.conn

type Message struct  {
	action  string `json:"action"`
	topic string `json:topic`
    message string `json:message`
}



