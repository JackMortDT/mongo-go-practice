package person

import (
	"my-rest-api/domain/person"
	"my-rest-api/utils/error_utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	PersonService personServiceInterface = &personService{}
)

type personService struct{}

type personServiceInterface interface {
	GetPerson(bson.M) ([]bson.M, error_utils.MessageErr)
	GetAllPeople() ([]bson.M, error_utils.MessageErr)
	CreatePerson(person *person.Person) (*mongo.InsertOneResult, error_utils.MessageErr)
	UpdatePerson(bson.M, primitive.M) (*mongo.UpdateResult, error_utils.MessageErr)
	DeletePerson(bson.M) (*mongo.DeleteResult, error_utils.MessageErr)
}

func (per *personService) GetPerson(personId bson.M) ([]bson.M, error_utils.MessageErr) {
	person, err := person.PersonRepo.Get(personId)
	if err != nil {
		return nil, err
	}
	return person, nil
}

func (per *personService) GetAllPeople() ([]bson.M, error_utils.MessageErr) {
	people, err := person.PersonRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return people, nil
}

func (per *personService) CreatePerson(personCreated *person.Person) (*mongo.InsertOneResult, error_utils.MessageErr) {
	pp, err := person.PersonRepo.Create(personCreated)
	if err != nil {
		return nil, err
	}
	return pp, nil
}

func (per *personService) UpdatePerson(pGet bson.M, pId primitive.M) (*mongo.UpdateResult, error_utils.MessageErr) {
	pp, err := person.PersonRepo.Update(pGet, pId)
	if err != nil {
		return nil, err
	}
	return pp, nil
}

func (per *personService) DeletePerson(personId bson.M) (*mongo.DeleteResult, error_utils.MessageErr) {
	pp, err := person.PersonRepo.Delete(personId)
	if err != nil {
		return nil, err
	}
	return pp, nil
}
