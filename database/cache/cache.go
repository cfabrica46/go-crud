package cache

import (
	"fmt"
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

func Close() (err error) {
	err = db.Close()
	if err != nil {
		return
	}
	return
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

func TokenIsValid(token string) (check bool, err error) {
	result, err := db.Get(token).Result()
	if err != nil {
		if err.Error() == redis.Nil.Error() {
			err = nil
			return
		}
		return
	}
	fmt.Println(result)

	if result == "1" {
		check = true
	}
	return
}
