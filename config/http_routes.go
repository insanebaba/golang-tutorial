package config

import "github.com/julienschmidt/httprouter"

// ConfigureRoutes common routes
func ConfigureRoutes(router *httprouter.Router) {
	ConfigureBookRoutes(router)
}
