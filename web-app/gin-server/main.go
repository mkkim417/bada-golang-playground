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

	/*
			// Binding from JSON
		type Login struct {
			User     string `form:"user" json:"user" xml:"user"  binding:"required"`
			Password string `form:"password" json:"password" xml:"password" binding:"required"`
		}

		func main() {
			router := gin.Default()

			// Example for binding JSON ({"user": "manu", "password": "123"})
			router.POST("/loginJSON", func(c *gin.Context) {
				var json Login
				if err := c.ShouldBindJSON(&json); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if json.User != "manu" || json.Password != "123" {
					c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
					return
				}

				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			})

			// Example for binding XML (
			//	<?xml version="1.0" encoding="UTF-8"?>
			//	<root>
			//		<user>manu</user>
			//		<password>123</password>
			//	</root>)
			router.POST("/loginXML", func(c *gin.Context) {
				var xml Login
				if err := c.ShouldBindXML(&xml); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if xml.User != "manu" || xml.Password != "123" {
					c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
					return
				}

				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			})

			// Example for binding a HTML form (user=manu&password=123)
			router.POST("/loginForm", func(c *gin.Context) {
				var form Login
				// This will infer what binder to use depending on the content-type header.
				if err := c.ShouldBind(&form); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
					return
				}

				if form.User != "manu" || form.Password != "123" {
					c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
					return
				}

				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			})

			// Listen and serve on 0.0.0.0:8080
			router.Run(":8080")
		}
	*/
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
