package builder

import (
	"github.com/Javellin01/go_course/internal/domain/agg"
	"github.com/Javellin01/go_course/internal/domain/dto"
)

type Campaign interface {
	BuildFromRequest(campaign dto.CampaignRequest) agg.Campaign
}

type Advertiser interface {
	BuildFromRequest(advertiser dto.AdvertiseRequest) agg.Advertiser
}
