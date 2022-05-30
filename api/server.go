package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	db "github.com/pArtour/buss_api/db/sqlc"
	"net/http"
)

// Server serves all HTTP requests
type Server struct {
	store  *db.Store
	router *mux.Router
}

// NewServer creates a new server and setups routing
func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	r := mux.NewRouter()

	// adding routes
	r.HandleFunc("/places", server.getPlaces).Methods("GET")
	r.HandleFunc("/stops", server.getStops).Methods("GET").Queries("place", "{place}")
	r.HandleFunc("/route-info", server.getStopRouteInfo).Methods("GET").Queries("stop", "{stop}")
	server.router = r

	return server
}

// Start starts HTTP server on a specific address
func (s *Server) Start(addr string) error {
	return http.ListenAndServe(addr, s.router)
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	enableCors(w)
	w.Header().Set("Content-Type", "application/json")
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(code)
	w.Write(response)
}

// enableCors sets cors header
func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}
