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

func getAllCities(db *sql.DB) ([]City, error) {
	rows, err := db.Query("SELECT id, cityName, latitude, longitude, fromDate, toDate FROM cities ORDER BY Id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cities := []City{}

	for rows.Next() {
		var city City
		if err := rows.Scan(&city.TokenID, &city.CityName, &city.Latitude,
			&city.Longitude, &city.FromDate, &city.ToDate); err != nil {

			return nil, err
		}
		cities = append(cities, city)
	}
	if err = rows.Err(); err != nil {
		return cities, err
	}
	return cities, nil
}

func getCityByID(db *sql.DB, id int64) (City, error) {
	row := db.QueryRow("SELECT id, cityName, latitude, longitude, fromDate, toDate FROM cities WHERE id = ?", id)
	var city City

	if err := row.Scan(&city.TokenID, &city.CityName, &city.Latitude,
		&city.Longitude, &city.FromDate, &city.ToDate); err != nil {

		return city, err
	}

	return city, nil
}

func getCityAt(db *sql.DB, at int64) (City, error) {
	row := db.QueryRow("SELECT id, cityName, latitude, longitude, fromDate, toDate FROM cities WHERE fromDate <= ? AND (toDate > ? OR toDate = 0)", at, at)
	var city City

	if err := row.Scan(&city.TokenID, &city.CityName, &city.Latitude,
		&city.Longitude, &city.FromDate, &city.ToDate); err != nil {

		return city, err
	}

	return city, nil
}

func getCitiesBetween(db *sql.DB, from int64, to int64) ([]City, error) {
	rows, err := db.Query("SELECT id, cityName, latitude, longitude, fromDate, toDate FROM cities WHERE fromDate < ? AND (toDate > ? OR toDate = 0) ORDER BY Id", to, from)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cities := []City{}

	for rows.Next() {
		var city City
		if err := rows.Scan(&city.TokenID, &city.CityName, &city.Latitude,
			&city.Longitude, &city.FromDate, &city.ToDate); err != nil {

			return nil, err
		}
		cities = append(cities, city)
	}
	if err = rows.Err(); err != nil {
		return cities, err
	}
	return cities, nil

}
