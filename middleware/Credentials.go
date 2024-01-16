package middleware
import (
	"net/http"
	"context"
  "fmt"
  "go.mongodb.org/mongo-driver/bson"
  "go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

func CredentialMiddleware(next http.Handler) http.Handler{
	
	// database.Collection("newCollection").InsertOne(ctx, bson.M{"key": "value"})
  next.ServeHTTP()
}


