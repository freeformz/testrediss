package main

import (
	"fmt"
	"os"

	"github.com/freeformz/testrediss"
	"gopkg.in/redis.v5"
)

func main() {
	u, err := testrediss.RedissURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	o, err := redis.ParseURL(u)
	if err != nil {
		panic(err)
	}
	c := redis.NewClient(o)
	pong, err := c.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
}
