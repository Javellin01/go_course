package usecase

import (
	"context"
	"github.com/Javellin01/go_course/internal/data/repository"
	"github.com/Javellin01/go_course/internal/domain/agg"
	"github.com/Javellin01/go_course/internal/domain/builder"
	"github.com/Javellin01/go_course/internal/domain/dto"
	"time"
)

type CampaignUsecase struct {
	ctx                context.Context
	campaignRepository repository.Campaign
}

func NewCampaignUsecase(ctx context.Context, repo repository.Campaign) Campaign {
	return CampaignUsecase{
		ctx:                ctx,
		campaignRepository: repo,
	}
}

func (cu CampaignUsecase) Create(dto dto.CampaignRequest) (string, error) {
	campaign := builder.NewCampaignBuilder().BuildFromRequest(dto)
	ctx, cancel := context.WithTimeout(cu.ctx, time.Second*3)
	defer cancel()

	return cu.campaignRepository.Create(ctx, campaign)
}

func (cu CampaignUsecase) Find(id string) (agg.Campaign, error) {
	ctx, cancel := context.WithTimeout(cu.ctx, time.Second*3)
	defer cancel()

	return cu.campaignRepository.Find(ctx, id)
}
