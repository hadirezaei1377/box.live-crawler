package database

import (
	"box-crawler/entities"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database interface {
	InsertUpcomingFight(entities.Fight)
	GetUpcomingFights()
	GetUpcomingFight()
}

type mongodb struct {
	db *mongo.Client
}

func New(dsn string) (*mongodb, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dsn))
	return &mongodb{
		db: client,
	}, err
}

func (db mongodb) InsertUpcomingFight(fihgt entities.Fight) {
	db.db.Database("default_db").Collection("upcoming_fights").UpdateOne(context.TODO(), fihgt, options.Update().SetUpsert(true))
}
func (db mongodb) GetUpcomingFights() {}
func (db mongodb) GetUpcomingFight()  {}
