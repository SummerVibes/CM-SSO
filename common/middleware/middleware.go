package middleware

import (
	"CM-SSO/common/cache"
	"CM-SSO/common/utils"
	"CM-SSO/model"
	"github.com/gin-gonic/gin"
	cache2 "github.com/patrickmn/go-cache"
	"net/http"
)

func CheckLoginState() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetToken(c)
		if token != "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized,"无token,请重新登录")
		}
		loginState := cache.GetLoginState(token)
		if loginState!=nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized,"token验证失败,请重新登录")
		}
		c.Next()
	}
}

func CheckMaliciousLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginForm model.LoginForm
		err := c.ShouldBindJSON(loginForm)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, model.Response{Msg: "用户名或密码错误"})
		}
		ca := cache.GetInstance()
		key := "fail_"+loginForm.Identifier
		failTimes,found := ca.Get(key)
		if !found {
			failTimes = 0
		}
		ca.Set(key,failTimes,cache2.DefaultExpiration)
		c.Next()
	}
}
