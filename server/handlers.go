package server

import (
	"net/http"
	"rest-api/models"
	"strconv"
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

		sendResponse(w, r, data, http.StatusOK)
	}
}

func (s *server) GetTransactionHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		block_param, trx_param, err := handleTransaction(w, r)
		if err != nil {
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		data, err := s.Controller.GetBlock(block_param)
		if err != nil {
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		// check txr_param index
		var result models.Transaction
		val, err := strconv.Atoi(trx_param)
		if err != nil {
			for _, trx := range data.Transactions {
				if trx.Hash == trx_param {
					result = trx
				}
			}
		} else {
			if len(data.Transactions) > val {
				result = data.Transactions[val]
			} else {
				sendResponse(w, r, nil, http.StatusInternalServerError)
				return
			}
		}

		sendResponse(w, r, result, http.StatusOK)
	}
}
