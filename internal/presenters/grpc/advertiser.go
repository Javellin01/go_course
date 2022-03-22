package grpchandler

import (
	"context"
	"github.com/Javellin01/go_course/internal/domain/dto"
	"github.com/Javellin01/go_course/internal/presenters/grpc/pb"
)

func (s PlatformsServer) CreateAdvertiser(ctx context.Context, advertiser *pb.Advertiser) (*pb.AdvertiserID, error) {
	advertiserDto := dto.AdvertiseRequest{
		Name:    advertiser.Name,
		Balance: advertiser.Balance,
	}

	createdId, err := s.usecase.Advertiser.Create(advertiserDto)
	if err != nil {
		return nil, err
	}

	return &pb.AdvertiserID{Id: createdId}, nil
}
