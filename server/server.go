package server

import (
	"rest-api/ethclient"

	"github.com/gorilla/mux"
)

type server struct {
	Router    *mux.Router
	ethClient *ethclient.CloudflareEthGateway
}

func New(client *ethclient.CloudflareEthGateway) *server {
	s := &server{
		ethClient: client,
		Router:    mux.NewRouter(),
	}
	s.initRoutes()
	return s
}

func (s *server) initRoutes() {
	// HEALTHZ endpoint
	//router.GET("/healthz", s.GetHealthz)
	s.Router.HandleFunc("/healthz", s.IndexHandler()).Methods("GET")

	// Get Block endpoint
	s.Router.HandleFunc("/block/{block_param}", s.GetBlockHandler()).Methods("GET")
	// router.GET("/block/:block_param", s.GetBlock)

	// // Get Transaction endpoint
	// router.GET("/block/:block_param/txs/:txs_param", s.GetTransaction)
}
