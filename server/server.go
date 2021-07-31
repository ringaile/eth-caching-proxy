package server

import (
	"log"
	"net/http"

	"rest-api/config"
	"rest-api/server/handler"

	"github.com/gorilla/mux"
)

// Server has router
type Server struct {
	Router *mux.Router
}

// Initialize initializes the app with predefined configuration
func (a *Server) Initialize(config *config.Config) {
	a.Router = mux.NewRouter()
	a.setRouters()
}

// setRouters sets the all required routers
func (a *Server) setRouters() {
	// Routing for handling the projects
	a.Get("/health", a.handleRequest(handler.GetHealthz))
}

// Get wraps the router for GET method
func (a *Server) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// // Post wraps the router for POST method
// func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Methods("POST")
// }

// // Put wraps the router for PUT method
// func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Methods("PUT")
// }

// // Delete wraps the router for DELETE method
// func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) {
// 	a.Router.HandleFunc(path, f).Methods("DELETE")
// }

// Run the app on it's router
func (a *Server) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}

type RequestHandlerFunction func(w http.ResponseWriter, r *http.Request)

func (a *Server) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	}
}
