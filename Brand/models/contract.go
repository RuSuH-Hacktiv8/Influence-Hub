package models

type Contract struct {
	ID           uint `json:"id,omitempty" bson:"_id,omitempty"`
	InfluencerID int  `json:"influencer_id,omitempty" bson:"influencer_id,omitempty"`
	CampaignID   int  `json:"campaign_id,omitempty" bson:"campaign_id,omitempty"`
}


