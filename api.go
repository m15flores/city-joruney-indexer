package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	db *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{db: db}
}

func (s *Server) routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /cities", s.handleGetCities)
	mux.HandleFunc("GET /cities/{id}", s.handleGetCityByID)
	return mux
}

func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.routes())
}

func (s *Server) handleGetCities(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	if query.Has("at") && (query.Has("from") || query.Has("to")) {
		http.Error(w, "Use either at or from/to, not both", http.StatusBadRequest)
		return
	}
	if query.Has("at") {
		s.respondCityAt(w, query.Get("at"))
		return
	}
	if query.Has("from") || query.Has("to") {
		s.respondCitiesBetween(w, query.Get("from"), query.Get("to"))
		return
	}
	s.respondAllCities(w)
}

func (s *Server) respondAllCities(w http.ResponseWriter) {
	response, err := getAllCities(s.db)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) respondCityAt(w http.ResponseWriter, atStr string) {
	at, err := parseDate(atStr)
	if err != nil {
		http.Error(w, "at must be YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	response, err := getCityAt(s.db, at)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) respondCitiesBetween(w http.ResponseWriter, fromStr string, toStr string) {
	if fromStr == "" || toStr == "" {
		http.Error(w, "Query needs both: from and to", http.StatusBadRequest)
		return
	}

	from, err := parseDate(fromStr)
	if err != nil {
		http.Error(w, "from must be YYYY-MM-DD", http.StatusBadRequest)
		return
	}
	to, err := parseDate(toStr)
	if err != nil {
		http.Error(w, "to must be YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	if from >= to {
		http.Error(w, "from must be before to", http.StatusBadRequest)
		return
	}

	response, err := getCitiesBetween(s.db, from, to)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (s *Server) handleGetCityByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	response, err := getCityByID(s.db, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		http.Error(w, "Something wrong happened", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func parseDate(date string) (int64, error) {
	dateParsed, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, err
	}
	return dateParsed.Unix(), nil
}
