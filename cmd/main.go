package main

import (
	"log"

	"github.com/temur-shamshidinov/univer_grades/config"
	db "github.com/temur-shamshidinov/univer_grades/pkg/database"
)

func main() {

	_, err := db.ConnectToDb(config.Load())
	if err != nil {
		log.Println("error on connect to ConToDb:", err)
		return
	}
	// storage := storage.NewStorage(db)

	// api.Api(storage)

}
