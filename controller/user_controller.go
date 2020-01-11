package controller

import (
	"CM-SSO/common/utils"
	. "CM-SSO/model"
	"CM-SSO/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)
var userService = service.UserService{}

//单点登录控制器
func SingleSignOn(c *gin.Context) {
	token := utils.GetToken(c)
	if token=="" {
		token,_ = c.Cookie("token")
	}
	if token==""  {
		c.AbortWithStatusJSON(http.StatusUnauthorized,Response{Msg: "认证失败,请重新登录"})
		c.Redirect(http.StatusMovedPermanently, viper.GetString("app.login_page"))
		return
	}
	loginState,err := userService.SingleSignOn(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized,Response{Msg: "单点登录失败"+err.Error()})
		c.Redirect(http.StatusMovedPermanently, viper.GetString("app.login_page"))
		return
	} else {
		c.JSON(http.StatusOK,ResponseData{Msg: "单点登录成功", Data: loginState})
		return
	}
}

func QuitLoginState(c *gin.Context)  {

}

func GetAllUserInfo(c *gin.Context)  {

}

func UpdateUserInfo(c *gin.Context)  {

}

func DeleteUser(c *gin.Context)  {

}