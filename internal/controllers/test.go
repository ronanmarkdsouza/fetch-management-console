package controllers

import "github.com/gin-gonic/gin"

func Test(ctx *gin.Context) {
	id := ctx.Query("id")
	ctx.JSON(200, gin.H{
		"message": id,
	})
}
