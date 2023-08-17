package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type Data struct {
	Response string `json:"response"`
}

func main() {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	PORT := ":" + os.Getenv("PORT")
	fmt.Printf("%T %v", os.Getenv("PORT"), os.Getenv("PORT"))

	if os.Getenv("PORT") == "" {
		PORT = ":3334"
	}

	fmt.Println(PORT)

	// Routes
	http.HandleFunc("/ping", homePage)
	http.HandleFunc("/test", createNewPing)

	// Start the server
	http.ListenAndServe(PORT, nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	resp := Data{Response: "valoorrr"}

	jsonResp, err := json.Marshal(resp)

	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set the appropriate content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)

	fmt.Println("Endpoint Hit: [GET]'/ping'")
}

func createNewPing(w http.ResponseWriter, r *http.Request) {
	// Parse the incoming JSON request body
	var requestData Data
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// Process the requestData and perform necessary actions

	// Create a response object
	response := map[string]interface{}{
		"message": "Data received successfully",
	}

	// Convert the response to JSON format
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to marshal JSON", http.StatusInternalServerError)
		return
	}

	// Set the appropriate content type header
	w.Header().Set("Content-Type", "application/json")

	// Write the JSON response
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

	fmt.Println(requestData)

	fmt.Println("Endpoint Hit: [POST]'/test'")
}
