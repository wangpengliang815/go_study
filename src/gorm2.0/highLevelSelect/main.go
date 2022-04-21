package main

import (
	"fmt"
	. "goProject/gorm2.0/common"
)

var db = CreateDbConn()

func main() {
	SubQuery()
}

// SelectField 自动选择字段,其实就是定义了一个小的结构体gorm内部反射进行名称匹配赋值
func SelectField() {
	type UserModel struct {
		Id          int    `gorm:"column:Id"`
		Name        string `gorm:"column:Name"`
		PhoneNumber string `gorm:"column:PhoneNumber"`
	}
	var vm []UserModel
	db.Model(new(User)).Debug().Find(&vm)
	// SELECT "User"."Id","User"."Name","User"."PhoneNumber" FROM "User"
	fmt.Printf("%#v", vm)
}

// SubQuery 子查询
func SubQuery() {
	var users []User
	db.Where("age > (?)", db.Table("User").Select("AVG(age)")).Debug().Find(&users)
	// SELECT "User"."Id","User"."Name","User"."Age","User"."PhoneNumber","User"."Address","User"."CreateTime" FROM "User" WHERE age > (SELECT AVG(age) FROM "User")
	fmt.Println(users)

	subQuery := db.Select("AVG(age)").Where("name LIKE ?", "%wang%").Table("User")
	db.Select("AVG(age) as avgAge").Group("name").Having("AVG(age) > (?)", subQuery).Debug().Find(&users)
	// SELECT AVG(age) as avgAge FROM "User" GROUP BY "name" HAVING AVG(age) > (SELECT AVG(age) FROM "User" WHERE name LIKE '%wang%')
}
