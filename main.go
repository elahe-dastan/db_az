package main

import (
	"fmt"
	"log"

	"github.com/elahe-dastan/db_az/db"
	"github.com/elahe-dastan/db_az/handler"
)

func main() {
	database, err := db.New()
	if err != nil {
		log.Fatal("database initiation failed", err)
	}

	song := handler.Song{Store: database}
	song.Create() // Create table
	errMessage := song.AddRows()
	if errMessage != "" {
		log.Println("couldn't add rows", errMessage)
	}

	songs, err := song.ReadTable()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(songs)
}
