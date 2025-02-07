package main

import (
	"encoding/json"

	"net/http"
)
type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope, headers http.Header) (error) {
	js,err:=json.MarshalIndent(data,"","\t")
	if err!=nil {
		return err
	}
	
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(status)
	w.Write(js)
	return  nil
}