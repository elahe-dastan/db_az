package handler

import (
	"database/sql"
	"log"

	"github.com/elahe-dastan/db_az/model"
)

type Song struct {
	Store *sql.DB
}

func (s *Song) Create() {
	_, err := s.Store.Exec("CREATE TABLE IF NOT EXISTS song (ID int primary key, name varchar(255), views int);")
	if err != nil {
		log.Fatal("create table failed", err)
	}
}

func (s *Song) AddRows() string {
	errMessage := ""
	if _, err := s.Store.Exec("INSERT INTO song VALUES ('1', 'happier', '200000');"); err != nil {
		errMessage += err.Error()
	}

	if _, err := s.Store.Exec("INSERT INTO song VALUES ('2', 'frozen', '320000');"); err != nil {
		errMessage += err.Error()
	}

	if _, err := s.Store.Exec("INSERT INTO song VALUES ('3', 'money', '15000');"); err != nil {
		errMessage += err.Error()
	}

	return errMessage
}

func (s *Song) ReadTable() ([]model.Song, error) {
	var songs []model.Song
	rows, err := s.Store.Query("SELECT * FROM song")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var song model.Song
		if err = rows.Scan(&song.ID, &song.Name, &song.Views); err != nil {
			panic(err)
		}
		songs = append(songs, song)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return songs, nil
}
