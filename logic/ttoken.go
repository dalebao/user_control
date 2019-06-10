package logic

import (
	"github.com/dalebao/user_control/pkg"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var tJwtSecret = []byte(setting.TjwtSecret)

type TClaims struct {
	UToken string `json:"u_token"`
	Uuid   string `json:"uuid"`
	Guard  string `json:"guard"`
	jwt.StandardClaims
}

func GenerateTToken(UToken, uuid, guard string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(10 * time.Minute)

	claims := Claims{
		UToken,
		uuid,
		guard,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(tJwtSecret)

	return token, err
}

func ParseTToken(token string) (*TClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &TClaims{}, func(token *jwt.Token) (interface{}, error) {
		return tJwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*TClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
