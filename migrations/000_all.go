// This is auto-generated file using 'gofr migrate' tool. DO NOT EDIT.
package migrations

import (
	"gofr.dev/cmd/gofr/migration/dbMigration"
)

func All() map[string]dbmigration.Migrator{
	return map[string]dbmigration.Migrator{
	
		"20240402163759": K20240402163759{},
	}
}
