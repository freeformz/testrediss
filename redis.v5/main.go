package main

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/heroku/x/hredis"
	"gopkg.in/redis.v5"
)

func main() {
	u, err := hredis.RedissURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	fmt.Println("url:", u)

	o, err := redis.ParseURL(u)
	if err != nil {
		panic(err)
	}
	o.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	c := redis.NewClient(o)
	pong, err := c.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
}
