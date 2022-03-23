package core

import "github.com/go_learning"

// implementation of service.go
type service struct {
	//database dependency will go here, we are using map for now
	m map[int64]go_learning.User
}

// NewService instantiates a new Service.
func NewService( /* a database connection would be injected here */ ) go_learning.Service {
	return &service{
		m: map[int64]go_learning.User{
			1: {ID: 1, NAME: "Alice"},
			2: {ID: 2, NAME: "Bob"},
			3: {ID: 3, NAME: "Carol"},
		},
	}
}

func (s *service) GetUser(id int64) (result go_learning.User, err error) {
	if result, ok := s.m[id]; ok {
		return result, nil
	}

	return result, go_learning.ErrNotFound
}

func (s *service) GetUsers(ids []int64) (result map[int64]go_learning.User, err error) {
	result = map[int64]go_learning.User{}

	for _, id := range ids {
		if u, ok := s.m[id]; ok {
			result[id] = u
		}
	}
	return result, nil
}
