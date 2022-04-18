// @title gorm-查询相关操作
// @description
// @author wangpengliang
// @date 2022-04-14 23:47:10
package main

import (
	"fmt"
	. "goProject/gorm2.0/common"

	"gorm.io/gorm/clause"
)

var user User
var users User
var db = CreateDbConn()

func main() {
	Select_GroupAndHaving()
}

// Select_First 查询单个对象(主键升序)
func Select_First() {
	var user User
	db.Debug().First(&user)
	// SELECT * FROM "User" WHERE "User"."deleted_at" IS NULL ORDER BY "User"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY

	fmt.Printf("user:%#v", user)
	// 或者
	// user := new(User)
	// db.First(user)
}

// Select_Take 查询单个对象(使用自然排序)
func Select_Take() {
	var user User
	db.Debug().Take(&user)
	//  SELECT * FROM "User" WHERE "User"."deleted_at" IS NULL ORDER BY "id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
}

// Select_Last 查询单个对象(主键倒序)
func Select_Last() {
	var user User
	db.Debug().Last(&user)
	// SELECT * FROM "User" WHERE "User"."deleted_at" IS NULL ORDER BY "User"."id" DESC OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
	fmt.Printf("user:%#v", user)
}

// 主键查询,如果主键是数字类型，可以使用内联条件来检索对象
func Select_First_ByPrimaryKey() {
	var user User
	db.Debug().First(&user, 1) // SELECT * FROM "User" WHERE "User"."id" = 1 ORDER BY "User"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
	fmt.Printf("user:%#v", user)
}

// 查询全部或者in操作
func Select_Find() {
	var users []User
	db.Debug().Find(&users, []int{1, 2}) // SELECT * FROM "User" WHERE "User"."id" IN (1,2)
	db.Debug().Find(&users)              // SELECT * FROM "User"
	fmt.Printf("users:%#v", users)
}

// 查询条件
func Select_Where() {
	var user User
	var users []User
	// 获取第一条匹配的记录
	db.Where("name = ?", "wangpengliang").Debug().First(&user) // SELECT * FROM "User" WHERE name = 'wangpengliang' ORDER BY "User"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY

	// 获取全部匹配的记录
	db.Where("name <> ?", "wangpengliang").Debug().Find(&users) // SELECT * FROM "User" WHERE name <> 'wangpengliang'

	// IN
	db.Where("name IN ?", []string{"wangpengliang", "lizimeng"}).Debug().Find(&users) // SELECT * FROM "User" WHERE name IN ('wangpengliang','lizimeng')

	// LIKE
	db.Where("name LIKE ?", "%wangpengliang%").Debug().Find(&users) // SELECT * FROM "User" WHERE name LIKE '%wangpengliang%'

	// AND
	db.Where("name = ? AND age >= ?", "wangpengliang", "10").Debug().Find(&users) // SELECT * FROM "User" WHERE name = 'wangpengliang' AND age >= '10'

	// Time
	db.Where("updated_at > ?", "2022-01-01").Debug().Find(&users) // SELECT * FROM "User" WHERE updated_at > '2022-01-01' AND "User"."deleted_at" IS NULL

	// BETWEEN
	db.Where("created_at BETWEEN ? AND ?", "2022-01-01", "2022-09-01").Debug().Find(&users) // SELECT * FROM "User" WHERE (created_at BETWEEN '2022-01-01' AND '2022-09-01') AND "User"."deleted_at" IS NULL
}

