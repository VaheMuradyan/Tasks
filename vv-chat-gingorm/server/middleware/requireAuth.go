package middleware

import (
	"fmt"
	"net/http"
	"time"
	"vvchat/server/internal/user"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type MiddleWare struct {
	db *gorm.DB
}

func NewMiddleware(d *gorm.DB) *MiddleWare {
	return &MiddleWare{
		db: d,
	}
}

func (m MiddleWare) RequireAuth(c *gin.Context) {
	tokenString, err := c.Cookie("jwt")
	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("alksdjf9182374laksjdfh"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		var user1 user.User
		m.db.First(&user1, claims["sub"])

		if user1.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set("user", user1)

		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}
