package main

import (
	_ "encoding/json"
	"example.com/user/hello/gin/milddle"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BookAddress struct {
	Books   []string       `form:"books" json:"books" binding:"required"`
	Address map[string]int `form:"address" json:"address" binding:"required"`
}

type BookAddressStr struct {
	Books   string `form:"books" json:"books" binding:"required"`
	Address string `form:"address" json:"address" binding:"required"`
}

func main() {
	r := gin.Default()

	r.Use(milddle.Time())

	r.GET("/ping", func(context *gin.Context) {
		name := context.Param("name")
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"name":    name,
		})
	})

	r.GET("/ping/:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")

		context.JSON(http.StatusOK, gin.H{
			"name":   name,
			"action": action,
		})
	})

	r.GET("/query", func(context *gin.Context) {
		firstName := context.Query("firstName")
		secondName, _ := context.GetQuery("secondName")
		context.JSON(http.StatusOK, gin.H{
			"firstName":  firstName,
			"secondName": secondName,
		})
	})

	// Content-Type:application/json
	/**
	{
		"books":["book1", "book2"],
		"address":{
			"No":1
		}
	}
	*/
	r.POST("/post", func(context *gin.Context) {
		var ba BookAddress
		if err := context.ShouldBindJSON(&ba); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		books := ba.Books
		address := ba.Address
		context.JSON(http.StatusOK, gin.H{
			"books":   books,
			"address": address,
		})
	})

	r.POST("/post1", func(context *gin.Context) {
		var ba BookAddressStr
		if err := context.ShouldBind(&ba); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		books := ba.Books
		address := ba.Address
		context.JSON(http.StatusOK, gin.H{
			"books":   books,
			"address": address,
		})
	})

	_ = r.Run(":8080")
}
