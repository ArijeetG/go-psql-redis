package authHandler

import (
	"admin/helper"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type loginQueryPayload struct {
	id       string
	name     string
	password string
}

type LoginReqPayload struct {
	Name     string
	Password string
}

func Login(c *gin.Context) {
	var requestBody LoginReqPayload
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	if requestBody.Password == "" || requestBody.Name == "" {
		c.JSON(400, gin.H{
			"message": "missing_parameters",
		})
		return
	}

	//find user
	dbClient := helper.ConnectToDatabase()
	query := `SELECT "name", "user_type","password" FROM "user" WHERE name=$1`
	fmt.Println("QUERY :", query)
	resp, err := dbClient.Query(query, requestBody.Name)
	if err != nil {
		log.Panic(err.Error())
	}
	var responseBody []map[string]interface{}
	for resp.Next() {
		var b user
		if err := resp.Scan(&b.name, &b.user_type, &b.password); err != nil {
			log.Panic(err.Error())
		}
		log.Println("parsing ", b)
		err := bcrypt.CompareHashAndPassword([]byte(b.password), []byte(requestBody.Password))
		if err != nil {
			log.Println(err.Error())
			c.JSON(400, gin.H{
				"message": "invalid name/password",
			})
			return
		}
		responseBody = append(responseBody, map[string]interface{}{
			"name":      b.name,
			"user_type": b.user_type,
		})
	}
	if len(responseBody) <= 0 {
		c.JSON(400, gin.H{
			"message": "user_not_found",
		})
		return
	}

	//create jwt
	tokenString := helper.GetJwtToken()
	if tokenString == "" {
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, gin.H{
		"message":  "OK",
		"response": responseBody,
		"token":    tokenString,
	})
}
