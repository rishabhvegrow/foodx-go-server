package controlers

import (
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)



var jwtSecretKey = []byte("jwdnjkdsckjdckjsd123456gfdtyu8hvghfcjds")

// JWT
func generateToken(userID uint) (string, error) {
    tokenClaims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 24).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
    tokenString, err := token.SignedString(jwtSecretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func JWTAuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        tokenString := c.GetHeader("Authorization")

        if tokenString == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
            c.Abort()
            return
        }

        parts := strings.Split(tokenString, " ")
        if len(parts) < 2 || parts[0] != "Bearer" { 
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }


        tokenString = parts[1]
        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
            return jwtSecretKey, nil
        })


        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
            c.Abort()
            return
        }

        userID, ok := claims["user_id"].(float64)
        if !ok {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID in token"})
            c.Abort()
            return
        }

        c.Set("user_id", uint(userID))
        c.Next()
    }
}