package server

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

var validate = validator.New()

type Bar struct {
	Name  string `json:"name" yaml:"name" validate:"required"`
	Email string `json:"email" yaml:"email" validate:"omitempty,email"`
}

type Foo struct {
	Authors  []Bar    `validate:"required,gt=0,dive"`
	Keywords []string `validate:"omitempty,gt=0,dive,gt=0"`
}

func TestFoo(t *testing.T) {

	f := Foo{
		Authors: []Bar{
			{Name: "test"},
		},
		Keywords: []string{
			"f",
		},
	}

	err := validate.Struct(f)
	fmt.Println("struct ERR:", err)
	require.Error(t, err)
}
