package server

import (
	"rest-api/proxy"

	"github.com/gorilla/mux"
)

type server struct {
	Router *mux.Router
	Proxy  *proxy.ProxyImpl
}

func New(proxy *proxy.ProxyImpl) *server {
	s := &server{
		Proxy:  proxy,
		Router: mux.NewRouter(),
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

	s.Router.HandleFunc("/block/{block_param}/txs/{txs_param}", s.GetTransactionHandler()).Methods("GET")

	// s.Router.HandleFunc("/save-cache", s.SaveCache()).Methods("GET")
	// s.Router.HandleFunc("/get-cache", s.GetCache()).Methods("GET")

	// // Get Transaction endpoint
	// router.GET("/block/:block_param/txs/:txs_param", s.GetTransaction)
}
