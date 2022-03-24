package main

import (
	"gweb/gee"
	"log"
)

func main() {

	engine := gee.New()
	err := engine.Run(":8080")
	if err != nil {
		log.Fatalf("run has error : %v", err)
	}
}
