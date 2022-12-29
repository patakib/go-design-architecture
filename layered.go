`
In this example, the Entity, EntityRepository, and EntityService interfaces form the application's inner layer, while the MySQLRepository and PostgreSQLRepository structs are concrete implementations of the EntityRepository interface that belong to the infrastructure layer. The main function is the outer layer, and it depends on the inner layer (EntityService) to perform business logic.

The databaseType variable determines which database to connect to, and the appropriate repository is instantiated based on the value of databaseType. This allows the application to switch between MySQL and PostgreSQL easily, without affecting the inner layer.

This separation of concerns allows for more maintainable and testable code, as the inner layer can be tested in isolation from the infrastructure layer. It also makes it easier to swap out the infrastructure layer (e.g., switching to a different database) without affecting the inner layer.
`


package main

import (
	"fmt"
	"log"
)

// DatabaseType represents the type of database to connect to.
type DatabaseType int

const (
	// MySQL represents a MySQL database.
	MySQL DatabaseType = iota
	// PostgreSQL represents a PostgreSQL database.
	PostgreSQL
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

// MySQLRepository is a concrete implementation of EntityRepository that stores and retrieves entities from a MySQL database.
type MySQLRepository struct{}

// Find retrieves an entity from the MySQL database.
func (r *MySQLRepository) Find(id int) (*Entity, error) {
	// code to retrieve entity from MySQL database goes here
	return &Entity{ID: id, Name: "My Entity"}, nil
}

// Save stores an entity in the MySQL database.
func (r *MySQLRepository) Save(entity *Entity) error {
	// code to store entity in MySQL database goes here
	return nil
}

// PostgreSQLRepository is a concrete implementation of EntityRepository that stores and retrieves entities from a PostgreSQL database.
type PostgreSQLRepository struct{}

// Find retrieves an entity from the PostgreSQL database.
func (r *PostgreSQLRepository) Find(id int) (*Entity, error) {
	// code to retrieve entity from PostgreSQL database goes here
	return &Entity{ID: id, Name: "My Entity"}, nil
}

// Save stores an entity in the PostgreSQL database.
func (r *PostgreSQLRepository) Save(entity *Entity) error {
	// code to store entity in PostgreSQL database goes here
	return nil
}

// Service is an interface for performing business logic.
type Service interface {
	DoSomething(id int) (*Entity, error)
	DoSomethingElse(id int, name string) (*Entity, error)
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

// DoSomethingElse retrieves an entity, updates its name, and saves it.
func (s *EntityService) DoSomethingElse(id int, name string) (*Entity, error) {
	entity, err := s.repo.Find(id)
	if err != nil {
		return nil, err
	}

	entity.Name = name

	if err := s.repo.Save(entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func main() {
	var repo EntityRepository
	var service Service

	databaseType := MySQL
	switch databaseType {
	case MySQL:
		repo = &MySQLRepository{}
	case PostgreSQL:
		repo = &PostgreSQLRepository{}
	}

	service = &EntityService{repo: repo}

	entity, err := service.DoSomething(123)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(entity)

	updatedEntity, err := service.DoSomethingElse(123, "New Name")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(updatedEntity)
}

