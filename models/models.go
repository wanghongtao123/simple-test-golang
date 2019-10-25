package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//为每个模型添加CRUD操作

type User struct {
	gorm.Model
	Name  string `gorm:"size:10"`
	State bool
}

// CRUD增删改查

func (u *User) CreateUser(db *gorm.DB) {
	// 编辑user的方法
	if db.NewRecord(u) {
		db.Create(&u)
	}
}


func (u *User) ReadUser(db *gorm.DB) {
	if u.ID == 0 {

	}
}


func (u *User) UpdateUser(db *gorm.DB) {
	
}

func (u *User) DeleteUser(db *gorm.DB) {
	
}

type MarkDown struct {
	gorm.Model
	Name  string `gorm:"size:10"`
	User  User `gorm:"foreignkey:UserRefer"` // 使用 UserRefer 作为外键
  	UserRefer uint
}


