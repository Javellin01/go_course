package agg

import (
	"github.com/Javellin01/go_course/internal/domain/entity"
	"github.com/Javellin01/go_course/internal/domain/vo"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Campaign struct {
	entity.Campaign `bson:",inline"`
	AdvertiserId    primitive.ObjectID `bson:"advertiserId"`
	vo.Timestamp    `bson:",inline"`
}
