package migrations

import (
	"gofr.dev/pkg/datastore"
	"gofr.dev/pkg/log"
)

type K20240402163759 struct {
}

const (
	createTableEmployee = `CREATE TABLE IF NOT EXISTS employee(
		id INT NOT NULL AUTO_INCREMENT,  
		name VARCHAR (20) NOT NULL,
		PRIMARY KEY (id)
	)`
)

func (k K20240402163759) Up(d *datastore.DataStore, logger log.Logger) error {
	_, err := d.DB().Exec(createTableEmployee)
	if err != nil {
		return err
	}

	return nil
}

func (k K20240402163759) Down(_ *datastore.DataStore, _ log.Logger) error {
	return nil
}
