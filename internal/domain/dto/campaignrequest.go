package dto

type CampaignRequest struct {
	ID           string      `json:"id"`
	Name         string      `json:"name"`
	AdvertiserId string      `json:"advertiserId"`
	SpentItems   []SpentItem `json:"spentItems"`
}
