package main

import "net/http"

func Handler_Err(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
}
