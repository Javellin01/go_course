package builder

import (
	"github.com/Javellin01/go_course/internal/domain/agg"
	"github.com/Javellin01/go_course/internal/domain/dto"
	"github.com/Javellin01/go_course/internal/domain/entity"
	"github.com/Javellin01/go_course/internal/domain/vo"
	"time"
)

type CampaignBuilder struct{}

func NewCampaignBuilder() Campaign {
	return CampaignBuilder{}
}

func (cb CampaignBuilder) BuildFromRequest(dto dto.CampaignRequest) agg.Campaign {
	campaign := entity.Campaign{
		ID:         dto.ID,
		Name:       dto.Name,
		SpentItems: make([]entity.SpentItem, len(dto.SpentItems)),
	}

	for i, item := range dto.SpentItems {
		campaign.SpentItems[i] = entity.SpentItem{
			TotalSpent:       item.TotalSpent,
			Cost:             item.Cost,
			Profit:           item.Profit,
			AgencyProfit:     item.AgencyProfit,
			Charge:           item.Charge,
			AdvertiserCharge: item.AdvertiserCharge,
			AgencyCharge:     item.AgencyCharge,
			TotalCost:        item.TotalCost,
		}
	}
	t := agg.Campaign{
		Campaign:     campaign,
		AdvertiserId: dto.AdvertiserId,
		Timestamp: vo.Timestamp{
			CreatedAt: time.Now(),
		},
	}
	_ = t

	return agg.Campaign{
		Campaign:     campaign,
		AdvertiserId: dto.AdvertiserId,
		Timestamp: vo.Timestamp{
			CreatedAt: time.Now(),
		},
	}
}
