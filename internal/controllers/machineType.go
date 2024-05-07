package controllers

import "github.com/gin-gonic/gin"

func ViewMachineTypes(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "List all Machine Types",
	})
}

func ViewMachineType(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"message": "List Machine Type by id",
		"id":      id,
	})
}

func EditMachineType(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"message": "Edit Machine Type",
		"id":      id,
	})
}

func CreateMachineType(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Create Machine Type",
	})
}

func DeleteMachineType(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"message": "Delete Machine Type",
		"id":      id,
	})
}
