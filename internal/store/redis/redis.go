package redis

import (
	redigo "github.com/gomodule/redigo/redis"
)

type Redis struct {
	pool *redigo.Pool
	conf Conf
}