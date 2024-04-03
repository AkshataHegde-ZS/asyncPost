package service

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"gofr.dev/pkg/gofr"

	"asyncPost/model"
	"asyncPost/store"
)

type service struct {
	st store.Store
}

// Create implements service.Service.
func (s *service) Create(ctx *gofr.Context, fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		ctx.Logger.Errorf("error while opening a file,err: %v", err)
	}

	defer file.Close()

	// Read data from the temporary file
	data, err := io.ReadAll(file)
	if err != nil {
		ctx.Logger.Errorf("error while reading from the file,err: %v", err)
	}

	fmt.Println("Processing data:", string(data))

	var emp model.Employee

	err = json.Unmarshal(data, &emp)
	if err != nil {
		ctx.Logger.Errorf("error while unmarshalling,err: %v", err)
	}

	err = s.st.Create(ctx, &emp)
	if err != nil {
		ctx.Logger.Errorf("error from store,err: %v", err)
	}
}

func New(st store.Store) *service {
	return &service{st: st}
}
