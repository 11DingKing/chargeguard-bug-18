package httpapi

import (
	"chargeguard/internal/charging"
	"errors"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	err := charging.ReportHazard("deleted")
	code := charging.ErrorCode(err)
	w.Header().Set("X-Error-Code", code)
	switch {
	case errors.Is(err, charging.ErrStationNotFound):
		http.Error(w, "station not found", http.StatusNotFound)
	case errors.Is(err, charging.ErrVersionConflict):
		http.Error(w, "version conflict", http.StatusConflict)
	case err != nil:
		http.Error(w, "storage error", http.StatusInternalServerError)
	default:
		w.WriteHeader(http.StatusCreated)
	}
}
