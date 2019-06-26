package controller

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

// Index displays application index page
func Index(c *gin.Context) {
    c.String(http.StatusOK, "Hello, world!!!!")
}