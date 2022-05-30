package api

import (
	"database/sql"
	"net/http"
)

func (s *Server) getStops(w http.ResponseWriter, r *http.Request) {
	place := r.FormValue("place")
	stops, err := s.store.ListStopsByPlace(r.Context(), place)

	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, stops)
}
