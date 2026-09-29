package store

import (
	"city-journey-indexer/internal/model"
	"database/sql"
	"errors"
	"path/filepath"
	"slices"
	"testing"
	"time"
)

func unix(y int, m time.Month, d int) int64 {
	return time.Date(y, m, d, 0, 0, 0, 0, time.UTC).Unix()
}

func newTestDB(t *testing.T) *sql.DB {
	t.Helper()
	path := filepath.Join(t.TempDir(), "test.db")

	db, err := Open(path)
	if err != nil {
		t.Fatalf("opening db: %v", err)
	}

	// Real data minted on Arbitrum One, copied from cities.db.
	var testCities = []model.City{
		{TokenID: 0, CityName: "Madrid", Latitude: 40416900, Longitude: -3703300, FromDate: 873072000, ToDate: 1659312000},
		{TokenID: 1, CityName: "Quebec City", Latitude: 46816100, Longitude: -71224200, FromDate: 1659312000, ToDate: 1669852800},
		{TokenID: 2, CityName: "Mont-Tremblant", Latitude: 46118000, Longitude: -74600000, FromDate: 1669852800, ToDate: 1682899200},
		{TokenID: 3, CityName: "Montreal", Latitude: 45508800, Longitude: -73561600, FromDate: 1682899200, ToDate: 1693526400},
		{TokenID: 4, CityName: "Golden", Latitude: 51301900, Longitude: -116966700, FromDate: 1693526400, ToDate: 1698796800},
		{TokenID: 5, CityName: "Squamish", Latitude: 49701700, Longitude: -123158800, FromDate: 1698796800, ToDate: 1714521600},
		{TokenID: 6, CityName: "Tofino", Latitude: 49152700, Longitude: -125904400, FromDate: 1717200000, ToDate: 1725148800},
		{TokenID: 7, CityName: "Madrid", Latitude: 40416900, Longitude: -3703300, FromDate: 1725148800, ToDate: 1751328000},
		{TokenID: 8, CityName: "Lyon", Latitude: 45759700, Longitude: 4842200, FromDate: 1751328000, ToDate: 1788220800},
		{TokenID: 9, CityName: "Aix-en-Provence", Latitude: 43527800, Longitude: 5445600, FromDate: 1788220800, ToDate: 0},
	}

	if err := InsertCities(db, testCities); err != nil {
		t.Fatal(err)
	}
	return db
}

func TestGetCityAt(t *testing.T) {
	testDB := newTestDB(t)

	tests := []struct {
		name string
		at   int64
		want string
	}{
		{"middle of Montreal", unix(2023, time.July, 15), "Montreal"},
		{"border between Madrid and Quebec City", unix(2022, time.August, 1), "Quebec City"},
		{"Aix with toDate 0", unix(2026, time.September, 28), "Aix-en-Provence"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			city, err := GetCityAt(testDB, tt.at)
			if err != nil {
				t.Fatalf("GetCityAt(%d) returned error: %v", tt.at, err)
			}
			if city.CityName != tt.want {
				t.Errorf("GetCityAt(%d) = %q, want %q", tt.at, city.CityName, tt.want)
			}
		})
	}
}

func TestGetCityAt_NotFound(t *testing.T) {
	testDB := newTestDB(t)

	_, err := GetCityAt(testDB, unix(1990, time.January, 1))
	if !errors.Is(err, sql.ErrNoRows) {
		t.Errorf("GetCityAt(%d) error = %v, want sql.ErrNoRows", unix(1990, time.January, 1), err)
	}
}

func TestGetCitiesBetween(t *testing.T) {
	testDB := newTestDB(t)

	tests := []struct {
		name string
		from int64
		to   int64
		want []string
	}{
		{"all of 2023", unix(2023, time.January, 1), unix(2024, time.January, 1), []string{"Mont-Tremblant", "Montreal", "Golden", "Squamish"}},
		{"border between Madrid and Quebec City", unix(2022, time.August, 1), unix(2022, time.September, 1), []string{"Quebec City"}},
		{"no cities", unix(1990, time.January, 1), unix(1991, time.January, 1), []string{}},
		{"until Aix", unix(2023, time.June, 1), unix(2026, time.October, 1), []string{"Montreal", "Golden", "Squamish", "Tofino", "Madrid", "Lyon", "Aix-en-Provence"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cities, err := GetCitiesBetween(testDB, tt.from, tt.to)
			if err != nil {
				t.Fatalf("GetCitiesBetween(%d, %d) returned error: %v", tt.from, tt.to, err)
			}
			citiesStr := []string{}
			for _, city := range cities {
				citiesStr = append(citiesStr, city.CityName)
			}
			if !slices.Equal(citiesStr, tt.want) {
				t.Errorf("GetCitiesBetween(%d, %d) = %v, want %v", tt.from, tt.to, citiesStr, tt.want)
			}
		})
	}
}

func TestGetCityByID(t *testing.T) {
	testDB := newTestDB(t)

	tests := []struct {
		name    string
		id      int64
		want    string
		wantErr bool
	}{
		{"existing city", 3, "Montreal", false},
		{"nonexistent city", 99, "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			city, err := GetCityByID(testDB, tt.id)

			if tt.wantErr {
				if !errors.Is(err, sql.ErrNoRows) {
					t.Errorf("GetCityByID(%d) error = %v, want sql.ErrNoRows", tt.id, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("GetCityByID(%d) returned error: %v", tt.id, err)
			}
			if city.CityName != tt.want {
				t.Errorf("GetCityByID(%d) = %q, want %q", tt.id, city.CityName, tt.want)
			}
		})
	}
}

func TestGetAllCities(t *testing.T) {
	testDB := newTestDB(t)

	cities, err := GetAllCities(testDB)
	if err != nil {
		t.Fatalf("GetAllCities() returned error: %v", err)
	}

	if len(cities) != 10 {
		t.Errorf("len(GetAllCities()) = %d, want 10", len(cities))
	}

	if cities[0].CityName != "Madrid" {
		t.Errorf("cities[0].CityName = %q, want %q", cities[0].CityName, "Madrid")
	}
	if cities[9].CityName != "Aix-en-Provence" {
		t.Errorf("cities[9].CityName = %q, want %q", cities[9].CityName, "Aix-en-Provence")
	}
}
