package service

import (
	"CM-SSO/common/cache"
	"CM-SSO/common/utils"
	"CM-SSO/model"
	"errors"
)

func (UserService) SingleSignOn(token string) (*model.LoginState,error) {
	loginState := cache.GetLoginState(token)
	if loginState!=nil {
		return loginState,nil
	}
	tokenClaim := utils.ParseToken(token)
	loginState = model.FindLoginState(tokenClaim.UserId,tokenClaim.AuthId)
	if loginState!=nil {
		return loginState,nil
	}
	return nil,errors.New("token无效")
}

func (UserService) QuitLoginState()  {

}

func (UserService) GetAllUserInfo()  {

}

func (UserService) UpdateUserInfo()  {

}

func (UserService) DeleteUser()  {

}

//登录服务,暂时不使用这种
//func (UserService) SignUp(authType string,identifier string,credential string) error {
//	//查询 判断
//	var err error
//	switch authType {
//	case EMAIL:
//		err = userService.EmailSignUp("","")
//	case QQ:
//		err = userService.QQSignUp("","")
//	}
//	return err
//}

//先判断是否已经登录
//已经登录则返回jwt
//未登录进入登录服务
/*
提供Http Header和Cookie两种方式
*/
//登录服务
//func (UserService) Login(authType string,identifier string,credential string) (string,error) {
//	//查询 判断
//	switch authType {
//	case "email":
//		return userService.EmailLogin("","")
//	case "qq":
//		return userService.QQLogin("")
//	}
//	return "",errors.New("非法请求")
//}


