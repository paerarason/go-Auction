package controller

type JSON struct{
    user_id int 
    current_bid int
} 

Chan:=make(chan JSON)



func Bidding(w http.ResponseWriter, r *http.Request){
	
}