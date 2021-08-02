package server

import (
	"rest-api/controller"

	"github.com/gorilla/mux"
)

type server struct {
	Router     *mux.Router
	Controller *controller.BlockController
}

func New(controller *controller.BlockController) *server {
	s := &server{
		Controller: controller,
		Router:     mux.NewRouter(),
	}
	s.initRoutes()
	return s
}

func (s *server) initRoutes() {

	// Get Block endpoint
	s.Router.HandleFunc("/block/{block_param}", s.GetBlockHandler()).Methods("GET")

	// Get Transaction endpoint
	s.Router.HandleFunc("/block/{block_param}/txs/{txs_param}", s.GetTransactionHandler()).Methods("GET")
}
