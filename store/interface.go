package store

import (
	"gofr.dev/pkg/gofr"

	"asyncPost/model"
)

type Store interface {
	Create(ctx *gofr.Context, emp *model.Employee) error
}
