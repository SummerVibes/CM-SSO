package controller

import (
	. "CM-SSO/model"
	"CM-SSO/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var emailService = service.EmailService{}

func SendEmail(c *gin.Context) {
	var signUpForm = SignUpForm{}
	if c.ShouldBindJSON(&signUpForm) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,Response{Msg: "请求非法,请检查参数"})
		return
	}
	err := emailService.SendSignUpEmail(signUpForm.Identifier, signUpForm.Credential)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,Response{Msg: "发送邮件失败"})
		return
	}
	c.JSON(http.StatusOK,Response{Msg: "成功"})
}

func EmailSignUp(c *gin.Context) {
	id := c.Query("i")
	pwd := c.Query("c")
	ticket := c.Query("t")
	if id=="" || pwd=="" || ticket==""{
		c.AbortWithStatusJSON(http.StatusBadRequest,Response{Msg: "链接无效"})
		return
	}
	err := emailService.EmailSignUp(id,pwd,ticket)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusBadRequest,Response{Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK,Response{Msg: "成功"})
}

func EmailLogin(c *gin.Context) {
	var loginForm = LoginForm{}
	if c.ShouldBindJSON(&loginForm) != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,Response{Msg: "请求非法,请检查参数"})
		return
	}
	token, err := emailService.EmailLogin(loginForm.Identifier,loginForm.Credential)
	if err != nil{
		c.AbortWithStatusJSON(http.StatusInternalServerError,Response{Msg: err.Error()})
		return
	}
	c.JSON(http.StatusOK,ResponseData{Msg: "成功",Data:token})
}
