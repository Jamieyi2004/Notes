package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "username:password@(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	// serInfo{}：这是一个复合字面量（Composite Literal），用于创建一个 UserInfo 类型的实例。复合字面量允许你以一种紧凑的方式初始化数据结构（如结构体、数组、切片和映射）。当你使用 {} 并且不传递任何初始值时，所有字段都将被赋予其类型的零值。
    // &：这是取地址操作符，在这里用来获取新创建的 UserInfo 实例的地址。这意味着你得到的是一个指向该实例的指针而不是实例本身。这对于需要修改原对象或避免复制大型结构体的情况特别有用。
	db.AutoMigrate(&UserInfo{})

	u1 := UserInfo{1, "七米", "男", "篮球"}
	u2 := UserInfo{2, "沙河娜扎", "女", "足球"}
	// 创建记录
	db.Create(&u1)
	db.Create(&u2)
	// 查询
	var u = new(UserInfo)
	db.First(u)
	fmt.Printf("%#v\n", u)

	var uu UserInfo
	db.Find(&uu, "hobby=?", "足球")
	fmt.Printf("%#v\n", uu)

	// 更新
	db.Model(&u).Update("hobby", "双色球")
	// 删除
	db.Delete(&u)
}
