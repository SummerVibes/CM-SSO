package model

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	. "github.com/spf13/viper"
	"log"
	"os"
	"time"
)

var db *gorm.DB

type BaseModel struct {
	ID        uint `json:"-" gorm:"primary_key;AUTO_INCREMENT"`//
	CreatedAt time.Time `json:"createdAt"`//将被设置为当前时间
	UpdatedAt time.Time `json:"updatedAt"`//将被设置为当前时间
	DeletedAt *time.Time `json:"-" sql:"index"`//具有该字段只能软删除
}

//func (j *JSONTime) UnmarshalJSON(b []byte) error {
//	var s string
//	if err := json.Unmarshal(b, &s); err != nil {
//		return err
//	}
//	*j = time.Time(s).Format("2006-01-02 15:04:05")
//	return nil
//}
//自定义json反序列化器
//func (j time.Time) MarshalJSON() ([]byte, error) {
//	var stamp = fmt.Sprintf("\"%s\"", j.Format("2006-01-02 15:04:05"))
//	return []byte(stamp), nil
//}

func ConfigureDB() {
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetString("mysql.user"),
		GetString("mysql.password"),
		GetString("mysql.host"),
		GetString("mysql.database"),
	)

	var err error
	db, err = gorm.Open("mysql", conn)
	if err!=nil{
		log.Fatalf("connect database failed:%s\n statement is :%s",err.Error(),conn)
		return
	}
	db.DB().SetMaxIdleConns(GetInt("mysql.max_idle_conns"))
	db.DB().SetMaxOpenConns(GetInt("mysql.max_open_conns"))
	db.SetLogger(log.New(os.Stdout, "\r\n", 0))
	//defer db.Close()
	db.AutoMigrate(&User{},&UserAuth{})
}