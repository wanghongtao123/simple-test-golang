package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//为每个模型添加CRUD操作

type User struct {
	gorm.Model
	Name string `gorm:"size:10"`
	State bool
}