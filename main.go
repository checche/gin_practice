package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"gin_practice/models"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	dbInit()

	router.GET("/", func(ctx *gin.Context) {
		todos := dbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos})
	})
	router.Run()
}
