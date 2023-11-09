package main

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type StoredShow struct {
	gorm.Model
	MazeId    uint
	Name      string
	MySeason  uint
	MyEpisode uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&StoredShow{})

	// load the user's shows
	// 	open db
	// 	if they don't exist, create a new db
	// use gorm to pull the shows
	// fetch the data needed from tvmaze api

}
