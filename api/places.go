package api

import (
	"net/http"
)

func (s *Server) getPlaces(w http.ResponseWriter, r *http.Request) {
	places, err := s.store.ListPlaces(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, places)
}
