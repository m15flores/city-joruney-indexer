package main

import (
	"database/sql"

	_ "modernc.org/sqlite"
)

func openDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS cities (
			id INTEGER PRIMARY KEY,
			cityName TEXT, 
			latitude INTEGER,
			longitude INTEGER,
			fromDate INTEGER,
			toDate INTEGER
		)
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func insertCities(db *sql.DB, cities []City) error {

	for _, city := range cities {
		_, err := db.Exec(`
			INSERT OR REPLACE INTO cities (id, cityName, latitude, longitude, fromDate, toDate)
			VALUES(?, ?, ?, ?, ?, ?)
		`, city.TokenID, city.CityName, city.Latitude, city.Longitude, city.FromDate, city.ToDate)
		if err != nil {
			return err
		}
	}

	return nil
}
