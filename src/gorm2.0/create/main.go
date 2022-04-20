// @title gorm-创建相关操作
// @description
// @author wangpengliang
// @date 2022-04-14 23:46:46
package main

import (
	"fmt"
	. "goProject/gorm2.0/common"
	"time"
)

var db = CreateDbConn()

func main() {
	CreateInBatches()
}

// DefaultCreate 默认创建方式
func DefaultCreate() {
	user := User{Name: "wangpengliang", Age: 18, PhoneNumber: "18888888888", Address: "beijing", CreateTime: time.Now()}
	result := db.Debug().Create(&user)
	// INSERT INTO "User" ("Name","Age","PhoneNumber","Address","CreateTime") OUTPUT INSERTED."Id" VALUES ('wangpengliang',18,'18888888888','beijing','2022-04-19 17:14:42.329');
	if result.Error != nil {
		fmt.Println(result.Error) // 返回error
	}
	fmt.Println(result.RowsAffected, user.Id) // 返回插入记录的条数 | 插入数据的主键
}

// SpecifiedField 使用指定字段创建
func SpecifiedField() {
	user := User{Name: "wangpengliang", Age: 18, PhoneNumber: "18888888888", Address: "beijing", CreateTime: time.Now()}
	// 这里只会插入Name,Age两个字段,虽然user是多个字段
	db.Select("Name", "Age").Debug().Create(&user)
	// INSERT INTO "User" ("Name","Age") OUTPUT INSERTED."Id" VALUES ('wangpengliang',18);
}

// ExcludeSpecifiedField 排除指定字段
func ExcludeSpecifiedField() {
	user := User{Name: "wangpengliang", Age: 18, PhoneNumber: "18888888888", Address: "beijing", CreateTime: time.Now()}
	// 这里会排除Age、PhoneNumber字段
	db.Omit("Age", "PhoneNumber").Debug().Create(&user)
	//  INSERT INTO "User" ("Name","Address","CreateTime") OUTPUT INSERTED."Id" VALUES ('wangpengliang','beijing','2022-04-19 17:19:52.02');
}

// BatchCreate 批量创建
func BatchCreate() {
	users := []User{
		{Name: "wangpengliang", Age: 18, PhoneNumber: "18888888888", Address: "beijing", CreateTime: time.Now()},
		{Name: "lizimeng", Age: 19, PhoneNumber: "166666666", Address: "shanghai", CreateTime: time.Now()},
	}
	db.Debug().Create(&users)
	// INSERT INTO "User" ("Name","Age","PhoneNumber","Address","CreateTime") OUTPUT INSERTED."Id" VALUES ('wangpengliang',18,'18888888888','beijing','2022-04-19 17:21:18.668'),('lizimeng',19,'166666666','shanghai','2022-04-19 17:21:18.668');
}

// CreateInBatches 分批批量创建
func CreateInBatches() {
	users := []User{
		{Name: "w1", Age: 18, PhoneNumber: "18888888888", Address: "beijing", CreateTime: time.Now()},
		{Name: "w2", Age: 19, PhoneNumber: "166666666", Address: "shanghai", CreateTime: time.Now()},
		{Name: "w3", Age: 20, PhoneNumber: "166666666", Address: "shanghai", CreateTime: time.Now()},
		{Name: "w4", Age: 21, PhoneNumber: "166666666", Address: "shanghai", CreateTime: time.Now()},
		{Name: "w5", Age: 22, PhoneNumber: "166666666", Address: "shanghai", CreateTime: time.Now()},
		{Name: "w6", Age: 23, PhoneNumber: "166666666", Address: "shanghai", CreateTime: time.Now()},
	}
	// 分批创建,每批数量为2
	db.Debug().CreateInBatches(users, 2)
	// INSERT INTO "User" ("Name","Age","PhoneNumber","Address","CreateTime") OUTPUT INSERTED."Id" VALUES ('w1',18,'18888888888','beijing','2022-04-20 14:09:02.946'),('w2',19,'166666666','shanghai','2022-04-20 14:09:02.946');
	// INSERT INTO "User" ("Name","Age","PhoneNumber","Address","CreateTime") OUTPUT INSERTED."Id" VALUES ('w3',20,'166666666','shanghai','2022-04-20 14:09:02.946'),('w4',21,'166666666','shanghai','2022-04-20 14:09:02.946');
	// INSERT INTO "User" ("Name","Age","PhoneNumber","Address","CreateTime") OUTPUT INSERTED."Id" VALUES ('w5',22,'166666666','shanghai','2022-04-20 14:09:02.946'),('w6',23,'166666666','shanghai','2022-04-20 14:09:02.946');
}

// CreateByMap 根据Map创建
func CreateByMap() {
	db.Model(&User{}).Debug().Create(map[string]interface{}{
		"Name":        "wangpengliang",
		"Age":         18,
		"PhoneNumber": "18888888888",
		"Address":     "beijing",
		"CreateTime":  time.Now(),
	})
	//  INSERT INTO "User" ("Address","Age","CreateTime","Name","PhoneNumber") OUTPUT INSERTED."Id" VALUES ('beijing',18,'2022-04-19 17:24:40.844','wangpengliang','18888888888');

	// batch 注意: 根据 map 创建时，association 不会被调用，且主键也不会自动填充,所以需要手动给主键赋值,但是由于主键列一般有primaryKey标识会引发异常:当 IDENTITY_INSERT 设置为 OFF 时，不能为表 'xxx' 中的标识列插入显式值。
	type Person struct {
		PersonId int
		Name     string
		Age      int
		Address  string
	}
	//db.Model(&Person{}).Debug().Create([]map[string]interface{}{
	//	{"PersonId": 999, "Name": "wangpengliang", "Age": 18, "Address": "beijing"},
	//	{"PersonId": 888, "Name": "lizimeng", "Age": 19, "Address": "shanghai"},
	//})
}
