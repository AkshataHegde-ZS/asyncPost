package handler

import (
	"encoding/json"
	"os"
	"path/filepath"

	"gofr.dev/pkg/gofr"
	"gofr.dev/pkg/gofr/types"

	"asyncPost/model"
	"asyncPost/service"
)

type handler struct {
	svc service.Service
}

func New(svc service.Service) *handler {
	return &handler{svc: svc}
}

func (h *handler) Create(ctx *gofr.Context) (interface{}, error) {
	var emp model.Employee

	err := ctx.Bind(&emp)
	if err != nil {
		return nil, err
	}

	jsonData, err := json.Marshal(emp)
	if err != nil {
		return nil, err
	}

	dir := "tmp/"

	// Create a file with the Name field
	filename := filepath.Join(dir, emp.Name)
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	if _, err := file.Write(jsonData); err != nil {
		return nil, err
	}

	go h.svc.Create(ctx, filename)

	return types.Response{
		Data: "success",
	}, nil
}
