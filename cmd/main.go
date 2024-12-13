package main

import (
	_ "velocity-technical-test/docs"
	"velocity-technical-test/internal/infraestructure/primary/server"
)

// @title Velocity Technical Test API
// @version 1.0
// @description This is a sample server for a technical test.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:60000
// @BasePath /api

func main() {
	server.RunServer()
}
