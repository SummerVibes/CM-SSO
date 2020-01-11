package model

type User struct {
	BaseModel
	NickName string `json:"nickname" gorm:"type:varchar(100);default:'快起个名字吧'"`
	Avatar string `json:"avatar" gorm:"type:varchar(255)"`
	//AuthId uint `json:"authId"`
}


func (u *User) InsertUser() error{
	err := db.Create(&u).Error
	return err
}

func (u *User) FindUserById(id uint) bool{
	return !db.First(u, id).RecordNotFound()
}

//全量更新
func (u *User) SaveUser() error {
	return db.Save(&u).Error
}

func (u *User) DeleteUser() error {
	return db.Delete(&u).Error
}