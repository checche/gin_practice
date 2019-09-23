package main

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"gin_practice/models"
)

func main() {
	//ルーティング定義
	router := gin.Default()
	//テンプレートの保存先を指定
	router.LoadHTMLGlob("templates/*.html")

	// importしたやつはこんな感じで使う
	models.DbInit()
	router.GET("/", func(ctx *gin.Context) {
		todos := models.DbGetAll()
		ctx.HTML(200, "index.html", gin.H{
			"todos": todos,
		})
	})

	router.POST("/new", func(ctx *gin.Context) {
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		models.DbInsert(text, status)
		// 302として'/'にredirect
		ctx.Redirect(302, "/")
	})

	// :idについてる数字が"id"という文字列として利用できる.
	router.GET("/detail/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		// 文字列をintにしている.
		id, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		todo := models.DbGetOne(id)
		ctx.HTML(200, "detail.html", gin.H{"todo": todo})
	})

	router.POST("/update/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		text := ctx.PostForm("text")
		status := ctx.PostForm("status")
		models.DbUpdate(id, text, status)
		ctx.Redirect(302, "/")
	})

	router.GET("/delete_check/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		todo := models.DbGetOne(id)
		ctx.HTML(200, "delete.html", gin.H{"todo": todo})
	})

	router.POST("/delete/:id", func(ctx *gin.Context) {
		n := ctx.Param("id")
		id, err := strconv.Atoi(n)
		if err != nil {
			panic("ERROR")
		}
		models.DbDelete(id)
		ctx.Redirect(302, "/")
	})

	router.Run()
}
