package generalhandler

import (
	"admin/helper"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type deleteRequestPayload struct {
	Name     string
	Password string
}

type user struct {
	name      string
	password  string
	phone     string
	address   string
	user_type string
}

func DeleteUserInfo(c *gin.Context) {
	var requestBody deleteRequestPayload
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Internal server error",
		})
	}

	if requestBody.Name == "" || requestBody.Password == "" {
		c.JSON(400, gin.H{
			"message": "missing_parameters",
		})
	}

	dbClient := helper.ConnectToDatabase()
	cacheClient := helper.GetCacheClient()

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
	if len(responseBody) == 0 {
		c.JSON(400, gin.H{
			"message": "user_not_present",
		})
		return
	}

	deleteQuery := `DELETE FROM "user" WHERE name=$1`
	_, er := dbClient.Query(deleteQuery, requestBody.Name)
	if er != nil {
		log.Println(er.Error())
		c.JSON(400, gin.H{
			"message": "Internal server error",
		})
		return
	}
	cacheClient.Del(ctx, requestBody.Name)
	c.JSON(200, gin.H{
		"message": "OK",
	})

}
