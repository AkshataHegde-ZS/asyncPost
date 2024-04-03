package service

import (
	"gofr.dev/pkg/gofr"
)

type Service interface {
	Create(ctx *gofr.Context, fileName string)
}
