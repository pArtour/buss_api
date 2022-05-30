package api

import (
	"net/http"
	"strconv"
)

type StopRouteInfoRequest struct {
	StopId int32 `json:"stopId"`
}

func (s *Server) getStopRouteInfo(w http.ResponseWriter, r *http.Request) {
	stopId := r.FormValue("stop")

	stopIdInt, err := strconv.ParseInt(stopId, 10, 32)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	info, err := s.store.GetStopRouteInfo(r.Context(), int32(stopIdInt))
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, info)
}
