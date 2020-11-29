package db

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func getConnectionConfig() redis.Options {
	host := os.Getenv("REDIS_HOST")

	if host == "" {
		panic("Can't get Redist host")
	}

	port := os.Getenv("REDIS_PORT")

	if port == "" {
		panic("Can't get Redist port")
	}

	addr := fmt.Sprintf("%v:%v", host, port)

	return redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	}
}

func extractPost(s string, err error) (Post, error) {
	p := &Post{}

	if err != nil {
		return *p, err
	}

	e := json.Unmarshal([]byte(s), &p)

	if e != nil {
		return *p, err
	}

	return *p, nil
}

// GetPersistentPost return is post exist and store to redis if not
func GetPersistentPost(p Post) (post Post, isNew bool, err error) {
	config := getConnectionConfig()
	rdb := redis.NewClient(&config)

	if value, err := rdb.Get(ctx, p.ID).Result(); value != "" {
		d, err := extractPost(value, err)

		return d, false, err
	}

	s, err := json.Marshal(&p)

	if err != nil {
		return p, true, err
	}

	err = rdb.Set(ctx, p.ID, string(s), 0).Err()

	return p, true, err
}
