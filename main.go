package main
import (
	
	"github.com/paerarason/go-Auction/ws"
	"net/http"
	"log"
)

func main(){

	mux := http.NewServeMux()
	mux.Handle("/",)
	http.HandleFunc("/",ws.WEBsocket)
	http.HandleFunc("/login",)
    
    err := http.ListenAndServe(":3333", nil)  
	if err==nil{
		log.Println("SERVER RUNNING On port 3333")
	}
	log.Println("SERVER RUNNING On port 3333")
}