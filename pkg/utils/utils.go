package utils

import (
	"regexp"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ParseInt64(value interface{}) int64 {
	switch v := value.(type) {
	case string:
		val, _ := strconv.Atoi(v)
		if val >= 0 {
			return int64(val)
		}
		return 0
	case float64:
		if v >= 0 {
			return int64(v)
		}
		return 0
	case uint:
		return int64(v)
	case int:
		if v >= 0 {
			return int64(v)
		}
		return 0
	case int64:
		return v
	default:
		return 0
	}
}

type JWTClaims struct {
	UserId int64 `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWTToken(userId int64, secretKey string) (string, error) {
	claims := JWTClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}

func ValidateEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
