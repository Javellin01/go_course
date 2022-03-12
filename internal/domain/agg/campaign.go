package agg

import (
	"github.com/Javellin01/go_course/internal/domain/entity"
	"github.com/Javellin01/go_course/internal/domain/vo"
)

type Campaign struct {
	entity.Campaign
	vo.Timestamp
}
