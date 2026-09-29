package api

import (
	"city-journey-indexer/internal/model"
	"city-journey-indexer/internal/store"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"slices"
	"testing"
)

func newTestDB(t *testing.T) *sql.DB {
	t.Helper()
	path := filepath.Join(t.TempDir(), "test.db")

	db, err := store.Open(path)
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

	if err := store.InsertCities(db, testCities); err != nil {
		t.Fatal(err)
	}
	return db
}

func TestHandleGetCities_All(t *testing.T) {
	testDB := newTestDB(t)

	req := httptest.NewRequest("GET", "/cities", nil)
	rec := httptest.NewRecorder()
	NewServer(testDB).routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", rec.Code, http.StatusOK)
	}

	var cities []model.City
	if err := json.NewDecoder(rec.Body).Decode(&cities); err != nil {
		t.Fatalf("decoding body: %v", err)
	}

	if len(cities) != 10 {
		t.Errorf("len(cities) = %d, want 10", len(cities))
	}
}

func TestHandleGetCityByID(t *testing.T) {
	testDB := newTestDB(t)

	tests := []struct {
		name     string
		path     string
		wantCode int
		wantCity string
	}{
		{name: "exists", path: "/cities/3", wantCode: http.StatusOK, wantCity: "Montreal"},
		{name: "not exists", path: "/cities/99", wantCode: http.StatusNotFound, wantCity: ""},
		{name: "non-numeric id", path: "/cities/abc", wantCode: http.StatusBadRequest, wantCity: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			rec := httptest.NewRecorder()
			NewServer(testDB).routes().ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantCode)
			}

			if rec.Code == http.StatusOK {
				var city model.City
				if err := json.NewDecoder(rec.Body).Decode(&city); err != nil {
					t.Fatalf("decoding body: %v", err)
				}

				if city.CityName != tt.wantCity {
					t.Fatalf("cityName = %q, want %q", city.CityName, tt.wantCity)
				}
			}
		})
	}
}

func TestHandleGetCityAt(t *testing.T) {
	testDB := newTestDB(t)

	tests := []struct {
		name     string
		path     string
		wantCode int
		wantCity string
	}{
		{name: "exists", path: "/cities?at=2023-07-15", wantCode: http.StatusOK, wantCity: "Montreal"},
		{name: "not exists", path: "/cities?at=1990-01-01", wantCode: http.StatusNotFound, wantCity: ""},
		{name: "invalid date", path: "/cities?at=hola", wantCode: http.StatusBadRequest, wantCity: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			rec := httptest.NewRecorder()
			NewServer(testDB).routes().ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantCode)
			}

			if rec.Code == http.StatusOK {
				var city model.City
				if err := json.NewDecoder(rec.Body).Decode(&city); err != nil {
					t.Fatalf("decoding body: %v", err)
				}

				if city.CityName != tt.wantCity {
					t.Fatalf("cityName = %q, want %q", city.CityName, tt.wantCity)
				}
			}

		})
	}
}

func TestHandleGetCitiesBetween(t *testing.T) {
	testDB := newTestDB(t)

	tests := []struct {
		name       string
		path       string
		wantCode   int
		wantCities []string
	}{
		{name: "all of 2023", path: "/cities?from=2023-01-01&to=2024-01-01", wantCode: http.StatusOK, wantCities: []string{"Mont-Tremblant", "Montreal", "Golden", "Squamish"}},
		{name: "border, only Quebec City", path: "/cities?from=2022-08-01&to=2022-09-01", wantCode: http.StatusOK, wantCities: []string{"Quebec City"}},
		{name: "no cities", path: "/cities?from=1990-01-01&to=1991-01-01", wantCode: http.StatusOK, wantCities: []string{}},
		{name: "missing to", path: "/cities?from=2023-01-01", wantCode: http.StatusBadRequest, wantCities: nil},
		{name: "invalid from", path: "/cities?from=hola&to=2024-01-01", wantCode: http.StatusBadRequest, wantCities: nil},
		{name: "from after to", path: "/cities?from=2024-01-01&to=2023-01-01", wantCode: http.StatusBadRequest, wantCities: nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.path, nil)
			rec := httptest.NewRecorder()
			NewServer(testDB).routes().ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Fatalf("status = %d, want %d", rec.Code, tt.wantCode)
			}

			if rec.Code == http.StatusOK {
				var cities []model.City
				if err := json.NewDecoder(rec.Body).Decode(&cities); err != nil {
					t.Fatalf("decoding body: %v", err)
				}

				citiesStr := []string{}
				for _, city := range cities {
					citiesStr = append(citiesStr, city.CityName)
				}

				if !slices.Equal(citiesStr, tt.wantCities) {
					t.Fatalf("cities = %v, want %v", citiesStr, tt.wantCities)
				}
			}
		})
	}
}

func TestHandleGetCities_AmbiguousQuery(t *testing.T) {
	testDB := newTestDB(t)

	req := httptest.NewRequest("GET", "/cities?at=2023-07-15&from=2023-01-01&to=2024-01-01", nil)
	rec := httptest.NewRecorder()
	NewServer(testDB).routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusBadRequest)
	}
}

func TestHandleGetCities_DBError(t *testing.T) {
	testDB := newTestDB(t)
	testDB.Close() // force any query to fail

	req := httptest.NewRequest("GET", "/cities", nil)
	rec := httptest.NewRecorder()
	NewServer(testDB).routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusInternalServerError)
	}
}

func TestHandleGetCityByID_DBError(t *testing.T) {
	testDB := newTestDB(t)
	testDB.Close()

	req := httptest.NewRequest("GET", "/cities/3", nil)
	rec := httptest.NewRecorder()
	NewServer(testDB).routes().ServeHTTP(rec, req)

	if rec.Code != http.StatusInternalServerError {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusInternalServerError)
	}
}
