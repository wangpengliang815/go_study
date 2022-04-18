// @title gorm-创建相关操作
// @description
// @author wangpengliang
// @date 2022-04-14 23:46:46
package main

import (
	"fmt"
	. "goProject/gorm2.0/common"
)

var db = CreateDbConn()

func main() {
	Insert_Batch()
}

// Insert 普通创建
func Insert() {
	user := User{Name: "wangpengliang", Age: 18, Address: "beijing"}

	result := db.Debug().Create(&user)
	//  INSERT INTO "User" ("created_at","updated_at","deleted_at","name","age","address") OUTPUT INSERTED."id" VALUES ('2022-04-15 23:12:13.627','2022-04-15 23:12:13.627',NULL,'wangpengliang',18,'beijing');
	if result.Error != nil {
		fmt.Println(result.Error)
	}
	fmt.Println(result.RowsAffected)
	// user.ID 插入数据的主键
	// result.Error        // 返回error
	// result.RowsAffected // 返回插入记录的条数
}

// Insert_SelectField 使用选定字段创建
func Insert_SelectField() {
	user := User{Name: "wangpengliang", Age: 18, Address: "beijing"}
	// 这里只会插入Name,Age两个字段,虽然传入的user是3个字段
	db.Select("Name", "Age").Debug().Create(&user)
	// INSERT INTO "User" ("created_at","updated_at","name","age") OUTPUT INSERTED."id" VALUES ('2022-04-15 23:15:32.688','2022-04-15 23:15:32.688','wangpengliang',18);
}

// Insert_ExcludeField 排除指定字段创建
func Insert_ExcludeField() {
	user := User{Name: "wangpengliang", Age: 18, Address: "beijing"}
	// 这里会排除Age字段
	db.Omit("Age").Debug().Create(&user)
	// INSERT INTO "User" ("created_at","updated_at","deleted_at","name","address") OUTPUT INSERTED."id" VALUES ('2022-04-15 23:17:28.892','2022-04-15 23:17:28.892',NULL,'wangpengliang','beijing');
}

// Insert_Batch 批量创建
func Insert_Batch() {
	users := []User{
		{Name: "wangpengliang", Age: 18, Address: "beijing"},
		{Name: "lizimeng", Age: 19, Address: "shanghai"},
	}
	db.Debug().Create(&users)
	// INSERT INTO "User" ("created_at","updated_at","deleted_at","name","age","address") OUTPUT INSERTED."id" VALUES ('2022-04-15 23:18:39.353','2022-04-15 23:18:39.353',NULL,'wangpengliang',18,'beijing'),('2022-04-15 23:18:39.353','2022-04-15 23:18:39.353',NULL,'lizimeng',19,'shanghai');
}

// Insert_ByMap 根据Map创建
func Insert_ByMap() {
	db.Model(&User{}).Debug().Create(map[string]interface{}{
		"Name": "wangpengliang", "Age": 18, "Address": "beijing",
	})
	//  INSERT INTO "User" ("address","age","name") OUTPUT INSERTED."id" VALUES ('beijing',18,'wangpengliang');

	// batch 注意: 根据 map 创建时，association 不会被调用，且主键也不会自动填充,所以需要手动给主键赋值,但是由于主键列一般有primaryKey标识会引发异常:当 IDENTITY_INSERT 设置为 OFF 时，不能为表 'xxx' 中的标识列插入显式值。
	db.Model(&Person{}).Debug().Create([]map[string]interface{}{
		{"PersonId": 999, "Name": "wangpengliang", "Age": 18, "Address": "beijing"},
		{"PersonId": 888, "Name": "lizimeng", "Age": 19, "Address": "shanghai"},
	})
	// INSERT INTO "Person" ("address","age","name","person_id") VALUES ('beijing',18,'wangpengliang',999),('shanghai',19,'lizimeng',888);
}
