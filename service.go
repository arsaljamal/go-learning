package go_learning

import (
	"errors"
	"log"
)

// interface which provides a set of methods

var ErrNotFound = errors.New("not found")

func init() {
	log.SetPrefix("Trace: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Println("service started")
}

type User struct {
	ID   int64
	NAME string
}

type Service interface {
	GetUser(id int64) (User, error)
	GetUsers(ids []int64) (map[int64]User, error)
}
