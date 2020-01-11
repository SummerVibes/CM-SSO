package model

import "time"

type UserAuth struct {
	BaseModel
	UserId uint `json:"userId" gorm:"not null"`
	AuthType string `json:"authType" gorm:"type:varchar(20);not null"`
	Identifier string `json:"identifier" gorm:"type:varchar(30);not null"`
	Credential string `json:"credential" gorm:"type:varchar(50);not null"`
	LastLoginTime time.Time `json:"lastLoginTime"`
}

func (u *UserAuth) InsertUserAuth() error{
	err := db.Create(&u).Error
	return err
}

func (u *UserAuth) UpdateLastLoginTime() error{
	u.LastLoginTime = time.Now()
	return db.Save(&u).Error
}

func CreateNewUser(user *User,auth *UserAuth) error{
	tx := db.Begin()
	err := user.InsertUser()
	auth.UserId = user.ID
	err = auth.InsertUserAuth()
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (u *UserAuth) FindUserAuthById(id uint) bool{
	return !db.First(u, id).RecordNotFound()
}

func (u *UserAuth) FindUserAuth(authType string,email string) bool {
	return !db.Where("auth_type = ? AND identifier = ?",authType,email).First(u).RecordNotFound()
}

//全量更新
func (u *UserAuth) SaveUserAuth() error {
	return db.Save(&u).Error
}

func (u *UserAuth) DeleteUserAuth() error{
	return db.Delete(&u).Error
}