package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type healthCheckResponse struct {
	Status      string `json:"status"`
	ProjectName string `json:"project_name"`
}

type versionCheckResponse struct {
	Version    string `json:"version"`
	ServerTime string `json:"server_time"`
}

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello, successfully connected")
}

func notFoundCheckHandler(w http.ResponseWriter, r *http.Request) {

	response := []byte("This page does not exists.")

	w.WriteHeader(http.StatusNotFound)

	w.Write(response)

}

func versionCheckHandler(w http.ResponseWriter, r *http.Request) {

	response := versionCheckResponse{
		Version:    "1.0.0",
		ServerTime: time.Now().Format(time.RFC3339),
	}

	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(JSONresponse)

}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	response := healthCheckResponse{
		Status:      "availible",
		ProjectName: "LibraryAPI",
	}

	JSONresponse, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(JSONresponse)

}

func Init() {
	http.HandleFunc("/healthcheck", healthCheckHandler)
	http.HandleFunc("/api/v1/info", versionCheckHandler)
	http.HandleFunc("/notfound", notFoundCheckHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
