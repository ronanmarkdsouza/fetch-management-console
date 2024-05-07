package controllers

import (
	"fetch/management-console/internal/db"
	"fetch/management-console/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ViewProducts(ctx *gin.Context) {
	db, _ := db.Db()

	rows, err := db.Query(`
		SELECT * FROM products
	`)

	if err != nil {
		log.Fatal(err)
	}

	var products []models.Product

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.PID, &product.Name, &product.Description, &product.ImageUrl, &product.Price, &product.Category); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, products)
}

func ViewProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	db, _ := db.Db()

	rows, err := db.Query(`
		SELECT * FROM products
		WHERE PID = ?
	`, id)

	if err != nil {
		log.Fatal(err)
	}

	var products []models.Product

	for rows.Next() {
		var product models.Product
		if err := rows.Scan(&product.PID, &product.Name, &product.Description, &product.ImageUrl, &product.Price, &product.Category); err != nil {
			log.Fatal(err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	ctx.JSON(http.StatusOK, products)
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
