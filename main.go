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
	err := engine.Run(":8080")
	if err != nil {
		log.Fatalf("run has error : %v", err)
	}
}
