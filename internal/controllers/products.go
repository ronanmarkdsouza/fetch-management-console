package controllers

import "github.com/gin-gonic/gin"

func ViewProducts(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "List all Products",
	})
}

func ViewProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"message": "List Product by id",
		"id":      id,
	})
}

func EditProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"message": "Edit Product",
		"id":      id,
	})
}

func CreateProduct(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Create Product",
	})
}

func DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"message": "Delete Product",
		"id":      id,
	})
}
