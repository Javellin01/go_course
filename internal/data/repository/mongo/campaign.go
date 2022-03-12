package mongorepository

import (
	"context"
	"errors"
	"github.com/Javellin01/go_course/internal/domain/agg"
	"github.com/Javellin01/go_course/internal/domain/entity"
	"github.com/Javellin01/go_course/internal/domain/vo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type Campaign struct {
	db      *mongo.Collection
	timeout time.Duration
}

func NewCampaignRepository(db *mongo.Database, timeout time.Duration) Campaign {
	return Campaign{
		db:      db.Collection("campaign"),
		timeout: timeout,
	}
}

func (c Campaign) Create(ctx context.Context, campaign agg.Campaign) (string, error) {
	doc := bson.M{
		"name":       campaign.Name,
		"spentItems": campaign.SpentItems,
		"createdAt":  campaign.CreatedAt,
	}

	localCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	result, err := c.db.InsertOne(localCtx, doc)
	if err != nil {
		return "", err
	}

	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return id.Hex(), nil
	}

	return "", errors.New("returned result is not primitive.ObjectID")
}

func (c Campaign) Find(ctx context.Context, id string) (agg.Campaign, error) {
	documentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return agg.Campaign{}, errors.New("invalid campaign id")
	}

	localCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	var campaign entity.Campaign
	var timestamp vo.Timestamp

	res := c.db.FindOne(localCtx, bson.M{"_id": documentId})

	if err := res.Decode(&campaign); err != nil {
		if err == mongo.ErrNoDocuments {
			return agg.Campaign{}, errors.New("document not found")
		}

		return agg.Campaign{}, err
	}

	if err := res.Decode(&timestamp); err != nil {
		return agg.Campaign{}, err
	}

	result := agg.Campaign{
		Campaign:  campaign,
		Timestamp: timestamp,
	}

	return result, nil
}

func (c Campaign) FindBy(ctx context.Context, filter bson.M) ([]agg.Campaign, error) {
	localCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	var result []agg.Campaign

	cursor, err := c.db.Find(localCtx, filter)
	if err != nil {
		return result, err
	}

	if cursor.RemainingBatchLength() == 0 {
		return []agg.Campaign{}, errors.New("documents not found")
	}

	for cursor.Next(localCtx) {
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

	return result, err
}

func (c Campaign) FindAll(ctx context.Context) ([]agg.Campaign, error) {
	localCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	var result []agg.Campaign

	cursor, err := c.db.Find(localCtx, bson.M{})
	if err != nil {
		return result, err
	}

	if cursor.RemainingBatchLength() == 0 {
		return []agg.Campaign{}, errors.New("documents not found")
	}

	for cursor.Next(localCtx) {
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

	return result, err
}

func (c Campaign) Update(ctx context.Context, campaign agg.Campaign) error {
	documentId, err := primitive.ObjectIDFromHex(campaign.ID)
	if err != nil {
		return errors.New("invalid campaign id")
	}

	doc := bson.M{
		"$set": bson.M{
			"name":       campaign.Name,
			"spentItems": campaign.SpentItems,
			"updatedAt":  time.Now(),
		},
	}

	localCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	filter := bson.M{"_id": documentId}
	result, err := c.db.UpdateOne(localCtx, filter, doc)
	if err != nil {
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("nothing was updated")
	}

	return nil
}

func (c Campaign) Delete(ctx context.Context, id string) error {
	documentId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid campaign id")
	}

	localCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	_, err = c.db.DeleteOne(localCtx, bson.M{"_id": documentId})

	if err != nil {
		return err
	}

	return nil
}
