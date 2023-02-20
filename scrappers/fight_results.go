package scrappers

import (
	"box-crawler/database"
	"box-crawler/entities"
)

func ScrapeFightResults(db database.Database) {
	// Get data
	// Extract correct data

	fight := entities.Fight{}
	db.InsertUpcomingFight(fight)
}
