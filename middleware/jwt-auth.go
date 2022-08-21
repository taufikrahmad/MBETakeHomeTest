package middleware

import (
	"MBETakeHomeTest/helper"
	"MBETakeHomeTest/services"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
)

func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process your request","No Token Found",nil)
			c.AbortWithStatusJSON(http.StatusBadRequest,response)
			return
		}

		token,err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim [user id] ",claims["user_id"])
			log.Println("Claim [ExpiresAt] ",claims["ExpiresAt"])
			log.Println("Claim [issuer] ",claims["issuer"])
		}else{
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid",err.Error(),nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized,response)
		}
	}
}
