package entity

/*
结构体定义,标签解释：
	1.`gorm:"primaryKey"`: 表示设置某列为主键
	2.`gorm:"default:'test'"`: 表示为某字段设置默认值
*/
type User struct {
	Id      int    `gorm:"primaryKey"`
	Name    string `gorm:"default:'test'"`
	Age     int
	Address string
}

// 重命名表名
func (User) TableName() string {
	return "User"
}
