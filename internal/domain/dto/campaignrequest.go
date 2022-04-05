package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CampaignRequest struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	AdvertiserId primitive.ObjectID `json:"advertiserId"`
	SpentItems   []SpentItem        `json:"spentItems"`
}
