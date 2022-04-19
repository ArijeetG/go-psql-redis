package helper

import (
	"fmt"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type claims struct {
	jwt.StandardClaims
}

func VerifyJwtToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		//get token
		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(400, gin.H{
				"message": "token_missing",
			})
			return
		}
		token := strings.Split(authHeader, " ")[1]
		fmt.Println(token)
		payload, _ := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
			return []byte(token), nil
		})
		if claim, ok := payload.Claims.(*claims); ok && payload.Valid {
			log.Println(claim.StandardClaims, claim)
			c.Next()
		} else {
			c.AbortWithStatusJSON(200, gin.H{
				"message": "auth_failed",
			})
		}
	}
}
