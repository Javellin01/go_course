package entity

type Advertiser struct {
	ID      string  `bson:"_id"`
	Name    string  `bson:"name"`
	Balance float32 `bson:"balance"`
}
