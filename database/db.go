package database

import (
	"github.com/go-redis/redis"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password []byte
}

type Note struct {
	Name string
}

type DatabaseService interface {
	RegisterUser(username, password string) error
	LoginUser(username, password string) error
	AddNote(note string) error
	GetNotes() ([]string, error)
	DeleteNote(note string) error
}


func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return client
}

func (u User) RegisterUser(username, password string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	redis := InitRedis()
	u.Username = username
	u.Password = hash
	return redis.Set(u.Username, u.Password, 0).Err()
}

func (u User) LoginUser(username, password string) error {
	redis := InitRedis() 
	hash, err := redis.Get(username).Bytes()
	if err != nil {
		return err
	}
	u.Username = username
	u.Password = hash
	err = bcrypt.CompareHashAndPassword(u.Password, []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func (n Note) AddNote(note string) error {
	redis := InitRedis()
	n.Name = note
	return redis.LPush("Notes", note).Err()
}

func (n Note) GetNotes() ([]string, error) {
	redis := InitRedis()
	return redis.LRange("Notes", 0, 100).Result()
}

func (n Note) DeleteNote(note string) error {
	redis := InitRedis()
	return redis.LRem("Notes", 1, note).Err()
}
