package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/suspect", func(c *gin.Context) {
		// get request body as string
		body, _ := c.GetRawData()

		// return the request body as string
		c.String(http.StatusOK, string(body))
	})

	r.Run()
}
