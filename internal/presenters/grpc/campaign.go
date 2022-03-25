package grpchandler

import (
	"context"
	"github.com/Javellin01/go_course/internal/domain/dto"
	"github.com/Javellin01/go_course/internal/presenters/grpc/pb"
)

func (s PlatformsServer) CreateCampaign(ctx context.Context, campaign *pb.CampaignRequest) (*pb.CampaignID, error) {
	campaignDto := dto.CampaignRequest{
		Name:         campaign.Name,
		AdvertiserId: campaign.AdvertiserId,
		SpentItems:   make([]dto.SpentItem, len(campaign.SpentItems)),
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

func (s PlatformsServer) FindCampaign(ctx context.Context, id *pb.CampaignID) (*pb.CampaignResponse, error) {
	c, err := s.usecase.Campaign.Find(id.Id)
	if err != nil {
		return nil, err
	}
	response := &pb.CampaignResponse{
		Id:   &pb.CampaignID{Id: c.Campaign.ID},
		Name: c.Campaign.Name,
	}
	return response, nil
}
