package routes

import (
	"waysbean/handlers"
	"waysbean/pkg/middleware"
	"waysbean/pkg/mysql"
	"waysbean/repositories"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	authReposutory := repositories.RepositoryAuth(mysql.DB)
	h := handlers.HandlerAuth(authReposutory)

	r.HandleFunc("/register", h.Register).Methods("POST")
	r.HandleFunc("/login", h.Login).Methods("POST")
	r.HandleFunc("/log-check", middleware.Auth(h.CheckAuth)).Methods("GET")
}
