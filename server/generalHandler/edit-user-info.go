package generalhandler

import (
	"admin/helper"
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

type payload struct {
	Name    string
	Address string
	Phone   string
}

func EditUserInfo(c *gin.Context) {
	var requestBody payload
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"error": "Internal server error",
		})
		return
	}

	if requestBody.Name == "" || requestBody.Address == "" || requestBody.Phone == "" {
		c.JSON(400, gin.H{
			"message": "missing_parameters",
		})
	}

	dbClient := helper.ConnectToDatabase()
	cacheClient := helper.GetCacheClient()

	//update user
	query := `UPDATE "user" SET "address" = $1 , "phone" = $2 WHERE "name" = $3 `
	_, err := dbClient.Query(query, requestBody.Address, requestBody.Phone, requestBody.Name)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	//find if user present and set to cache
	findQuery := `SELECT "name", "user_type", "address", "phone" FROM "user" where name=$1`
	resp, err := dbClient.Query(findQuery, requestBody.Name)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}
	var responseBody []map[string]interface{}
	for resp.Next() {
		var u userPayload
		if err := resp.Scan(&u.name, &u.user_type, &u.address, &u.phone); err != nil {
			log.Panic(err.Error())
			return
		}
		responseBody = append(responseBody, map[string]interface{}{
			"name":      u.name,
			"phone":     u.phone,
			"address":   u.address,
			"user_type": u.user_type,
		})
	}
	if len(responseBody) > 0 {
		marshalResponseBody, err := json.Marshal(responseBody)
		if err != nil {
			log.Println(err.Error())
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		r := helper.SetCache(cacheClient, ctx, requestBody.Name, marshalResponseBody)
		if r == 0 {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
			return
		}
		c.JSON(299, gin.H{
			"message":  "OK",
			"response": responseBody,
		})
	} else {
		c.JSON(400, gin.H{
			"message": "user_not_present",
		})
	}

}
