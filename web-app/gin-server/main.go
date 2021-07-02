package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	engine.GET("/query", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		c.JSON(200, gin.H{
			"id":   id,
			"page": page,
		})
	})

	engine.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	//http://localhost:8081/map?ids[a]=1234&ids[b]=hello
	engine.GET("/map", func(c *gin.Context) {

		ids := c.QueryMap("ids")
		names := c.PostFormMap("names")

		fmt.Printf("ids: %v; names: %v", ids, names)
	})

	engine.GET("/to-google", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://www.google.co.kr/")
	})

	/* router grouping */
	v1 := engine.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	v2 := engine.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)
		v2.POST("/submit", submitEndpoint)
		v2.POST("/read", readEndpoint)
	}

	engine.Run(":8081") // default: 8080
}

func loginEndpoint(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.google.co.kr/")
}
func submitEndpoint(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.google.co.kr/")
}
func readEndpoint(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "https://www.google.co.kr/")
}
