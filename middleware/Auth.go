package middleware
import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"golang.org/x/crypto/bcrypt"
	"time"
	//"fmt"
	"log"
	//"os"
)

type Claims struct {
	AccountID int `json:"account_id"`
	jwt.StandardClaims
}

func JWTokenMiddlerware(next http.Handler) http.Handler{
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return []byte("authkey"), nil // Use your secret key here
        })
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		c.Abort()
		return
	}

	 if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
         c.Set("account_id", claims["account_id"])
	     c.Next()
	 }
}



func GenerateToken() {
	username := c.PostForm("username")
	password := c.PostForm("password")
	
	db,err:=database.DB_connection()
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Authentication failed"})
			return
	}
     
   log.Println(username,password)
    //get the account password from the database 
	var hash string 
	var  accountID int
	Query:=`SELECT password,ID FROM account WHERE account.username=$1` 
    err=db.QueryRow(Query,username).Scan(&hash,&accountID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to get Account details"})
		return
	}
    
	//Error handle for password Comaparison
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
		c.SetCookie("token", tokenString, 3600, "/", "localhost", false, true)
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}

	
func CheckPasswordHash(password string , hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}