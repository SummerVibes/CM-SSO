package service

import (
	"CM-SSO/common/cache"
	"CM-SSO/common/utils"
	"CM-SSO/model"
	"errors"
	"github.com/spf13/viper"
	"log"
	"time"
)

func (EmailService) EmailSignUp(identifier string,credential string,ticket string) error{
	var userAuth = new(model.UserAuth)
	if userAuth.FindUserAuth(EMAIL,identifier){
		return errors.New("用户已存在")
	}
	c := cache.GetInstance()
	cTicket,found:= c.Get(identifier)
	cTicket = cTicket.(string)
	if !found||cTicket!=ticket{
		log.Printf("found is %v,ticket in cache is %s",found,ticket)
		return errors.New("邮箱验证失败")
	}
	var user = new(model.User)
	userAuth.Credential = credential
	userAuth.Identifier = identifier
	userAuth.AuthType = EMAIL
	err := model.CreateNewUser(user,userAuth)
	if err != nil {
		return err
	}
	return nil
}

//发送邮件,用户确认后再注入用户信息,邮箱链接有效期为30min
func (EmailService) SendSignUpEmail(identifier string,credential string) error{
	var userAuth = new(model.UserAuth)
	if userAuth.FindUserAuth(EMAIL,identifier){
		return errors.New("用户已存在")
	}
	credential = utils.Encrypt(credential,identifier)
	//生成认证字符串
	ticket := utils.GenRandomString(16)
	c := cache.GetInstance()
	c.Set(identifier,ticket,30*time.Minute)
	//构造链接
	body := viper.GetString("app.name")+"mail?i="+identifier+"&c="+credential+"&t="+ticket
	log.Printf("链接:%s",body)
	err := utils.SendMail([]string{identifier},body)
	if err != nil {
		log.Printf("邮件发送失败:%s",err.Error())
	}

	return err
}

func (EmailService) EmailLogin(identifier string,credential string) (string,error) {
	userAuth := new(model.UserAuth)
	if !userAuth.FindUserAuth(EMAIL,identifier)||!utils.Verify(credential,identifier,userAuth.Credential){
		return "",errors.New("用户名或密码不正确")
	}
	err := userAuth.UpdateLastLoginTime()
	if err != nil {
		log.Printf("更新登录时间失败:%s",err)
		return "",errors.New("服务器异常")
	}
	//生成token
	token,err := utils.GenToken(userAuth)
	if err != nil {
		return "",errors.New("生成token失败")
	}
	//将登录态放入缓存
	user := new(model.User)
	if !user.FindUserById(userAuth.UserId){
		return "",errors.New("无用户数据存在")
	}

	loginState := model.FindLoginState(user.ID,userAuth.ID)
	cache.SetLoginState(token,loginState)
	return token,nil
}