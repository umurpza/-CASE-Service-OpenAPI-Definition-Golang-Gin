package main

import (
	sw "_case_api/go"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"log"
)

func main() {
	log.Printf("Server started")

	// create a router, with all CASE REST APIs registered
	router := sw.NewRouter()

	// host and serve the openapi yaml
	router.StaticFile("/swagger", "./api/openapi.yaml")
	opts1 := middleware.RedocOpts{SpecURL: "/swagger", Path: "/redoc"}
	sh1 := middleware.Redoc(opts1, nil)
	router.GET("/redoc", func(ctx *gin.Context) {
		sh1.ServeHTTP(ctx.Writer, ctx.Request)
	})

	log.Fatal(router.Run(":8080"))
}
