package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GreetingsHandler(c *gin.Context) {
	c.String(http.StatusOK, "hello gin")
}
