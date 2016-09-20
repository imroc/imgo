package main

import (
	"fmt"
	"testing"

	"github.com/beego/redigo/redis"
)

func TestRedisGet(t *testing.T) {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("连接redis失败")
		return
	}
	defer conn.Close()
	reply, err := redis.Int64(conn.Do("HGET", "15984419985", "token"))
	fmt.Printf("%#v-%v", reply, err)
}
