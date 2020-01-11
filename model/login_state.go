package model

import "time"

type LoginState struct {
	User
	Identifier string `json:"identifier" gorm:"type:varchar(30);not null"`
	LastLoginTime time.Time `json:"lastLoginTime"`//不是单点登录的时间
}

func FindLoginState(userID uint,authId uint) *LoginState {
	var loginState LoginState
	sql := "select u.created_at, u.updated_at, u.nick_name, u.avatar,ua.identifier,ua.last_login_time from"+
		 " users u join user_auths ua on u.id = ? where ua.id = ?"
	_ = db.Raw(sql, userID,authId).Scan(&loginState)
	return &loginState
}
