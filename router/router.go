package router

import (
	"net/http"
)

// Router will return a new servemux instance
func Router() *http.ServeMux {
	router := http.ServeMux{}

	router.HandleFunc("/", middleware.CreateUser)

	return router
}
