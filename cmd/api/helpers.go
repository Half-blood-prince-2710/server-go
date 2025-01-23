package main

import (
	"maps"
	"net/http"
)


func (app *application) writeJSON(w http.ResponseWriter, status int, headers http.Header) {
	maps.Insert(w.Header(),maps.All(headers))
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
}