package httpapi

import (
	"chargeguard/internal/charging"
	"errors"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	station := r.URL.Query().Get("station")
	if station == "" {
		station = "deleted"
	}
	err := charging.ReportHazard(station)
	if err != nil {
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, charging.ErrStationNotFound):
			status = http.StatusNotFound
		case errors.Is(err, charging.ErrVersionConflict):
			status = http.StatusConflict
		}
		http.Error(w, err.Error(), status)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
