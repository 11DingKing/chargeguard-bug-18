package httpapi

import (
	"chargeguard/internal/charging"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	err := charging.ReportHazard("deleted")
	if err != nil {
		http.Error(w, "state conflict", http.StatusConflict)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