// 指定Struct或者Map查询
// 注意:当使用结构作为条件查询时，GORM 只会查询非零值字段。这意味着如果字段值为 0、''、false 或其他 零值，该字段不会被用于构建查询条件
//      如果想要包含零值查询条件，可以使用 map，会包含所有 key-value 的查询条件
func Select_Struct() {
	var user User

	var users []User

	// Struct
	db.Where(&User{Name: "wangpengliang", Age: 18}).Debug().First(&user)
	// SELECT * FROM "User" WHERE "User"."name" = 'wangpengliang' AND "User"."age" = 18 AND "User"."deleted_at" IS NULL ORDER BY "User"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY

	// Map
	db.Where(map[string]interface{}{"name": "wangpengliang", "age": 18}).Debug().Find(&users)
	// SELECT * FROM "User" WHERE "age" = 18 AND "name" = 'wangpengliang' AND "User"."deleted_at" IS NULL

	// 主键切片条件
	db.Where([]int64{20, 21, 22}).Debug().Find(&users)
	// SELECT * FROM "User" WHERE "User"."id" IN (20,21,22) AND "User"."deleted_at" IS NULL
}

// 内联查询,将查询条件内联到 First 和 Find 之类的方法中，用法类似于 Where
func Select_Inline() {
	var user User

	var users []User

	// 根据主键获取记录，如果是非整型主键
	db.Debug().First(&user, "id = ?", "string_primary_key")
	// SELECT * FROM "User" WHERE id = 'string_primary_key' AND "User"."deleted_at" IS NULL ORDER BY "User"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONL

	// Plain SQL
	db.Debug().Find(&user, "name = ?", "wangepngliang")
	// SELECT * FROM "User" WHERE name = 'wangepngliang' AND "User"."deleted_at" IS NULL

	db.Debug().Find(&users, "name <> ? AND age > ?", "jinzhu", 20)
	// SELECT * FROM "User" WHERE (name <> 'jinzhu' AND age > 20) AND "User"."deleted_at" IS NULL

	// Struct
	db.Debug().Find(&users, User{Age: 20})
	// SELECT * FROM "User" WHERE "User"."age" = 20 AND "User"."deleted_at" IS NULL

	// Map
	db.Debug().Find(&users, map[string]interface{}{"age": 20})
	// SELECT * FROM "User" WHERE "age" = 20 AND "User"."deleted_at" IS NULL
}

// 构建 NOT 条件，用法和Where类似不过是取反值
func Select_Not() {
	var user User
	var users []User
	db.Not("name = ?", "wangpengliang").Debug().First(&user)
	// SELECT * FROM "User" WHERE NOT name = 'wangpengliang' AND "User"."deleted_at" IS NULL ORDER BY "User"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY

	// Not In
	db.Not(map[string]interface{}{"name": []string{"wangpengliang", "wangpengliang 2"}}).Debug().Find(&users)
	// SELECT * FROM "User" WHERE "name" NOT IN ('wangpengliang','wangpengliang 2') AND "User"."deleted_at" IS NULL

	// Struct
	db.Not(User{Name: "wangpengliang", Age: 18}).Debug().First(&user)
	// SELECT * FROM "User" WHERE ("User"."name" <> 'wangpengliang' AND "User"."age" <> 18) AND "User"."deleted_at" IS NULL AND "User"."id" = 2 ORDER BY "User"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY

	// 不在主键切片中的记录
	db.Not([]int64{1, 2, 3}).Debug().First(&user)
	// SELECT * FROM "User" WHERE "User"."id" NOT IN (1,2,3) AND "User"."deleted_at" IS NULL AND "User"."id" = 2 ORDER BY "User"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
}

// 构建 OR 条件
func Select_Or() {
	var users []User
	db.Where("name = ?", "wangpengliang").Or("age = ?", 18).Debug().Find(&users)
	// SELECT * FROM "User" WHERE (name = 'wangpengliang' OR age = 18) AND "User"."deleted_at" IS NULL

	// Struct
	db.Where("name = 'wangpengliang'").Or(User{Age: 18}).Debug().Find(&users)
	// SELECT * FROM "User" WHERE (name = 'wangpengliang' OR "User"."age" = 18) AND "User"."deleted_at" IS NULL

	// Map
	db.Where("name = 'wangpengliang'").Or(map[string]interface{}{"age": 18}).Debug().Find(&users)
	// SELECT * FROM "User" WHERE (name = 'wangpengliang' OR "age" = 18) AND "User"."deleted_at" IS NULL
}

