package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	context context.Context
	mclient *mongo.Client
}

// New creates new mongodb client.
func New(ctx context.Context, uri string) (MongoDB, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return MongoDB{}, err
	}
	return MongoDB{context: ctx, mclient: client}, nil
}

// Close closes sockets to the topology referenced
// by this client.
func (db MongoDB) Close() {
	_ = db.mclient.Disconnect(db.context)
}

// Ping sends a ping command to verify that
// the client can connect to the deployment.
func (db MongoDB) Ping() error {
	return db.mclient.Ping(db.context, readpref.Primary())
}

// Database returns.
func (db MongoDB) Database(name string) *mongo.Database {
	return db.mclient.Database(name)
}
