package test

import (
	. "CM-SSO/model"
	"testing"
)

func TestUser(t *testing.T)  {
	user := User{NickName:"asd"}
	user.InsertUser()
	if user.ID<=0 {
		t.Error("插入失败")
	}
	//user.AuthId=2
	//user.SaveUser()
	//user.FindUserById(1)
	//fmt.Println(user)
	//user.DeleteUser()
}

func TestUserAuth(t *testing.T)  {
	//ua := UserAuth{BaseModel{},2,"email","1111","1111",time.Now()}
	//ua.InsertUserAuth()
}

//事务测试
func TestTransaction(t *testing.T)  {
	//db.Begin()
	//tx.Commit()
}