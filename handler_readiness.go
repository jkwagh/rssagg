package main

import (
	"net/http"
)

// should respond if everything is working properly
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, 200, struct{}{}) //calling on respondWithJSON function from json.go
}
