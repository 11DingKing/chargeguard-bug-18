package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTaskBehavior(t *testing.T) {
	rr := httptest.NewRecorder()
	TaskHTTPHandler(rr, httptest.NewRequest("POST", "/task", nil))
	if rr.Code != http.StatusNotFound {
		t.Fatalf("status=%d body=%s", rr.Code, rr.Body.String())
	}
}
