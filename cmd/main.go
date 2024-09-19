package main

import (
	"log"

	"github.com/temur-shamshidinov/univer_grades/api"
	"github.com/temur-shamshidinov/univer_grades/config"
	db "github.com/temur-shamshidinov/univer_grades/pkg"
	"github.com/temur-shamshidinov/univer_grades/storage"
)

func main() {
	cfg := config.Load()

	db, err := db.ConnectToDb(cfg.PgConfig)
	if err != nil {
		log.Println("error on connect to ConToDb:", err)
		return
	}
	storage := storage.NewStorage(db)

	api.Api(storage)

	
}

