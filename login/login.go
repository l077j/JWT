package login

import (
	"net/http"
	"time"

	jwt2 "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/l077j/JWT/jwt"
)

func Login(c *gin.Context) {
	data := jwt.Data{
		Name:   "Kol",
		Age:    24,
		Gender: "male",
		StandardClaims: jwt2.StandardClaims{
			ExpiresAt: time.Now().Unix() + 60*60*24*30,
			IssuedAt:  time.Now().Unix(),
			Issuer:    "Kol",
		},
	}

	sign, err := jwt.Sign(data)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": sign,
	})
}
