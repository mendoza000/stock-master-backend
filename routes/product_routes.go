package routes

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mendoza000/stockmaster/helpers"
	"github.com/mendoza000/stockmaster/models"
)

func ProductsRoutes(r *gin.Engine, db *sql.DB){
	r.GET("/api/products", func (c *gin.Context) {
		rows, err := db.Query("SELECT * FROM products")
		if err != nil {
			fmt.Println(err.Error())
		}
		// defer db.Close()
	
		var products []models.Product
		for rows.Next() {
			var product models.Product
			// revisamos si cada uno de los productos contiene la estructura de Prodgin
			if err := rows.Scan(&product.ID, &product.Title, &product.Brand, &product.Stock, &product.Price, &product.Details, &product.Amount); err != nil {
				fmt.Println(err.Error())
			}
			// si esta bien el producto lo agregamos a la lista
			products = append(products, product)
		}
		// revisamos si hay errores en la row
		if err := rows.Err(); err != nil {
			fmt.Println(err.Error())
		}
		// devolvemos la lista en tipo json
		c.JSON(http.StatusOK, products)
	})

	r.GET("api/products/:id", func(c *gin.Context){
		var product models.Product
		id := c.Param("id")
		rows, err := db.Query("SELECT * FROM products WHERE id = ?", id)
		if err != nil {
			fmt.Println("error en la consulta")
			c.JSON(http.StatusNotFound, "")
		}
		for rows.Next() {
			var tempProduct models.Product
			if err := rows.Scan(&tempProduct.ID, &tempProduct.Title, &tempProduct.Brand, &tempProduct.Stock, &tempProduct.Price, &tempProduct.Details, &tempProduct.Amount); err != nil {
				fmt.Println("")
				fmt.Println(err.Error())
				fmt.Println("")
			}

			product = tempProduct
		}

		if err := rows.Err(); err != nil {
			fmt.Println("")
			fmt.Println(err.Error())
			fmt.Println("")
		}

		if product.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Id not found"})
			return
		}

		c.JSON(http.StatusOK, product)
	})

	r.POST("api/products", func (c *gin.Context) {
		var body models.Product
		
		if err := c.BindJSON(&body);err!=nil {
			fmt.Println("")
			fmt.Println(err.Error())
			fmt.Println("")
			c.JSON(500, gin.H{"error": "error al leer el body!"})
			return
		}

		if err := helpers.ValidateBodyProduct(body); err != ""{
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		_, err := db.Exec("INSERT INTO products (title, brand, stock, price, details, amount) VALUES (?, ?, ?, ?, ?, ?)", body.Title, body.Brand, body.Stock, body.Price, body.Details, body.Amount)
		if err != nil {
			fmt.Println("")
			fmt.Println(err.Error())
			fmt.Println("")
			c.JSON(500, gin.H{"err": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, &body)
	})

	r.DELETE("/api/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		_, err := db.Query("DELETE FROM products WHERE id = ?", id)
		if err != nil {
			fmt.Println("error en la consulta")
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.Status(http.StatusOK)
	})

	r.PUT("/api/products/:id", func(c *gin.Context) {
		id := c.Param("id")
		// obtenemos el body y lo validamos
		var body models.Product
		
		if err := c.BindJSON(&body);err!=nil {
			fmt.Println("")
			fmt.Println(err.Error())
			fmt.Println("")
			c.JSON(500, gin.H{"error": "error al leer el body!"})
			return
		}

		if err := helpers.ValidateBodyProduct(body); err != ""{
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		// actualizamos el producto
		_, err := db.Query("UPDATE products SET title = ?, brand = ?, stock = ?, price = ?, details = ?, amount = ? WHERE id = ?", body.Title, body.Brand, body.Stock, body.Price, body.Details, body.Amount, id)
		
		if err != nil {
			fmt.Println("")
			fmt.Println(err.Error())
			fmt.Println("")
			c.JSON(http.StatusBadRequest, gin.H{"error": "error al actualizar el producto!"})
			return
		}
		// parseamos y agramos el id al producto
		body.ID, err = strconv.Atoi(id)
		if err != nil {
			fmt.Println("")
			fmt.Println(err.Error())
			fmt.Println("")
			c.JSON(500, gin.H{"error": "error al parsear el id!"})
			return
		}

		c.JSON(http.StatusOK, body)
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
	})
}