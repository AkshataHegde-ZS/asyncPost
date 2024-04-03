package store

import (
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"

	"asyncPost/model"
)

type employee struct{}

func New() *employee {
	return &employee{}
}

func (s *employee) Create(ctx *gofr.Context, employee *model.Employee) error {
	_, err := ctx.DB().Exec(createQuery, employee.Name)
	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

const (
	createQuery = `INSERT INTO employee(name) VALUES(?)`
)