// 查询指定字段
func Select_FixedField() {
	db.Select("name", "age").Debug().Find(&users)
	// SELECT "name","age" FROM "User" WHERE "User"."deleted_at" IS NULL

	db.Select([]string{"name", "age"}).Debug().Find(&users)
	// SELECT "name","age" FROM "User" WHERE "User"."deleted_at" IS NULL

	db.Table("user").Select("COALESCE(age,?)", 42).Debug().Rows()
	// SELECT COALESCE(age,42) FROM "user"
}

// Order排序
func Select_Order() {
	db.Order("age desc, name").Debug().Find(&users)
	// SELECT * FROM "User" WHERE "User"."deleted_at" IS NULL ORDER BY age desc, name

	// 多个 order
	db.Order("age desc").Order("name").Debug().Find(&users)
	// SELECT * FROM "User" WHERE "User"."deleted_at" IS NULL ORDER BY age desc,name

	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Debug().Find(&User{})
	// SELECT * FROM "User" WHERE "User"."deleted_at" IS NULL ORDER BY FIELD(id,1,2,3)
	// 注意：mssql: 'FIELD' 不是可以识别的 内置函数名称
}

// SQLServer中的TOP
func Select_LimitAndOffset() {
	var users1 []User
	var users2 []User

	db.Limit(1).Debug().Find(&users)
	// SELECT * FROM "User" WHERE "User"."deleted_at" IS NULL ORDER BY "id" OFFSET 0 ROW FETCH NEXT 3 ROWS ONLY

	// 通过 -1 消除 Limit 条件
	db.Limit(1).Find(&users1).Limit(-1).Debug().Find(&users2)
	// SELECT * FROM users LIMIT 10; (users1)
	// SELECT * FROM users; (users2)

	// offset表示跳过几行
	db.Offset(1).Debug().Find(&users)
	// SELECT * FROM "User" WHERE "User"."deleted_at" IS NULL ORDER BY "id" OFFSET 1 ROWS

	// 跳过1行获取第一条
	db.Limit(1).Offset(1).Find(&users)
	// SELECT * FROM "User" WHERE "User"."deleted_at" IS NULL ORDER BY "id" OFFSET 1 ROWS

	// 通过 -1 消除 Offset 条件,跟上面同理
	db.Offset(1).Find(&users1).Offset(-1).Find(&users2)
	// SELECT * FROM users OFFSET 10; (users1)
	// SELECT * FROM users; (users2)
}

type Result struct {
	name  string
	total int
}

// Group And Having
func Select_GroupAndHaving() {
	var result []Result
	db.Model(&User{}).Select("name, sum(age) as total").Group("name").Debug().Find(&result)
	// SELECT name, sum(age) as total FROM "User" WHERE "User"."deleted_at" IS NULL GROUP BY "name"
	fmt.Printf("result:%#v", result)

	db.Model(&User{}).Select("name, sum(age) as total").Group("name").Having("name = ?", "wangpengliang").Debug().Find(&result)
	// SELECT name, sum(age) as total FROM "User" WHERE "User"."deleted_at" IS NULL GROUP BY "name" HAVING name = 'wangpengliang'

	rows, _ := db.Table("user").Select("name, sum(age) as total").Group("name").Debug().Rows()
	// SELECT name, sum(age) as total FROM "user" GROUP BY "name"
	defer rows.Close()

	for rows.Next() {
		//
	}

	rows2, _ := db.Table("user").Select("name, sum(age) as total").Group("name").Having("sum(age) > ?", 10).Debug().Rows()
	// SELECT name, sum(age) as total FROM "user" GROUP BY "name" HAVING sum(age) > 10
	defer rows2.Close()
	for rows2.Next() {
		//
	}

	db.Table("user").Select("name, sum(age) as total").Group("name").Having("sum(age) > ?", 30).Debug().Scan(&result)
	// SELECT name, sum(age) as total FROM "user" GROUP BY "name" HAVING sum(age) > 30
}
