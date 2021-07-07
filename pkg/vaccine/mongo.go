package vaccine

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoClient struct {
	*mongo.Client
}

func (MI MongoClient) Create() (Vaccine, error) {
	return Vaccine{}, nil
}

func NewMongoConnection() MongoClient {
	return MongoClient{}
}
