package cache

import (
	"CM-SSO/model"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

var (
	cacheInstance *cache.Cache
	cacheLock sync.Once
)

func init()  {
	GetInstance()
}

//单例模式
func GetInstance() *cache.Cache{
	exp := time.Duration(viper.GetInt("app.exp"))*time.Minute
	clearInterval := time.Duration(viper.GetInt("app.cleanup_interval"))*time.Minute
	cacheLock.Do(func() {
		cacheInstance = cache.New(exp, clearInterval)
	})
	return cacheInstance
}

func GetLoginState(key string) *model.LoginState {
	loginState,found := cacheInstance.Get(key)
	if !found {
		return nil
	}else {
		loginState := loginState.(model.LoginState)
		log.Printf("从缓存中获取登录态:%v",loginState)
		return &loginState
	}
}

func SetLoginState(key string,loginState *model.LoginState)  {
	err := cacheInstance.Add(key,*loginState,cache.DefaultExpiration)
	if err != nil {
		log.Printf("缓存写入失败:%s",err.Error())
	}
}
