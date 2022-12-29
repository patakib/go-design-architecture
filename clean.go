`
In this example, the Entity, EntityRepository, and EntityService interfaces form the application's inner layer, while the MySQLEntityRepository struct is a concrete implementation of the EntityRepository interface that belongs to the infrastructure layer. The main function is the outer layer, and it depends on the inner layer (EntityService) to perform business logic.

This separation of concerns allows for more maintainable and testable code, as the inner layer can be tested in isolation from the infrastructure layer. It also makes it easier to swap out the infrastructure layer (e.g., switching to a different database) without affecting the inner layer.
`

package main

import (
	"fmt"
	"log"
)

// Entity represents a business entity.
type Entity struct {
	ID   int
	Name string
}

// EntityRepository is an interface for storing and retrieving entities.
type EntityRepository interface {
	Find(id int) (*Entity, error)
	Save(entity *Entity) error
}

// MySQLEntityRepository is a concrete implementation of EntityRepository that stores and retrieves entities from a MySQL database.
type MySQLEntityRepository struct{}

// Find retrieves an entity from the database.
func (r *MySQLEntityRepository) Find(id int) (*Entity, error) {
	// code to retrieve entity from database goes here
	return &Entity{ID: id, Name: "My Entity"}, nil
}

// Save stores an entity in the database.
func (r *MySQLEntityRepository) Save(entity *Entity) error {
	// code to store entity in database goes here
	return nil
}

// Service is an interface for performing business logic.
type Service interface {
	DoSomething(id int) (*Entity, error)
}

// EntityService is a concrete implementation of Service that uses EntityRepository to perform business logic.
type EntityService struct {
	repo EntityRepository
}

// DoSomething retrieves an entity and does some business logic with it.
func (s *EntityService) DoSomething(id int) (*Entity, error) {
	entity, err := s.repo.Find(id)
	if err != nil {
		return nil, err
	}

	// business logic goes here

	return entity, nil
}

func main() {
	repo := &MySQLEntityRepository{}
	service := &EntityService{repo: repo}

	entity, err := service.DoSomething(123)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(entity)
}
