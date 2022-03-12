package usecase

import (
	"github.com/Javellin01/go_course/internal/domain/agg"
	"github.com/Javellin01/go_course/internal/domain/dto"
)

type Usecase struct {
	Campaign Campaign
}

type Campaign interface {
	Create(campaign dto.Campaign) (string, error)
	Find(id string) (agg.Campaign, error)
}
