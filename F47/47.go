// gin program
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/progress", func(c *gin.Context) {
		c.String(http.StatusOK, "Here Aishwarya")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8088")
}

//  http://0.0.0.0:8088/progress
