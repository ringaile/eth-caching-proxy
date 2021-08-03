package server

import (
	"net/http"
)

func (s *server) GetBlockHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		param, err := handleBlock(w, r)
		if err != nil {
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		// get block
		data, err := s.Controller.GetBlock(param)
		if err != nil {
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		if data == nil {
			sendResponse(w, r, nil, http.StatusNotFound)
		}

		sendResponse(w, r, data, http.StatusOK)
	}
}

func (s *server) GetTransactionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		blockParam, trxParam, err := handleTransaction(w, r)
		if err != nil {
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		trx, err := s.Controller.GetTransaction(blockParam, trxParam)

		if err != nil {
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		if trx == nil {
			sendResponse(w, r, nil, http.StatusNotFound)
			return
		}

		sendResponse(w, r, trx, http.StatusOK)
	}
}
