package rediscache

import (
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func initPool() {
	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial("tcp", "localhost:6379")
			if err != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			return conn, err
		},
	}
}

func ping(conn redis.Conn) {
	_, err := redis.String(conn.Do("PING"))
	if err != nil {
		log.Fatalf("ERROR: failed to ping redis connection: %s", err.Error())
	}
}
