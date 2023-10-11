package controller

import "influence-hub-influencer/repository"

type Controller struct {
	Controller repository.Repository
}

func NewController(repo *repository.Repository) *Controller {
	return &Controller{Controller: *repo}
}
