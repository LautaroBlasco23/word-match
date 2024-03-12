package handlers

import (
	"context"
	"fmt"

	"match-words/initializers"
)

func Login(username string, password string) {
	client := initializers.ConnectWithDatabase()
	ctx := context.Background()

	err := client.Set(ctx, "foo", "bar", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get(ctx, "foo").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("foo", val)

}

func Register(email string, username string, password string) {

}

func VerifyToken(email string, username string, password string) {

}
