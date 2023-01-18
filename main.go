package main

import (
	sw "_case_api/go"
	"github.com/gin-gonic/gin"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"net/http"
	"os"
	"strconv"
)

// dev - host the swagger UI (dev) page if true, or the redoc page if false.
var dev = true

func main() {
	log.Printf("Server started")

	// check for dev variable, else use value above
	getDev := os.Getenv("SWAGGER_DEV")
	if getDev != "" {
		dev, _ = strconv.ParseBool(getDev)
	}

	// create a router, with all CASE REST APIs registered
	router := sw.NewRouter()

	// host and serve the openapi yaml
	router.StaticFile("/swagger", "./api/openapi.yaml")

	if dev {
		opts1 := middleware.SwaggerUIOpts{SpecURL: "/swagger", Path: "/docs"}
		sh1 := middleware.SwaggerUI(opts1, nil)
		router.GET("/docs", func(ctx *gin.Context) {
			sh1.ServeHTTP(ctx.Writer, ctx.Request)
		})
	} else {
		opts1 := middleware.RedocOpts{SpecURL: "/swagger", Path: "/docs"}
		sh1 := middleware.Redoc(opts1, nil)
		router.GET("/docs", func(ctx *gin.Context) {
			sh1.ServeHTTP(ctx.Writer, ctx.Request)
		})
	}

	router.GET("/", Index)

	log.Fatal(router.Run(":8080"))
}

func Index(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/docs")
}
