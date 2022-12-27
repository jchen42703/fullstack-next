package main

import (
	"fmt"
	"os"

	"github.com/jchen42703/echo-api/auth"
	"github.com/jchen42703/echo-api/controllers/baseRoutes"
	"github.com/jchen42703/echo-api/router"
)

func main() {
	baseKratosUrl := os.Getenv("ORY_KRATOS_BASE_URL")
	oryClient := auth.NewOryAPIClient(baseKratosUrl)
	r := router.New(oryClient)
	// connections, err := db.NewConnections()
	// defer connections.DB.Close()

	v1 := r.Group("/api")
	baseRoutes.RegisterRoutes(v1)

	serverUrl := os.Getenv("SERVER_URL")
	fmt.Println("server url: ", serverUrl)
	r.Logger.Fatal(r.Start(serverUrl))
}
