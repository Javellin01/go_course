package repository

import (
	"context"
	"github.com/Javellin01/go_course/internal/domain/agg"
	"go.mongodb.org/mongo-driver/bson"
)

type Campaign interface {
	Create(context.Context, agg.Campaign) (string, error)
	Find(context.Context, string) (agg.Campaign, error)
	FindAll(context.Context) ([]agg.Campaign, error)
	FindBy(context.Context, bson.M) ([]agg.Campaign, error)
	Update(context.Context, agg.Campaign) error
	Delete(context.Context, string) error
}
