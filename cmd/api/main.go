package main

import (
	_ "crudapp/docs"
	"crudapp/internal/routes"
	"log"

	_ "github.com/swaggo/files"
)

// @title           CrudApp API
// @version         1.0
// @description     This is a sample server celler server.

// @host      localhost:8000
// @BasePath  /

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	router := routes.NewRouter()
	log.Fatal(router.Listen(":8000"))
}
