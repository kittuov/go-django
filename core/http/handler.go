package http

import "net/http"

// GoDjHandle is created for URL configuration of Every App.
// This gets every request and runs through the middleware before going through mux.
type GoDjHandle map[string]interface{}

func (h *GoDjHandle) Handle(w http.ResponseWriter, r *http.Request) {
	//TODO: make middleware framework and add them here
}
