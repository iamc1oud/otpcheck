package redis

import (
	"encoding/json"
	"fmt"
	"time"

	redigo "github.com/gomodule/redigo/redis"
)

type Redis struct {
	pool *redigo.Pool
	conf Conf
}

type Conf struct {
	Host      string        `json:"host"`
	Port      int           `json:"port"`
	Username  string        `json:"username"`
	Password  string        `json:"password"`
	DB        int           `json:"db"`
	MaxActive int           `json:"max_active"`
	MaxIdle   int           `json:"max_idle"`
	Timeout   time.Duration `json:"timeout"`
	KeyPrefix string        `json:"key_prefix"`

	// If this is set, 'check' and 'close' events will be PUBLISHed to
	// to this Redis key (Redis PubSub).
	PublishKey string `json:"publish_key"`
}

type event struct {
	Type      string          `json:"type"`
	Namespace string          `json:"namespace"`
	ID        string          `json:"id"`
	Data      json.RawMessage `json:"data"`
}

// New returns a Redis implementation of store.
func New(c Conf) *Redis {
	if c.KeyPrefix == "" {
		c.KeyPrefix = "OTP"
	}

	pool := &redigo.Pool{
		Wait:      true,
		MaxActive: c.MaxActive,
		MaxIdle:   c.MaxIdle,
		Dial: func() (redigo.Conn, error) {
			c, err := redigo.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", c.Host, c.Port),
				redigo.DialPassword(c.Password),
				redigo.DialConnectTimeout(c.Timeout),
				redigo.DialReadTimeout(c.Timeout),
				redigo.DialWriteTimeout(c.Timeout),
				redigo.DialDatabase(c.DB),
			)

			return c, err
		},
	}

	return &Redis{
		conf: c,
		pool: pool,
	}
}

func (r *Redis) Ping() error {
	c := r.pool.Get()
	defer c.Close()
	_, err := c.Do("PING")
	return err
}

// Check checks the attempt count and TTL duration against an ID.
// Passing count=true increments the attempt counter.
func (r *Redis) Check(namespace, id string, counter bool) (models.OTP, error) {
	
}
