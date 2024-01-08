package main
import (
	"github.com/paerarason/go-Auction/controller"
	"github.com/paerarason/go-Auction/ws"
	"net/http"
	"log"
)

func main(){

	mux := http.NewServeMux()
	mux.Handle("/api")
	http.HandleFunc("/ws",ws.WEBsocket)
	http.HandleFunc("/login",controller.LoginHandler)
   
	
    err := http.ListenAndServe(":3333", nil)  
	if err==nil{
		log.Println("SERVER RUNNING On port 3333")
	}
	log.Println("SERVER RUNNING On port 3333")
}