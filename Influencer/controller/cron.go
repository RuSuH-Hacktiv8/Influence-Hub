package controller

import (
	"influence-hub-influencer/middleware"
	"log"
)

func (cn *Controller) UpdateFollowerCount() {
	influencers, err := cn.Controller.FindAll()
	if err != nil {
		log.Printf("Unable to get data from db: %s\n", err.Error())
	}
	for _, influencer := range influencers {
		count, err := middleware.GetInstagramFollowers(influencer.InstagramUsername)
		if err != nil {
			log.Printf("Unable to get follower count: %s\n", err.Error())
			continue
		}

		influencer.InstagramFollowers = count
		err = cn.Controller.UpdateFollowerCount(influencer)
		if err != nil {
			log.Printf("Unable to update follower count: %s\n", err.Error())
		}
	}
	log.Printf("Updated %v influencers followers count\n", len(influencers))
}
