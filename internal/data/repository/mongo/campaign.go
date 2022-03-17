package mongorepository

import (
	"context"
	"errors"
	"github.com/Javellin01/go_course/internal/data/repository/mongo/mapper"
	"github.com/Javellin01/go_course/internal/domain/agg"
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
	obj := mapper.NewCampaignMapper().BuildObjectCreate(campaign)

	localCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	result, err := c.db.InsertOne(localCtx, obj)
	if err != nil {
		return "", err
	}

	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return id.Hex(), nil
	}

	return "", errors.New("returned result is not primitive.ObjectID")
}

func (c Campaign) Find(ctx context.Context, id string) (agg.Campaign, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return agg.Campaign{}, errors.New("invalid campaign id")
	}

	localCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	res := c.db.FindOne(localCtx, bson.M{"_id": objectID})

	result, err := mapper.NewCampaignMapper().MapSingle(res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return agg.Campaign{}, errors.New("document not found")
		}

		return agg.Campaign{}, err
	}

	return result, nil
}

func (c Campaign) FindBy(ctx context.Context, filter bson.M) ([]agg.Campaign, error) {
	localCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	cursor, err := c.db.Find(localCtx, filter)
	if err != nil {
		return []agg.Campaign{}, err
	}

	if cursor.RemainingBatchLength() == 0 {
		return []agg.Campaign{}, errors.New("documents not found")
	}

	result, err := mapper.NewCampaignMapper().MapSlice(localCtx, cursor)
	if err != nil {
		return []agg.Campaign{}, err
	}

	return result, nil
}

func (c Campaign) FindAll(ctx context.Context) ([]agg.Campaign, error) {
	localCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	cursor, err := c.db.Find(localCtx, bson.M{})
	if err != nil {
		return []agg.Campaign{}, err
	}

	if cursor.RemainingBatchLength() == 0 {
		return []agg.Campaign{}, errors.New("documents not found")
	}

	result, err := mapper.NewCampaignMapper().MapSlice(localCtx, cursor)
	if err != nil {
		return []agg.Campaign{}, err
	}

	return result, nil
}

func (c Campaign) Update(ctx context.Context, campaign agg.Campaign) error {
	objectID, err := primitive.ObjectIDFromHex(campaign.ID)
	if err != nil {
		return errors.New("invalid campaign id")
	}

	obj := mapper.NewCampaignMapper().BuildObjectCreate(campaign)

	localCtx, cancel := context.WithTimeout(ctx, c.timeout)
	defer cancel()

	filter := bson.M{"_id": objectID}
	result, err := c.db.UpdateOne(localCtx, filter, obj)
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
