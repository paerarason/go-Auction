package ws 
import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/paerarason/go-Auction/controller"

)



func WEBsocket(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool { return true },}
    var chan 
  
	conn,err:=upgrader.Upgrade(w,r,nil)
	if err != nil {
            log.Print("upgrade failed: ", err)
            return
        }
    defer conn.Close()

	for {
         msg:=<-controller.Chan
		}

		
}
