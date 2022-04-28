package main

import (
	"gweb/gee"
	"log"
	"net/http"
)

func main() {

	engine := gee.New()
	engine.GET("/test", func(context *gee.Context) {
		context.Json(http.StatusOK, "SUCCESS")
	})

	engine.GET("/test/:name/", func(context *gee.Context) {
		context.Json(http.StatusOK, "SUCCESS", context.Prams)
	})

	groupYYY := engine.Group("/yyy")
	{
		groupYYY.GET("/xxx", func(context *gee.Context) {
			context.Json(http.StatusNotFound, "not found")
		})
	}
	group := engine.Group("/test")
	{
		group.GET("/xxx/:name", func(context *gee.Context) {
			context.Json(http.StatusNotFound, "test xxx", context.Prams)
		})
	}

	err := engine.Run("127.0.0.1:8080")
	if err != nil {
		log.Fatalf("run has error : %v", err)
	}
}
