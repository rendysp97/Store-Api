package helper

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	User_id  int
	No_telp  string
	Is_admin bool
	jwt.RegisteredClaims
}

func GenerateJwt(phone string, Is_admin bool, user_id int) (string, error) {

	claim := &Claims{
		No_telp:  phone,
		Is_admin: Is_admin,
		User_id:  user_id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenStr, nil

}
