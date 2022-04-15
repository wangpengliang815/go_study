// Package common @title gorm-结构体定义
// @description
// @author wangpengliang
// @date 2022-04-14 23:46:22
package common

import "gorm.io/gorm"

/*
	1.`gorm:"primaryKey"`: 表示设置某列为主键
	2.`gorm:"default:'test'"`: 表示为某字段设置默认值;插入记录到数据库时，默认值会被用于填充值为零值的字段
	需要注意 对于声明了默认值的字段，像 0、''、false 等零值是不会保存到数据库。需要使用指针类型或 Scanner/Valuer 来避免这个问题
*/
type User struct {
	gorm.Model        // gorm里自定义的结构包含ID/CreatedAt/UpdatedAt/DeletedAt四个字段,其中DeletedAt标记当前数据是否被删除(gorm中默认使用逻辑删除)如果使用的话对应的Sql语句会自动添加相关筛选语句
	Name       string `gorm:"default:'test'"`
	Age        int
	Address    string
}

// TableName 重命名表名User否则为小写的user+复数s也就是users
func (User) TableName() string {
	return "User"
}

type Person struct {
	PersonId int
	Name     string
	Age      int
	Address  string
}

func (Person) TableName() string {
	return "Person"
}
