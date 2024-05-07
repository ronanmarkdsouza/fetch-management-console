package controllers

import (
	"fetch/management-console/internal/db"
	"fetch/management-console/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewVendingMachines(ctx *gin.Context) {
	db, _ := db.Db()

	rows, err := db.Query(`
		SELECT * FROM vending_machines
	`)

	if err != nil {
		log.Fatal(err)
	}

	var vendingmachines []models.VendingMachine

	for rows.Next() {
		var vendingmachine models.VendingMachine
		if err := rows.Scan(&vendingmachine.VID, &vendingmachine.HID, &vendingmachine.TID, &vendingmachine.Location, &vendingmachine.MachineName, &vendingmachine.Status, &vendingmachine.LastServiceDate, &vendingmachine.Latitude, &vendingmachine.Longitude); err != nil {
			log.Fatal(err)
		}
		vendingmachines = append(vendingmachines, vendingmachine)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, vendingmachines)
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
