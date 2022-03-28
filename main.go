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

	group := engine.Group("/yyy")
	{
		group.GET("/xxx", func(context *gee.Context) {
			context.Json(http.StatusNotFound, "not found")
		})
	}
	err := engine.Run(":8080")
	if err != nil {
		log.Fatalf("run has error : %v", err)
	}
}
