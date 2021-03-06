package memory

import (
	"log"
	"saver/config"

	"github.com/gomodule/redigo/redis"
)

func New(config config.Redis) redis.Conn {
	conn, err := redis.Dial("tcp", config.Host+":"+config.Port)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}
