package server

import (
	"encoding/json"
	"log"
	"net/http"
	"rest-api/ethclient"
	"rest-api/models"

	"github.com/gorilla/mux"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func handleBlock(w http.ResponseWriter, r *http.Request) (string, error) {
	vars := mux.Vars(r)
	param := vars["block_param"]
	return param, nil
}

func sendResponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format json")
	}
}

func mapBlockToJson(b ethclient.Block) models.BlockResponse {
	return models.BlockResponse{
		Timestamp:  b.Timestamp,
		Number:     b.Number,
		Difficulty: b.Difficulty,
	}
}
