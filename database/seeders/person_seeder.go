package seeders

import (
	"github.com/fahrigunadi/playground/app/models"
	"github.com/goravel/framework/facades"
)

type PersonSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *PersonSeeder) Signature() string {
	return "PersonSeeder"
}

// Run executes the seeder logic.
func (s *PersonSeeder) Run() error {
	var persons []models.Person
	err := facades.Orm().Factory().Count(10).Create(&persons)

	return err
}
