package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"ir.safari.shortlink/model"
)

type OriginalUrlRepository interface {
	InsertOne(*model.Url, int64) error
	FetchUrl(string) (*model.Url, error)
}

type originalUrlRepository struct {
	collection *mongo.Collection
	ctx        context.Context
}

func (o *originalUrlRepository) FetchUrl(s string) (result *model.Url, _ error) {
	filter := bson.D{{"code", s}}
	findErr := o.collection.FindOne(o.ctx, filter).Decode(result)
	return result, findErr
}

func (o *originalUrlRepository) InsertOne(url *model.Url, expiredTime int64) error {
	//todo set expired time
	_, err := o.collection.InsertOne(o.ctx, url)
	return err
}

func NewOriginalUrlRepository(mongoClient *mongo.Client, ctx context.Context) OriginalUrlRepository {
	return &originalUrlRepository{
		collection: mongoClient.Database("short_link").Collection("origin_url"),
		ctx:        ctx,
	}
}
