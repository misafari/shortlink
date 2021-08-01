package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Url struct {
	ID          primitive.ObjectID `bson:"_id"`
	OriginalUrl string             `bson:"original_url"`
	Key         string             `bson:"key"`
	CreatedAt   time.Time          `bson:"created_at"`
	UserId      int32              `bson:"user_id"`
}

type UrlBuilder struct {
	url *Url
}

func NewUrlBuilder() *UrlBuilder {
	return &UrlBuilder{url: new(Url)}
}

func (b *UrlBuilder) SetOriginalUrl(original string) *UrlBuilder {
	b.url.OriginalUrl = original
	return b
}

func (b *UrlBuilder) SetKey(key string) *UrlBuilder {
	b.url.Key = key
	return b
}

func (b *UrlBuilder) SetUserId(userId int32) *UrlBuilder {
	b.url.UserId = userId
	return b
}

func (b *UrlBuilder) Build() *Url {
	b.url.CreatedAt = time.Now()
	return b.url
}
