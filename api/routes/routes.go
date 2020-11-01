package routes

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Path string
	Method string
	Handler http.HandlerFunc
}

func Install(router *mux.Router, categoryRoutes CategoryRoutes, productRoutes ProductRoutes) {

	allRoutes := categoryRoutes.Routes()
	allRoutes = append(allRoutes, productRoutes.Routes()...)

	for _, route := range allRoutes {
		router.HandleFunc(route.Path, route.Handler).Methods(route.Method)
	}
}