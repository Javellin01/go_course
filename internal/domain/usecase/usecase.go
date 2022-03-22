package usecase

import (
	"github.com/Javellin01/go_course/internal/domain/agg"
	"github.com/Javellin01/go_course/internal/domain/dto"
)

type Usecase struct {
	Campaign   Campaign
	Advertiser Advertiser
}

type Campaign interface {
	Create(campaign dto.CampaignRequest) (string, error)
	Find(id string) (agg.Campaign, error)
}

type Advertiser interface {
	Create(advertiser dto.AdvertiseRequest) (string, error)
	Find(id string) (agg.Advertiser, error)
}
