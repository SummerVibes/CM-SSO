package model

//tag:  Id int `json:"id" gorm:"AUTO_INCREMENT"`
type (
	ResponseData struct {
		Msg string `json:"msg"`
		Data interface{} `json:"data"`
	}
	Response struct {
		Msg string `json:"msg"`
	}
)