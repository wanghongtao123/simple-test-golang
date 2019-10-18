package models

import (
	"fmt"
	"testing"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func TestModels(t *testing.T){
	db, err := gorm.Open("mysql", "root:123456@/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil { 
		t.Errorf("cannot connect the database")
	}
	// 默认使用后关闭数据库
	defer db.Close()

	t.Run("testing user CRUD", func(t *testing.T) {
		db.AutoMigrate(&User{}) // 数据库中表单和模型同步
		user := User{
			Name: "test",
			State: false,
		}

		newUser := db.Create(&user)
		fmt.Println("The new one is", newUser)
		

		// db.First(&first_user, "name = ?", "test")



		// 读取
		// var first_user User
		// db.First(&first_user, 1) // 查询id为1的product
		// fmt.Println(first_user)
		// db.Delete(&product)
		
	})
}