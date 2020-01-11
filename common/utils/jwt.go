package utils

import (
	. "CM-SSO/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"time"
)

type TokenClaim struct {
	AuthId uint
	UserId uint
}

//盐值最好是动态的
var key = []byte("asdasd")
func GenToken(userAuth *UserAuth) (string,error) {
	jwtExp := time.Duration(viper.GetInt("app.jwt_exp"))*time.Minute
	tokenClaim := TokenClaim{AuthId: userAuth.ID, UserId: userAuth.UserId}
	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"AuthId":tokenClaim.AuthId,
		"UserId":tokenClaim.UserId,
		"exp": time.Now().Add(jwtExp).Unix(),
		"iat": time.Now().Unix(),
	}) // Sign and get the complete encoded token as a string using the secret

	tokenString, err := token.SignedString(key)
	if err != nil{
		log.Printf("生成token失败:%s",err)
	}
	return tokenString,err
}

//失败返回  UserId,AuthId
func ParseToken(tokenStr string) *TokenClaim {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("签名方法错误: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return key, nil
	})
	if err != nil {
		log.Printf("解析token失败:%s",err)
		return nil
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return &TokenClaim{AuthId: uint(claims["UserId"].(float64)), UserId: uint(claims["AuthId"].(float64))}
	} else {
		log.Printf("解析token失败:%s",err)
	}
	return nil
}


func GetToken(c *gin.Context) string {
	return c.Request.Header.Get("token")
}
