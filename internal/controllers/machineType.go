package controllers

import (
	"fetch/management-console/internal/db"
	"fetch/management-console/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewMachineTypes(ctx *gin.Context) {

	db, _ := db.Db()

	rows, err := db.Query(`
		SELECT * FROM machine_type
	`)

	if err != nil {
		log.Fatal(err)
	}

	var machinetypes []models.MachineType

	for rows.Next() {
		var machinetype models.MachineType
		if err := rows.Scan(&machinetype.TID, &machinetype.TypeName, &machinetype.Capacity, &machinetype.Refrigeration); err != nil {
			log.Fatal(err)
		}
		machinetypes = append(machinetypes, machinetype)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, machinetypes)
}

func ViewMachineType(ctx *gin.Context) {
	id := ctx.Param("id")

	db, _ := db.Db()

	rows, err := db.Query(`
		SELECT * FROM machine_type
		WHERE TID =?
	`, id)

	if err != nil {
		log.Fatal(err)
	}

	var machinetypes []models.MachineType

	for rows.Next() {
		var machinetype models.MachineType
		if err := rows.Scan(&machinetype.TID, &machinetype.TypeName, &machinetype.Capacity, &machinetype.Refrigeration); err != nil {
			log.Fatal(err)
		}
		machinetypes = append(machinetypes, machinetype)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, machinetypes)
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
