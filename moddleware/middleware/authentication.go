package middleware

import (
	"moddleware/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Authentications() gin.HandlerFunc {
	return func(c *gin.Context) {
		// jika error maka batalkan sehingga tidak mencapai header
		verifyToken, err := helpers.VerifyToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "unauthenticated",
				"message": err.Error(),
			})
			return
		}

		// bila benar maka data user akan diambiil
		c.Set("userData", verifyToken)
		c.Next()
	}
}
