package cache

import (
	"time"

	"github.com/go-redis/redis"
)

var db *redis.Client

func init() {
	Open()
}

func Open() {
	options := &redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}

	db = redis.NewClient(options)
}

func SetToken(token string) (err error) {
	err = db.Set(token, true, time.Minute*10).Err()
	if err != nil {
		return
	}
	return
}

func DeleteTokenUsingValue(token string) (err error) {
	return db.Del(token).Err()
}

func CheckIfTokenItsValid(token string) (check bool, err error) {
	result, err := db.Get(token).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			err = nil
			return
		}
		return
	}

	if result {
		check = true
	}
	return
}
