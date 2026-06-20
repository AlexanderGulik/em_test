package controller

import (
	"em_test/src/subscription/service"
)

type SubController struct {
	service *service.SubService
}

func NewSubController() *SubController {
	return &SubController{service: service.NewSubscriptionsService()}
}






