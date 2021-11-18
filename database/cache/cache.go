package cache

import (
	"time"

	"github.com/go-redis/redis"
)

var db *redis.Client

func Open() {
	options := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	db = redis.NewClient(options)
}

func SetToken(token string) (err error) {
	index, err := getTokenHash(token)
	if err != nil {
		return
	}

	err = db.Set(index, token, time.Minute*10).Err()
	if err != nil {
		return
	}
	return
}

func CheckIfTokenIsInTheBlackList(token string) (check bool, err error) {
	index, err := getTokenHash(token)
	if err != nil {
		return
	}

	result, err := db.Get(index).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			err = nil
			return
		}
		return
	}

	if result == token {
		check = true
	}

	return
}