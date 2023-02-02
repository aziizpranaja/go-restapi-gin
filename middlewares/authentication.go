package middlewares

import (
	"fmt"
	"go-restapi-gin/models"
	"go-restapi-gin/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func CheckCredetial(c *gin.Context){
	// Get the cookie off req
	// tokenString, err := c.Cookie("Authorization")
	const BEARER_SCHEMA = "Bearer "
	authHeader := c.GetHeader("Authorization")
	if len(authHeader) == 0 {
		fmt.Println("test")
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	}
	
	tokenString := authHeader[len(BEARER_SCHEMA):]

	// Decode/validate it
	// Parse takes the token string and a function for looking up the key. The latter is especially
	token, err := services.JWTAuthService().ValidateToken(tokenString)

	if token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		//Check the Exp 
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
		}

		// Find the user with token sub
		var user models.User

		models.DB.First(&user, claims["sub"])

		if user.Id == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
		}

		//Attach to req
		c.Set("user", user)

		// Continue
		c.Next()
	} else {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
		})
	}
}