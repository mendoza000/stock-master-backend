package routes

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mendoza000/stockmaster/helpers"
	"github.com/mendoza000/stockmaster/models"
)

func UserRoutes(r *gin.Engine, db *sql.DB) {
	r.GET("api/users", func(c *gin.Context){
		rows, err := db.Query("SELECT * FROM users")
		if err != nil {
			fmt.Println(err.Error())
		}
		// defer db.Close()
	
		var users []models.User
		for rows.Next() {
			var user models.User
			// revisamos si cada uno de los productos contiene la estructura de Prodgin
			if err := rows.Scan(&user.ID, &user.Username, &user.Mail, &user.Password, &user.UserRange); err != nil {
				fmt.Println(err.Error())
			}
			// si esta bien el producto lo agregamos a la lista
			users = append(users, user)
		}
		// revisamos si hay errores en la row
		if err := rows.Err(); err != nil {
			fmt.Println(err.Error())
		}
		// devolvemos la lista en tipo json
		c.JSON(http.StatusOK, users)
	})

	r.POST("api/users", func(c *gin.Context) {
		var body models.User
		if err := c.BindJSON(&body);err!=nil {
			fmt.Println("")
			fmt.Println(err.Error())
			fmt.Println("")
			c.JSON(500, gin.H{"error": "error al leer el body!"})
			return
		}

		err2 := helpers.ValidateBodyUser(body)
		if err2 != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2})
			return
		}

		tempPass, err :=helpers.HashPassword(body.Password)
		if err != nil {
			fmt.Println("")
			fmt.Println(err.Error())
			fmt.Println("")
			c.JSON(500, gin.H{"error": "error al encriptar la contraseña!"})
			return
		}
		body.Password = tempPass

		_, err = db.Query("INSERT INTO users (username, mail, password, user_range) VALUES (?, ?, ?, ?)", body.UserRange, body.Mail, body.Password, body.UserRange)
		if err != nil {
			fmt.Println("")
			fmt.Println(err.Error())
			fmt.Println("")
			c.JSON(500, gin.H{"error": "error al crear el usuario!"})
			return
		}

		c.JSON(http.StatusCreated, body)
	})

	r.POST("api/users/login", func(c *gin.Context) {
		var body models.UserLogin
		var user models.User
		if err := c.BindJSON(&body);err!=nil {
			fmt.Println("")
			fmt.Println(err.Error())
			fmt.Println("")
			c.JSON(500, gin.H{"error": "error al leer el body!"})
			return
		}

		err2 := helpers.ValidateBodyUserLogin(body)
		if err2 != "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": err2})
			return
		}

		rows, err := db.Query("SELECT * FROM users WHERE username = ?", body.Username)
		if err != nil {
			fmt.Println("")
			fmt.Println(err.Error())
			fmt.Println("")
			c.JSON(http.StatusBadGateway, gin.H{"error": "error en la consulta!"})
			return
		}

		for rows.Next() {
			var tempUser models.User
			if err := rows.Scan(&tempUser.ID, &tempUser.Username, &tempUser.Mail, &tempUser.Password, &tempUser.UserRange); err != nil {
				fmt.Println("")
				fmt.Println(err.Error())
				fmt.Println("")
				c.Status(http.StatusBadGateway)
				return
			}
			user = tempUser
		}

		c.JSON(200, user)

	})
}