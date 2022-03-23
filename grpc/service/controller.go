package main

import (
	"context"
	"github.com/go_learning"
	grpcsvc "github.com/go_learning/grpc"
	"log"
)

// userServiceController implements the gRPC UserServiceServer interface.
type userServiceController struct {
	userService go_learning.Service
}

// NewUserServiceController instantiates a new UserServiceServer.
func NewUserServiceController(userService go_learning.Service) grpcsvc.UserServiceServer {
	return &userServiceController{
		userService: userService,
	}
}

// GetUsers calls the core service's GetUsers method and maps the result to a grpc service response.
func (u *userServiceController) GetUsers(ctx context.Context, request *grpcsvc.GetUsersRequest) (response *grpcsvc.GetUsersResponse, err error) {
	resultMap, err := u.userService.GetUsers(request.GetIds())

	if err != nil {
		return
	}

	response = &grpcsvc.GetUsersResponse{}

	for _, u := range resultMap {
		response.Users = append(response.Users, marshalUser(&u))
	}

	log.Println("handled GetUsers(%v)\n", request.GetIds())
	return
}

// marshalUser marshals a business object User into a gRPC layer User.
func marshalUser(u *go_learning.User) *grpcsvc.User {
	return &grpcsvc.User{Id: u.ID, Name: u.NAME}
}
