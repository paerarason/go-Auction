package reddis
import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

cert, err := tls.LoadX509KeyPair("redis_user.crt", "redis_user_private.key")
if err != nil {
    log.Fatal(err)
}

// Load CA cert
caCert, err := os.ReadFile("redis_ca.pem")
if err != nil {
    log.Fatal(err)
}
caCertPool := x509.NewCertPool()
caCertPool.AppendCertsFromPEM(caCert)

client := redis.NewClient(&redis.Options{
        Addr:     "your_redis_host:your_redis_port", // Replace with your Redis server's address
        Password: "your_password_here",              // Replace with your Redis password
        DB:       0,                                 // Use default DB or specify the desired database number
    })
//send SET command
err = client.Set(ctx, "foo", "bar", 0).Err()
if err != nil {
    panic(err)
}

//send GET command and print the value
val, err := client.Get(ctx, "foo").Result()
if err != nil {
    panic(err)
}

fmt.Println("foo", val)