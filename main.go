package main

import (
	"net/http"

	user_routes "api1-multi.com/a/src/Users/infrastructure/routes"
	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	r := gin.Default()
	c := cors.New(cors.Options{
        AllowedOrigins:   []string{"http://localhost:4200"},
        AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
        AllowCredentials: true,
        Debug:            true,
    })

	user_routes.UserRoutes(r)
	handler := c.Handler(r)

	http.ListenAndServe(":8080", handler)
}