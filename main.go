package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "hello world",
			"data": "",
		})
	})
	r.GET("/foo", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "bar",
			"data": "",
		})
	})
	r.Run()
}
