package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	ProductRoutes(r)
	ProfileRoutes(r)
	UserRoutes(r)
	Transaction(r)
	AuthRoutes(r)
	CartRoutes(r)
}
