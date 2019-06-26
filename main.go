package main

import (
    "github.com/ikotan/go-todo-template/src/controller"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/", controller.Index)
    router.Run(":3001")
}