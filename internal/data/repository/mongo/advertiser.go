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

type Advertiser struct {
	db      *mongo.Collection
	timeout time.Duration
}

func NewAdvertiserRepository(db *mongo.Database, timeout time.Duration) Advertiser {
	return Advertiser{
		db:      db.Collection("advertiser"),
		timeout: timeout,
	}
}

func (a Advertiser) Create(ctx context.Context, advertiser agg.Advertiser) (string, error) {
	obj := mapper.NewAdvertiserMapper().BuildObjectCreate(advertiser)

	localCtx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	result, err := a.db.InsertOne(localCtx, obj)
	if err != nil {
		return "", err
	}

	if id, ok := result.InsertedID.(primitive.ObjectID); ok {
		return id.Hex(), nil
	}

	return "", errors.New("returned result is not primitive.ObjectID")
}

func (a Advertiser) Find(ctx context.Context, id string) (agg.Advertiser, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return agg.Advertiser{}, errors.New("invalid campaign id")
	}

	localCtx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	res := a.db.FindOne(localCtx, bson.M{"_id": objectID})

	result, err := mapper.NewAdvertiserMapper().MapSingle(res)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return agg.Advertiser{}, errors.New("document not found")
		}

		return agg.Advertiser{}, err
	}

	return result, nil
}

func (a Advertiser) FindAll(ctx context.Context) ([]agg.Advertiser, error) {
	//TODO implement me
	panic("implement me")
}

func (a Advertiser) FindBy(ctx context.Context, m bson.M) ([]agg.Advertiser, error) {
	//TODO implement me
	panic("implement me")
}

func (a Advertiser) Update(ctx context.Context, advertiser agg.Advertiser) error {
	//TODO implement me
	panic("implement me")
}

func (a Advertiser) Delete(ctx context.Context, s string) error {
	//TODO implement me
	panic("implement me")
}
