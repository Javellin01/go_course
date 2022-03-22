package usecase

import (
	"context"
	"github.com/Javellin01/go_course/internal/data/repository"
	"github.com/Javellin01/go_course/internal/domain/agg"
	"github.com/Javellin01/go_course/internal/domain/builder"
	"github.com/Javellin01/go_course/internal/domain/dto"
	"time"
)

type AdvertiserUsecase struct {
	ctx                  context.Context
	advertiserRepository repository.Advertiser
}

func NewAdvertiserUsecase(ctx context.Context, repo repository.Advertiser) Advertiser {
	return AdvertiserUsecase{
		ctx:                  ctx,
		advertiserRepository: repo,
	}
}

func (au AdvertiserUsecase) Create(dto dto.AdvertiseRequest) (string, error) {
	advertiser := builder.NewAdvertiserBuilder().BuildFromRequest(dto)
	ctx, cancel := context.WithTimeout(au.ctx, time.Second*3)
	defer cancel()

	return au.advertiserRepository.Create(ctx, advertiser)
}

func (au AdvertiserUsecase) Find(id string) (agg.Advertiser, error) {
	ctx, cancel := context.WithTimeout(au.ctx, time.Second*3)
	defer cancel()

	return au.advertiserRepository.Find(ctx, id)
}
