package main

import (
	"log"

	"github.com/elahe-dastan/db_az/db"
)

func main() {
	_, err := db.New()
	if err != nil {
		log.Fatal("database initiation failed", err)
	}
}
