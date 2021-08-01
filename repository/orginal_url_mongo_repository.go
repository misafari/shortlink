package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"ir.safari.shortlink/model"
)

type OriginalUrlRepository interface {
	InsertOne(*model.Url) error
}

type originalUrlRepository struct {
	collection *mongo.Collection
	ctx context.Context
}

func (o *originalUrlRepository) InsertOne(url *model.Url) error {
	_, err := o.collection.InsertOne(o.ctx, url)
	return err
}

func NewOriginalUrlRepository(mongoClient *mongo.Client, ctx context.Context) OriginalUrlRepository {
	return &originalUrlRepository{
		collection: mongoClient.Database("short_link").Collection("origin_url"),
		ctx: ctx,
	}
}
