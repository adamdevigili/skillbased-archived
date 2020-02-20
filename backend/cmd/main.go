package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func main() {

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "the future home of balancer.team",
		})
	})

	r.Run()
}