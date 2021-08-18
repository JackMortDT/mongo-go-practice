package app

import (
	"my-rest-api/domain"
	"my-rest-api/domain/person"

	"go.mongodb.org/mongo-driver/mongo"
)

func startDatabase(mongoUrl string) {
	client := domain.Initialize(mongoUrl)
	repositories(client)
}

func repositories(client *mongo.Client) {
	person.PersonRepo.Initialize(client)
}
