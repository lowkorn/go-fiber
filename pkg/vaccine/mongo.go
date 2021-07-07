package vaccine

import (
	"context"
	"fmt"
	"time"

	"github.com/lowkorn/vaccine-reservation/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database   = "main"
	collection = "vaccines"
)

type MongoClient struct {
	DB     *mongo.Database
	Client *mongo.Client
}

func (mc MongoClient) Create(vc Vaccine) (Vaccine, error) {
	seed := time.Now().Unix()
	vc.ID = util.Hash(fmt.Sprint(seed))
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	_, err := mc.DB.Collection(collection).InsertOne(ctx, vc)
	if err != nil {
		return Vaccine{}, err
	}
	return vc, nil
}

func (mc MongoClient) GetAll() ([]Vaccine, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	cur, err := mc.DB.Collection(collection).Find(ctx, bson.M{})
	var vaccineReservations []Vaccine
	if err != nil {
		return vaccineReservations, err
	}
	cur.All(ctx, &vaccineReservations)
	return vaccineReservations, nil
}

func (mc MongoClient) GetByID(ID string) (Vaccine, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	filter := bson.M{"_id": ID}
	result := mc.DB.Collection(collection).FindOne(ctx, filter)
	var vaccine Vaccine
	result.Decode(&vaccine)
	return vaccine, nil
}

func (mc MongoClient) Edit(ID string, vc Vaccine) (Vaccine, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	filter := bson.M{"_id": ID}
	update := bson.M{
		"$set": vc,
	}
	_, err := mc.DB.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return Vaccine{}, err
	}
	return vc, nil
}

func (mc MongoClient) Delete(ID string) error {
	return nil
}

func NewMongoClient(conn *mongo.Client) MongoClient {
	return MongoClient{
		DB: conn.Database(database),
	}
}
