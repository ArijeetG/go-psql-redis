package authHandler

import (
	"admin/helper"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type SignupReqPayload struct {
	Name      string
	Password  string
	User_type string
}

type user struct {
	name      string
	user_type string
	password  string
}

func Signup(c *gin.Context) {
	dbClient := helper.ConnectToDatabase()

	var request SignupReqPayload
	if err := c.ShouldBindJSON(&request); err != nil {
		log.Panic(err.Error())
		return
	}
	fmt.Println(request)

	//Check if user is present
	findQuery := `SELECT "name" FROM "user" WHERE name=$1`
	resp, err := dbClient.Query(findQuery, request.Name)
	if err != nil {
		log.Panic(err.Error())
		return
	}
	var queryPayload []string
	for resp.Next() {
		var u string
		if err := resp.Scan(&u); err != nil {
			log.Panic(err.Error())
		}
		queryPayload = append(queryPayload, string(u))
	}
	if len(queryPayload) > 0 {
		c.JSON(400, gin.H{
			"message": "user_already_exists",
		})
		return
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(request.Password), 14)
	if err != nil {
		log.Panic(err.Error())
		c.JSON(400, gin.H{
			"message": "something_went_wrong",
		})
		return
	}

	hashedPassword := string(bytes)
	insertQuery := `INSERT INTO "user" ("name","password","user_type") VALUES ('` + request.Name + `','` + hashedPassword + `','` + request.User_type + `');`
	_, er := dbClient.Exec(insertQuery)
	if er != nil {
		log.Panic(er.Error())
		c.JSON(500, gin.H{
			"message": "something_went_wrong",
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
		"message": "OK",
		"token":   tokenString,
	})

}
