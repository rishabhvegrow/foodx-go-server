package controlers

import (
    "os"
	"time"
	"github.com/golang-jwt/jwt"
)


var secret = os.Getenv("JWT_SECRET_KEY")
var jwtSecretKey = []byte(secret)


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

