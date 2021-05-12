package database

import (
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

var client *redis.Client

func InitRedis() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func RegisterUser(username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	return client.Set(username, hash, 0).Err()
}

func LoginUser(username, password string) error {
	hash, err := client.Get(username).Bytes()
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword(hash, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func AddNote(note string) error {
	return client.LPush("Notes", note).Err()
}

func DeleteNote(note string) error {
	return client.LRem("Notes", 1, note).Err()
}

func GetNotes()([]string, error) {
	return client.LRange("Notes", 0, 100000).Result()
}
