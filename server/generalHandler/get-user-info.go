package generalhandler

import (
	"admin/helper"
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

type requestPayload struct {
	Name string
}

type userPayload struct {
	name      string
	user_type string
	address   string
	phone     string
}

var ctx = context.Background()

func GetUserInfo(c *gin.Context) {
	var requestBody requestPayload
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		log.Println(err.Error())
		c.JSON(400, gin.H{
			"message": "Internal server error",
		})
		return
	}

	if requestBody.Name == "" {
		c.JSON(400, gin.H{
			"message": "missing_parameter",
		})
		return
	}

	dbClient := helper.ConnectToDatabase()
	cacheClient := helper.GetCacheClient()
	fmt.Println("name: ", requestBody.Name)
	isDataPresent := helper.GetCache(cacheClient, ctx, requestBody.Name)
	fmt.Println(isDataPresent)
	switch v := isDataPresent.(type) {
	case bool:
		if v {
			// IF VALUE IS NOT PRESENT IN CACHE
			query := `SELECT "name", "user_type", "address", "phone" FROM "user" WHERE name=$1`
			fmt.Println("QUERY: ", query)
			resp, err := dbClient.Query(query, requestBody.Name)
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
				fmt.Println(u.name, u.user_type)
				responseBody = append(responseBody, map[string]interface{}{
					"name":      u.name,
					"user_type": u.user_type,
				})
			}
			marshalRequestBody, err := json.Marshal(responseBody)
			if len(responseBody) > 0 {
				r := helper.SetCache(cacheClient, ctx, requestBody.Name, marshalRequestBody)
				if r == 0 {
					c.JSON(500, gin.H{
						"message": "Internal server error",
					})
					return
				}

			}
			c.JSON(200, gin.H{
				"message":  "OK",
				"response": responseBody,
			})
			return
		} else {
			c.JSON(500, gin.H{
				"message": "Internal server error",
			})
		}
	case string:
		fmt.Println(string(v))
		c.JSON(200, gin.H{
			"message":  "OK",
			"response": v,
		})
	default:
		c.JSON(500, gin.H{
			"message": "Something went wrong",
		})
	}

}
