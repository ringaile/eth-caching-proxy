package server

import (
	"fmt"
	"log"
	"net/http"
)

func (s *server) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "welcome!")
	}
}

func (s *server) GetBlockHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param, err := handleBlock(w, r)
		if err != nil {
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		// get latest block
		data, err := s.ethClient.GetBlock(param)
		if err != nil {
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapBlockToJson(*data)
		log.Print(resp.Difficulty)
		sendResponse(w, r, resp, http.StatusOK)
	}
}
