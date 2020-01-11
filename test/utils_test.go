package test

import (
	"testing"
)

func TestSendMail(t *testing.T)  {
	//utils.SendMail([]string{"1144272434@qq.com"})
}

func TestJwt(t *testing.T)  {
	//base := model.BaseModel{1,time.Now(),time.Now(),nil}
	//token := utils.GenToken(&model.UserAuth{base,1,"email","sdasda","asdadsd",time.Now()})
	//utils.ParseToken(token)
}

//基准测试,框架自行决定该运行多少次
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		//utils.Hello()
	}
}

