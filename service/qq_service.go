package service


func (QQService) QQSignUp(identifier string,credential string) error {
	return nil
}

//前端获取authorizationCode
//根据authorizationCode,appid,appkey等获取accessToken
//根据accessToken获取openid
//使用openid和appid等调用接口获取用户信息
func (QQService) QQLogin(authorizationCode string) (string,error) {
	return "",nil
}
