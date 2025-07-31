package models

import (
	"github.com/fahrigunadi/playground/database/factories"
	"github.com/goravel/framework/contracts/database/factory"
	"github.com/goravel/framework/database/orm"
)

type Person struct {
	orm.Model
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func (u *Person) Factory() factory.Factory {
	return &factories.PersonFactory{}
}

func (r *Person) TableName() string {
	return "persons"
}
