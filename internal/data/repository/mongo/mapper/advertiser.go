package mapper

import (
	"github.com/Javellin01/go_course/internal/domain/agg"
	"github.com/Javellin01/go_course/internal/domain/entity"
	"github.com/Javellin01/go_course/internal/domain/vo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Advertiser struct {
}

func NewAdvertiserMapper() Advertiser {
	return Advertiser{}
}

func (a Advertiser) BuildObjectCreate(advertiser agg.Advertiser) bson.M {
	return bson.M{
		"name":    advertiser.Name,
		"balance": advertiser.Balance,
	}
}

func (a Advertiser) MapSingle(r *mongo.SingleResult) (agg.Advertiser, error) {
	var advertiser entity.Advertiser
	var timestamp vo.Timestamp

	if err := r.Decode(&advertiser); err != nil {
		return agg.Advertiser{}, err
	}

	if err := r.Decode(&timestamp); err != nil {
		return agg.Advertiser{}, err
	}

	return agg.Advertiser{
		Advertiser: advertiser,
		Timestamp:  vo.Timestamp{},
	}, nil
}
