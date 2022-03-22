package mapper

import (
	"context"
	"github.com/Javellin01/go_course/internal/domain/agg"
	"github.com/Javellin01/go_course/internal/domain/entity"
	"github.com/Javellin01/go_course/internal/domain/vo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Campaign struct {
}

func NewCampaignMapper() Campaign {
	return Campaign{}
}

func (c Campaign) BuildObjectCreate(campaign agg.Campaign) bson.M {
	return bson.M{
		"name":         campaign.Name,
		"spentItems":   campaign.SpentItems,
		"advertiserId": campaign.AdvertiserId,
		"createdAt":    campaign.CreatedAt,
	}
}

func (c Campaign) BuildObjectUpdate(campaign agg.Campaign) bson.M {
	return bson.M{
		"$set": bson.M{
			"name":       campaign.Name,
			"spentItems": campaign.SpentItems,
			"updatedAt":  time.Now(),
		},
	}
}

func (c Campaign) MapSingle(r *mongo.SingleResult) (agg.Campaign, error) {
	var aggregate agg.Campaign
	var campaign entity.Campaign
	var timestamp vo.Timestamp

	if err := r.Decode(&aggregate); err != nil {
		return agg.Campaign{}, err
	}
	if err := r.Decode(&campaign); err != nil {
		return agg.Campaign{}, err
	}
	if err := r.Decode(&timestamp); err != nil {
		return agg.Campaign{}, err
	}

	aggregate.Campaign = campaign
	aggregate.Timestamp = timestamp

	return aggregate, nil
}

func (c Campaign) MapSlice(ctx context.Context, cursor *mongo.Cursor) ([]agg.Campaign, error) {
	var result []agg.Campaign

	for cursor.Next(ctx) {
		var campaign entity.Campaign
		var timestamp vo.Timestamp

		err := cursor.Decode(&campaign)
		if err != nil {
			return result, err
		}

		err = cursor.Decode(&timestamp)
		if err != nil {
			return result, err
		}

		resItem := agg.Campaign{
			Campaign:  campaign,
			Timestamp: timestamp,
		}

		result = append(result, resItem)
	}

	return result, nil
}
