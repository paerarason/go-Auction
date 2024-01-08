package ws 
import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/google/uuid"
	"github.com/paerarason/go-Auction/controller"
)

server:=&Server{subscriptionubscriptions: make(Subscription)}
func WEBsocket(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },}
	conn,err:=upgrader.Upgrade(w,r,nil)
	if err != nil {
            log.Print("upgrade failed: ", err)
            return
        }
	defer conn.Close()
	
    clientID := uuid.New().String()
	server.Send(conn,fmt.Sprintf("Hello! Your where connected to Server is %s", clientID))
	
	for {
		// read incoming message
		_, msg, err := conn.ReadMessage()
		// if error occured
		if err != nil {
			server.RemoveClient(clientID)
			break
		}
		server.ProcessMessage(conn,clientID,msg) 
	}

}