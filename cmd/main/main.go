package main

import (
	"fetch/management-console/internal/config"
	"fetch/management-console/internal/routes"
)

func main() {
	config.LoadEnv()
	routes.NewRoute(config.API_PORT)
}
