package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"net/http"
	"os"
)

type StringResponse struct {
	Message string `json:"Message"`
}

func main() {
	var portStr string
	err := godotenv.Load()
	if err != nil {
		println("Error loading .env file")
	}
	portStr = ":" + os.Getenv("PORT")

	mux_router := mux.NewRouter()
	mux_router.HandleFunc("/ping", ping)
	println("Listening on port " + portStr + "...")
	err = http.ListenAndServe(portStr, mux_router)
	if err != nil {
		panic(err)
	}
}

func ping(w http.ResponseWriter, _ *http.Request) {
	response := StringResponse{
		Message: "Hello, World!",
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	_, err = w.Write(jsonResponse)
	if err != nil {
		panic(err)
	}
}
