package tests

import (
	"github.com/goravel/framework/testing"

	"github.com/fahrigunadi/playground/bootstrap"
)

func init() {
	bootstrap.Boot()
}

type TestCase struct {
	testing.TestCase
}
