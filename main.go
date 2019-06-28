package main

import (
	"github.com/gin-gonic/gin"

	"github.com/ikotan/go-todo-template/src/controller"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.Index)
	router.Run(":3001")
}
