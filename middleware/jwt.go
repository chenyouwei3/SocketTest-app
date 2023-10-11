package middleware

import (
	"SocketTest-app/model"
	"SocketTest-app/utils"
	"github.com/golang-jwt/jwt"
	"time"
)

type CustomClaims struct {
	User model.User `json:"user"`
	jwt.StandardClaims
}

var SigningKey = []byte("Cyw-github")

func CreateToken(user model.User) (string, error) {
	//获取token，前两部分
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{User: user,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix(),           //签名生效时间
			ExpiresAt: time.Now().Unix() + 60*60*2, //2小时过期
			Issuer:    "cyw",                       //签发人，
		},
	})
	//根据密钥生成加密token，token完整三部分
	tokenString, err := token.SignedString(SigningKey)
	if err != nil {
		return "", err
	}
	//存入redis
	err = utils.Redis{}.SetValue(tokenString, user.Account, 60*60*24*time.Second)
	return tokenString, err
}
