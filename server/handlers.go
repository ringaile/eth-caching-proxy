package server

import (
	"fmt"
	"net/http"
	"rest-api/models"
	"strconv"
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
		data, err := s.Proxy.GetBlock(param)
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

		// get block
		data, err := s.Proxy.GetBlock(block_param)
		if err != nil {
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		// check if txr_param index
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

// func (s *server) SaveCache() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var emp proxy.Emp
// 		decoder := json.NewDecoder(r.Body)
// 		err := decoder.Decode(&emp)

// 		if err != nil {
// 			w.WriteHeader(http.StatusBadRequest)
// 			fmt.Fprintf(w, "Error saving cache")
// 			log.Print(err)
// 			return
// 		}
// 		proxy.SetCache("emp_data", emp)

// 		w.WriteHeader(http.StatusOK)
// 		fmt.Fprintf(w, "successfully! data has been saved in cache")
// 	}
// }

// func (s *server) GetCache() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		data, found := proxy.GetCache("emp_data")
// 		if found {
// 			fmt.Println(data)
// 			w.WriteHeader(http.StatusOK)
// 			fmt.Fprintf(w, "successfully! data has been saved in cache")
// 		} else {
// 			w.WriteHeader(http.StatusBadRequest)
// 			fmt.Fprintf(w, "Error! Not found key into cache")
// 			return
// 		}
// 	}
// }
