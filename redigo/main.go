package main

import (
	"fmt"
	"os"

	"github.com/freeformz/testrediss"
	"github.com/garyburd/redigo/redis"
)

func main() {
	u, err := testrediss.RedissURL(os.Getenv("REDIS_URL"))
	if err != nil {
		panic(err)
	}
	fmt.Println("url:", u)
	c, err := redis.DialURL(u, redis.DialTLSSkipVerify(true))
	if err != nil {
		panic(err)
	}
	if err := c.Send("SET", "foo", "bar"); err != nil {
		panic(err)
	}
	if err := c.Flush(); err != nil {
		panic(err)
	}
	r, err := c.Receive()
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
