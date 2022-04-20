package main

import (
	"fmt"
	. "goProject/gorm2.0/common"
	"gorm.io/gorm"
)

func main() {
	ExprUpdate()
}

var db = CreateDbConn()

// DefaultUpdate 默认更新方式
func DefaultUpdate() {
	var user User
	db.Debug().First(&user, 1)
	// SELECT * FROM "User" WHERE "User"."Id" = 1 ORDER BY "User"."Id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
	if (user != User{}) {
		user.Name = "wangpengliang2"
		user.Age = 20
		db.Debug().Save(user) // Save 会保存所有的字段，即使字段是零值
		// UPDATE "User" SET "Name"='wangpengliang2',"Age"=20,"PhoneNumber"='18888888888',"Address"='beijing',"CreateTime"='2022-04-19 17:14:42.329' WHERE "Id" = 1
	}
}

// UpdateFieldSelect 更新选定字段
func UpdateFieldSelect() {
	var user User
	db.Debug().First(&user, 1)
	// 更新单个字段,Model(&user) 需要有主键
	db.Model(&user).Debug().Update("Name", "wangpengliang22")
	// UPDATE "User" SET "Name"='wangpengliang22' WHERE "Id" = 1

	// 根据条件更新单个字段
	db.Model(new(User)).Where("Id = ?", 1).Debug().Update("Name", "wangpengliang")
	// UPDATE "User" SET "Name"='wangpengliang' WHERE Id = 1

	// 通过 struct 更新多个字段，不会更新零值字段(注意Address字段并没有构建在sql语句中,因为它是字段零值)
	db.Model(new(User)).Where("Id = ?", 1).Debug().Updates(User{Name: "wangpengliang222", Age: 18, PhoneNumber: "133333333", Address: ""})
	// UPDATE "User" SET "Name"='wangpengliang222',"Age"=18,"PhoneNumber"='133333333' WHERE Id = 1

	// 通过 map 更新多个字段，零值字段也会更新(注意通过map更新零值字段也会更新，所以Address字段虽然是字段零值但是也被构建到sql中)
	db.Model(new(User)).Where("Id = ?", 1).Debug().Updates(map[string]interface{}{"Name": "wangpengliang", "Age": 20, "Address": ""})
	//  UPDATE "User" SET "Address"='',"Age"=20,"Name"='wangpengliang' WHERE Id = 1

	var user2 User
	db.Debug().First(&user2, 1)
	// 这里只会更新Name字段
	db.Model(&user2).Select("Name").Debug().Updates(map[string]interface{}{"Name": "wangpengliang", "Age": 18})
	// UPDATE "User" SET "Name"='wangpengliang' WHERE "Id" = 1

	var user3 User
	db.Debug().First(&user3, 1)
	// 这里会排除Name字段
	db.Model(user3).Omit("Name").Debug().Updates(map[string]interface{}{"Name": "wangpengliang22", "age": 20})
	// UPDATE "User" SET "age"=20 WHERE "Id" = 1

	var user4 User
	db.Debug().First(&user4, 1)
	db.Model(&user4).Select("Name", "Age").Debug().Updates(User{Name: "wangpengliang"})
	// UPDATE "User" SET "Name"='wangpengliang',"Age"=0 WHERE "Id" = 1
}

// BatchUpdate 批量更新,如果没有通过Model指定记录的主键，则GORM会执行批量更新
func BatchUpdate() {
	db.Model(User{}).Where("id = ?", 1).Debug().Updates(User{Name: "wangpengliang", Age: 18})
	// UPDATE "User" SET "Name"='wangpengliang',"Age"=18 WHERE id = 1

	db.Table("User").Where("id IN (?)", []int{1, 2}).Debug().Updates(map[string]interface{}{"Name": "Name", "Age": 20})
	// UPDATE "User" SET "Age"=20,"Name"='Name' WHERE id IN (1,2)

	//如果在没有任何条件的情况下执行批量更新，GORM 不会执行该操作，并返回ErrMissingWhereClause
	result := db.Model(&User{}).Debug().Update("Name", "test") // gorm.ErrMissingWhereClause:WHERE conditions required
	fmt.Println(result.RowsAffected, result.Error)

	// 可以使用 1 = 1 之类的条件来强制全局更新
	result1 := db.Model(&User{}).Where("1 = 1").Debug().Update("Name", "test")
	//  UPDATE "User" SET "Name"='test' WHERE 1 = 1
	fmt.Println(result1.RowsAffected) // 更新的记录数
	fmt.Println(result1.Error)        // 更新操作返回的错误
}

// ExprUpdate 通过Sql表达式更新
func ExprUpdate() {
	var user User
	db.First(&user, 1)
	db.Model(&user).Debug().Update("age", gorm.Expr("age * ? + ?", 2, 100))
	// UPDATE "User" SET "age"=age * 2 + 100 WHERE "Id" = 1

	db.Model(&user).Debug().Updates(map[string]interface{}{"age": gorm.Expr("age * ? + ?", 3, 100)})
	// UPDATE "User" SET "age"=age * 3 + 100 WHERE "Id" = 1

	db.Model(&user).Debug().UpdateColumn("age", gorm.Expr("age - ?", 10))
	// UPDATE "User" SET "age"=age - 10 WHERE "Id" = 1

	db.Model(&user).Where("age > ?", 18).Debug().UpdateColumn("age", gorm.Expr("age - ?", 1))
	// UPDATE "User" SET "age"=age - 1 WHERE age > 18 AND "Id" = 1
}
