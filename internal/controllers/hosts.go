package controllers

import (
	"fetch/management-console/internal/db"
	"fetch/management-console/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewHosts(ctx *gin.Context) {

	db, _ := db.Db()

	rows, err := db.Query(`
		SELECT * FROM hosts
	`)

	if err != nil {
		log.Fatal(err)
	}

	var hosts []models.Host

	for rows.Next() {
		var host models.Host
		if err := rows.Scan(&host.HID, &host.Name, &host.Email, &host.PasswordHash, &host.PhoneNumber); err != nil {
			log.Fatal(err)
		}
		hosts = append(hosts, host)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, hosts)
}

func ViewHost(ctx *gin.Context) {
	db, _ := db.Db()
	id := ctx.Param("id")

	rows, err := db.Query(`
		SELECT * 
		FROM hosts
		WHERE HID = ?
	`, id)

	if err != nil {
		log.Fatal(err)
	}

	var hosts []models.Host

	for rows.Next() {
		var host models.Host
		if err := rows.Scan(&host.HID, &host.Name, &host.Email, &host.PasswordHash, &host.PhoneNumber); err != nil {
			log.Fatal(err)
		}
		hosts = append(hosts, host)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, hosts)
}

func EditHost(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"message": "Edit Host",
		"id":      id,
	})
}

func CreateHost(ctx *gin.Context) {

	ctx.JSON(200, gin.H{
		"message": "Create Host",
	})
}

func DeleteHost(ctx *gin.Context) {
	id := ctx.Param("id")

	ctx.JSON(200, gin.H{
		"message": "Delete Host",
		"id":      id,
	})
}
