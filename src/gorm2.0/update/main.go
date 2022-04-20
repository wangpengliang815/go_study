package main

import (
	. "goProject/gorm2.0/common"
)

func main() {
	UpdateFieldSelect()
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
}
