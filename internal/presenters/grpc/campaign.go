package grpchandler

import (
	"context"
	"github.com/Javellin01/go_course/internal/domain/dto"
	"github.com/Javellin01/go_course/internal/presenters/grpc/pb"
)

func (s PlatformsServer) CreateCampaign(ctx context.Context, campaign *pb.Campaign) (*pb.CampaignID, error) {
	campaignDto := dto.Campaign{
		Name:       campaign.Name,
		SpentItems: make([]dto.SpentItem, len(campaign.SpentItems)),
	}

	for i, item := range campaign.SpentItems {
		campaignDto.SpentItems[i] = dto.SpentItem{
			TotalSpent:       int(item.TotalSpent),
			Cost:             int(item.Cost),
			Profit:           int(item.Profit),
			AgencyProfit:     float64(item.AgencyProfit),
			Charge:           float64(item.Charge),
			AdvertiserCharge: float64(item.AdvertiserCharge),
			AgencyCharge:     float64(item.AgencyCharge),
			TotalCost:        int(item.TotalCost),
		}
	}

	createdId, err := s.usecase.Campaign.Create(campaignDto)
	if err != nil {
		return nil, err
	}

	return &pb.CampaignID{Id: createdId}, nil
}
