package test

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"testing"
)

func TestConfig(t *testing.T)  {
	viper.SetConfigName("start-dev")
	viper.AddConfigPath("../")
	err := viper.ReadInConfig()
	//检测配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件变化:", e.Name)
	})
	if err != nil { // Handle errors reading the start file
		panic(fmt.Errorf("读取配置失败:%s \n", err))
	}
	fmt.Println()
}