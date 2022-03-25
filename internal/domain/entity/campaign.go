package entity

type Campaign struct {
	ID         string      `bson:"_id"`
	Name       string      `bson:"name"`
	SpentItems []SpentItem `bson:"spentItems"`
}
