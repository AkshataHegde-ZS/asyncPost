package main

import (
	"log"
	"os"

	"gofr.dev/cmd/gofr/migration"
	dbmigration "gofr.dev/cmd/gofr/migration/dbMigration"
	"gofr.dev/pkg/gofr"

	"asyncPost/handler"
	"asyncPost/migrations"
	service "asyncPost/service/employee"
	store "asyncPost/store/employee"
)

func main() {
	app := gofr.New()

	//run migrations to create database
	runMigrations(app)

	//create a dir to store files
	if err := os.MkdirAll("tmp/test", 0755); err != nil {
		log.Fatal("unable to create dir")
	}

	st := store.New()
	ser := service.New(st)
	http := handler.New(ser)

	app.POST("/post", http.Create)

	app.Start()
}

func runMigrations(app *gofr.Gofr) {
	err := migration.Migrate(app.Config.Get("APP_NAME"), dbmigration.NewGorm(app.GORM()), migrations.All(), dbmigration.UP, app.Logger)
	if err != nil {
		app.Logger.Fatal(err)
	}
}
