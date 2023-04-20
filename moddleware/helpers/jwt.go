// untuk membuat token jwt

package helpers

import (
	"errors"
	"log"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secretKey = "secret" // jng disini nnt pake env

func GenerateToken(id uint, email string, level string) string {
	// ambil data untuk enkripsi token
	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"level": level,
	}

	// token tahap awal/mentah
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// ubah oken jadi string
	signedToken, err := parseToken.SignedString([]byte(secretKey))
	if err != nil {
		log.Fatal(err)
	}

	return signedToken
}

// untuk bearer token
func VerifyToken(c *gin.Context) (interface{}, error) {
	errResp := errors.New("Failed to verify token")
	headerToken := c.Request.Header.Get("Authorization")
	isBearer := strings.HasPrefix(headerToken, "Bearer")

	if !isBearer {
		return nil, errResp
	}
	// token di indeks 1 dengan separator spasi
	stringToken := strings.Split(headerToken, " ")[1]

	// anonymous function
	token, _ := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // check signature method
		if !ok {
			return nil, errResp
		}

		return []byte(secretKey), nil
	})

	// validate
	_, ok := token.Claims.(jwt.MapClaims) // adakah error ketika convert token claim ke map claims
	if !ok && !token.Valid {
		return nil, errResp
	}

	return token.Claims.(jwt.MapClaims), nil
}
