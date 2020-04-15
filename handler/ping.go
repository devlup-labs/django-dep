package handler

import (
	"net/http"
)

// Ping handler is used for health check purpose
var Ping = func(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("pong"))
}
