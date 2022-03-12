package entity

type Campaign struct {
	ID         string `bson:"_id"`
	Name       string
	SpentItems []SpentItem
}
