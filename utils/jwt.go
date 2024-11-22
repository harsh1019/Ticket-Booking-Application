package utils
import (
	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(claims jwt.Claims, method jwt.SigningMethod, secret string) (string, error) {
	token := jwt.NewWithClaims(method, claims)
	return token.SignedString([]byte(secret))
}