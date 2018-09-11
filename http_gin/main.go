package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", RetunJson)
	router.GET("/user/:name", GetName)
	router.GET("/someGet", RetunString)
	router.POST("/somePost", RetunJson)
	router.PUT("/somePut", RetunJson)
	router.DELETE("/someDelete", RetunJson)
	router.PATCH("/somePatch", RetunJson)
	router.HEAD("/someHead", RetunJson)
	router.POST("/post", PostName)
	//router.OPTIONS("/someOptions", options)
	router.Run(":80") // listen and serve on 0.0.0.0:8080
}

func GetName(c *gin.Context) {
	name := c.Param("name")
	c.String(200, "Hello %s", name)
}

func RetunJson(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello,world",
	})
}

func PostName(c *gin.Context) {
	id := c.Query("id")
	page := c.DefaultQuery("page", "0")
	name := c.PostForm("name")
	message := c.PostForm("message")

	fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
}

func RetunString(c *gin.Context) {
	c.String(200, "Hello,This is string")
}
