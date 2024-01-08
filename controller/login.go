package cotroller
import (
	"net/http"
)

type Login struct {
   username string 
   password string 
}

func LoginHandler(w http.ResponseWriter,r *http.Request){
	if r.Method!="POST"{
       Response(&w,"Status Created",400)
	   return 
	}
	//jSON DECODING
	var   user  Login
    err:=json.Unmarshal(&user,r.Body())
	if err=nil{
		log.Fatal(err)
	}
	
	username := user["username"]
	password :=  user["password"]
	db,err:=database.DB_connection()
	if err!=nil{
        Response(&w,"Authentication failed",http.StatusInternalServerError)
		return
	}
    
	//get the account password from the database 
	var hash string 
	var  accountID int
	Query:=`SELECT password,ID FROM account WHERE account.username=$1` 
    err=db.QueryRow(Query,username).Scan(&hash,&accountID)
	if err != nil {
		return
	}
    
	controller//Error handle for password Comaparison
	if !CheckPasswordHash(password,hash){
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Password Doesn't match"})
		return
	}
	
	expirationTime := time.Now().Add(10 * time.Hours)
	claims := jwt.MapClaims{
        "account_id": accountID,
		"username": username
		"exp":expirationTime.Unix(),
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	tokenString, err := token.SignedString([]byte("authkey"))
	if err != nil  {
		Response(&w,"INTERNAL ERROR",http.StatusInternalServerError)
		return
	}
	c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
	Response(&w,"INTERNAL ERROR",http.StatusInternalServerError)
	c.JSON(http.StatusOK, gin.H{"token": tokenString})
}

	
func CheckPasswordHash(password string , hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
