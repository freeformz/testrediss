package main

import (
	"crypto/tls"
	"fmt"
	"os"

	"github.com/garyburd/redigo/redis"
	"github.com/heroku/x/hredis"
)

func testTLSSkipVerify() error {
	fmt.Println("testTLSSkipVerify")
	u, err := hredis.RedissURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return err
	}
	fmt.Println("url:", u)
	c, err := redis.DialURL(u, redis.DialTLSSkipVerify(true))
	if err != nil {
		return err
	}
	defer c.Close()
	if err := c.Send("SET", "foo", "bar"); err != nil {
		return err
	}
	if err := c.Flush(); err != nil {
		return err
	}
	r, err := c.Receive()
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func testCustomTLSConfig() error {
	fmt.Println("testCustomTLSConfig")
	u, err := hredis.RedissURL(os.Getenv("REDIS_URL"))
	if err != nil {
		return err
	}
	fmt.Println("url:", u)
	c, err := redis.DialURL(u, redis.DialTLSConfig(&tls.Config{InsecureSkipVerify: true}))
	if err != nil {
		return err
	}
	defer c.Close()
	if err := c.Send("SET", "foo", "bar"); err != nil {
		return err
	}
	if err := c.Flush(); err != nil {
		return err
	}
	r, err := c.Receive()
	if err != nil {
		return err
	}
	fmt.Println(r)
	return nil
}

func main() {
	fmt.Println(testTLSSkipVerify())
	fmt.Println(testCustomTLSConfig())
}
