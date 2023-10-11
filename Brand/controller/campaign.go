package controller

import "influence-hub-brand/repository"

type CampaignController struct {
	Repo repository.Repository
}

func NewCampaignController(repo repository.Repository) CampaignController {
	return CampaignController{repo}
}
