package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

// HeaderResponse defines the structure of the json response
type HeaderResponse struct {
	IP        *string `json:"ip"`
	UserAgent *string `json:"user_agent"`
	Language  *string `json:"language"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	userAgent := r.UserAgent()
	acceptLang := r.Header.Get("Accept-Language")
	var language string

	if acceptLang != "" {
		language = strings.Split(acceptLang, ",")[0]
	} else {
		language = ""
	}

	headerResponse := HeaderResponse{&r.Host, &userAgent, &language}
	json.NewEncoder(w).Encode(headerResponse)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "localhost:8080"
	} else {
		port = ":" + port
	}
	http.HandleFunc("/", indexHandler)

	http.ListenAndServe(port, nil)
}
