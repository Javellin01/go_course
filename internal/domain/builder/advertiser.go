package builder

import (
	"github.com/Javellin01/go_course/internal/domain/agg"
	"github.com/Javellin01/go_course/internal/domain/dto"
	"github.com/Javellin01/go_course/internal/domain/entity"
	"github.com/Javellin01/go_course/internal/domain/vo"
	"time"
)

type AdvertiserBuilder struct {
}

func NewAdvertiserBuilder() Advertiser {
	return AdvertiserBuilder{}
}

func (ab AdvertiserBuilder) BuildFromRequest(dto dto.AdvertiseRequest) agg.Advertiser {
	advertiser := entity.Advertiser{
		ID:      dto.ID,
		Name:    dto.Name,
		Balance: dto.Balance,
	}

	return agg.Advertiser{
		Advertiser: advertiser,
		Timestamp: vo.Timestamp{
			CreatedAt: time.Now(),
		},
	}
}
