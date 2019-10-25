package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

func TestModels(t *testing.T) {
	db, err := gorm.Open("mysql", "root:123456@/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Errorf("cannot connect the database")
	}
	// 默认使用后关闭数据库
	defer db.Close()

	t.Run("testing user CRUD", func(t *testing.T) {
		db.AutoMigrate(&User{}, &MarkDown{}) // 数据库中表单和模型同步
		user := &User{
			Model: gorm.Model {
				ID: 6,
			},
		}

		user.ID = uint(6)
		db.Where("ID = ?", user.ID).Find(&user)


		// newUser := db.Create(&user)
		fmt.Println("The new one is", user)

		// db.First(&first_user, "name = ?", "test")

		// 读取
		// var first_user User
		// db.First(&first_user, 1) // 查询id为1的product
		// fmt.Println(first_user)
		// db.Delete(&product)

	})
}

func TestStruct(t *testing.T) {
	t.Run("testing struct source and value", func(t *testing.T) {
		var u User
		u.Name = "123132"
		var u_p *User = &u
		fmt.Println(u, u_p)
	})
}
