package controllers

import "github.com/gin-gonic/gin"

func ViewVendingMachines(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "List all Vending Machines",
	})
}

func ViewVendingMachine(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"message": "List Vending Machine by id",
		"id":      id,
	})
}

func EditVendingMachine(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"message": "Edit Vending Machine",
		"id":      id,
	})
}

func CreateVendingMachine(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Create Vending Machine",
	})
}

func DeleteVendingMachine(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"message": "Delete Vending Machine",
		"id":      id,
	})
}
