package repo

import (
	"context"
	"fmt"
	"time"

	"github.com/lowkorn/vaccine-reservation/pkg/entity"
	"github.com/lowkorn/vaccine-reservation/pkg/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	database   = "main"
	collection = "vaccines"
)

type VaccineMongo struct {
	DB     *mongo.Database
	Client *mongo.Client
}

func (r VaccineMongo) Create(vc entity.Vaccine) (entity.Vaccine, error) {
	seed := time.Now().Unix()
	vc.ID = util.Hash(fmt.Sprint(seed))
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	_, err := r.DB.Collection(collection).InsertOne(ctx, vc)
	if err != nil {
		return entity.Vaccine{}, err
	}
	return vc, nil
}

func (r VaccineMongo) GetAll() ([]entity.Vaccine, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	cur, err := r.DB.Collection(collection).Find(ctx, bson.M{})
	var vaccineReservations []entity.Vaccine
	if err != nil {
		return vaccineReservations, err
	}
	cur.All(ctx, &vaccineReservations)
	return vaccineReservations, nil
}

func (r VaccineMongo) GetByID(ID string) (entity.Vaccine, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	filter := bson.M{"_id": ID}
	result := r.DB.Collection(collection).FindOne(ctx, filter)
	var vaccine entity.Vaccine
	result.Decode(&vaccine)
	return vaccine, nil
}

func (r VaccineMongo) Edit(ID string, vc entity.Vaccine) (entity.Vaccine, error) {
	ctx, _ := context.WithTimeout(context.Background(), 60*time.Second)
	filter := bson.M{"_id": ID}
	update := bson.M{
		"$set": vc,
	}
	_, err := r.DB.Collection(collection).UpdateOne(ctx, filter, update)
	if err != nil {
		return entity.Vaccine{}, err
	}
	return vc, nil
}

func (r VaccineMongo) Delete(ID string) error {
	return nil
}

func NewVaccineMongo(conn *mongo.Client) VaccineMongo {
	return VaccineMongo{
		DB: conn.Database(database),
	}
}
