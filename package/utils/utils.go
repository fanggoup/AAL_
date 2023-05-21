package utils

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

type Claims struct{
	Id       uint   `json:"id"`
	UserName string `json:"user_name"`
	Authority int    `json:"authority"`
	jwt.StandardClaims
}

// 签发token
func GenerateToken(id uint,username string ,authority int)(string,error){
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour)
	claims := Claims{
		Id: id,
		UserName: username,
		Authority: authority,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer: "AAL_time",
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,err := tokenClaims.SignedString(jwtSecret)
	return token,err
}

// 验证token
func ParseToken(token string) (*Claims, error) {
	// 去掉 token 字符串中的 bearer 前缀
	token = strings.ReplaceAll(token, "Bearer ", "")
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, fmt.Errorf("无效的token: %s", err)
	}
	if tokenClaims == nil {
		return nil, errors.New("token解析失败")
	}
	if !tokenClaims.Valid {
		return nil, errors.New("token无效")
	}
	claims, ok := tokenClaims.Claims.(*Claims)
	if !ok {
		return nil, errors.New("无效的claims")
	}
	return claims, nil
}