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
	Scan()
}

// First 查询单个对象(主键升序)
func First() {
	db.Debug().First(&user)
	// SELECT * FROM "User" ORDER BY "User"."Id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
	fmt.Printf("user:%#v", user)
	// 或者
	user := new(User)
	db.Debug().First(user)
	fmt.Printf("user:%#v", user)
}

// Take 查询单个对象(使用自然排序)
func Take() {
	db.Debug().Take(&user)
	// SELECT * FROM "User" ORDER BY "Id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
}

// Last 查询单个对象(主键倒序)
func Last() {
	db.Debug().Last(&user)
	// SELECT * FROM "User" ORDER BY "User"."Id" DESC OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
}

// SelectByIntPrimaryKey 主键查询,主键为int类型
func SelectByIntPrimaryKey() {
	db.Debug().First(&user, 1)
	// SELECT * FROM "User" WHERE "User"."Id" = 1 ORDER BY "User"."Id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
}

// SelectByStringPrimaryKey 主键查询,主键为string类型可以这么写
func SelectByStringPrimaryKey() {
	db.Debug().First(&user, "id = ?", "1b74413f-f3b8-409f-ac47-e8c062e3472a")
	// SELECT * FROM "User" WHERE "User"."Id" = "1b74413f-f3b8-409f-ac47-e8c062e3472a"
}

// Find 查询全部
func Find() {
	db.Debug().Find(&users)
	// SELECT * FROM "User"
}

// In 数据库In操作
func In() {
	db.Debug().Find(&users, []int{1, 2})
	// SELECT * FROM "User" WHERE "User"."Id" IN (1,2)
}

// Where 查询条件
func Where() {
	// 获取第一条匹配的记录
	db.Where("name = ?", "wangpengliang").Debug().First(new(User))
	// SELECT * FROM "User" WHERE name = 'wangpengliang' ORDER BY "User"."Id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY

	// 获取全部匹配的记录
	db.Where("name <> ?", "wangpengliang").Debug().Find(new(User))
	// SELECT * FROM "User" WHERE name <> 'wangpengliang'

	// IN
	db.Where("name IN ?", []string{"wangpengliang", "lizimeng"}).Debug().Find(new(User))
	// SELECT * FROM "User" WHERE name IN ('wangpengliang','lizimeng')

	// LIKE
	db.Where("name LIKE ?", "%wangpengliang%").Debug().Find(new(User))
	// SELECT * FROM "User" WHERE name LIKE '%wangpengliang%'

	// AND
	db.Where("name = ? AND age >= ?", "wangpengliang", "10").Debug().Find(new(User))
	// SELECT * FROM "User" WHERE name = 'wangpengliang' AND age >= '10'

	// Time
	db.Where("CreateTime > ?", "2022-01-01").Debug().Find(new(User))
	// SELECT * FROM "User" WHERE CreateTime > '2022-01-01'

	// BETWEEN
	db.Where("CreateTime BETWEEN ? AND ?", "2022-01-01", "2022-09-01").Debug().Find(new(User))
	// SELECT * FROM "User" WHERE CreateTime BETWEEN '2022-01-01' AND '2022-09-01'
}

// SelectByStructOrMap 指定Struct或者Map查询
// 注意:当使用结构作为条件查询时，GORM 只会查询非零值字段。这意味着如果字段值为 0、''、false 或其他 零值，该字段不会被用于构建查询条件
// 如果想要包含零值查询条件，可以使用 map，会包含所有 key-value 的查询条件
func SelectByStructOrMap() {
	// Struct
	db.Where(&User{Name: "wangpengliang", Age: 18}).Debug().First(new(User))
	// SELECT * FROM "User" WHERE "User"."Name" = 'wangpengliang' AND "User"."Age" = 18 ORDER BY "User"."Id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY

	// Map
	db.Where(map[string]interface{}{"name": "wangpengliang", "age": 18}).Debug().Find(new([]User))
	// SELECT * FROM "User" WHERE "age" = 18 AND "name" = 'wangpengliang'

	// 主键切片条件
	db.Where([]int64{20, 21, 22}).Debug().Find(new([]User))
	// SELECT * FROM "User" WHERE "User"."Id" IN (20,21,22)
}

// InlineSelect 内联查询,将查询条件内联到 First 和 Find 之类的方法中，用法类似于 Where
func InlineSelect() {
	// 根据主键获取记录，如果是非整型主键
	db.Debug().First(new(User), "id = ?", "string_primary_key")

	// Plain SQL
	db.Debug().Find(new(User), "name = ?", "wangpengliang")

	db.Debug().Find(new(User), "name <> ? AND age > ?", "wangpengliang", 20)

	// Struct
	db.Debug().Find(new(User), User{Age: 20})

	// Map
	db.Debug().Find(new(User), map[string]interface{}{"age": 20})
}

// Not 构建 NOT 条件，用法和Where类似不过是取反值
func Not() {
	db.Not("name = ?", "wangpengliang").Debug().First(new(User))
	// SELECT * FROM "User" WHERE NOT name = 'wangpengliang' ORDER BY "User"."Id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY

	// Not In
	db.Not(map[string]interface{}{"name": []string{"wangpengliang", "wangpengliang2"}}).Debug().Find(new([]User))
	// SELECT * FROM "User" WHERE "name" NOT IN ('wangpengliang','wangpengliang2')

	// Struct
	db.Not(User{Name: "wangpengliang", Age: 18}).Debug().First(new(User))
	// SELECT * FROM "User" WHERE ("User"."Name" <> 'wangpengliang' AND "User"."Age" <> 18) ORDER BY "User"."Id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY

	// 不在主键切片中的记录
	db.Not([]int64{1, 2, 3}).Debug().First(new(User))
	// SELECT * FROM "User" WHERE "User"."Id" NOT IN (1,2,3) ORDER BY "User"."Id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
}

// Or 构建 OR 条件
func Or() {
	db.Where("name = ?", "wangpengliang").Or("age = ?", 18).Debug().Find(new([]User))
	// SELECT * FROM "User" WHERE name = 'wangpengliang' OR age = 18

	// Struct
	db.Where("name = 'wangpengliang'").Or(User{Age: 18}).Debug().Find(new(User))
	// SELECT * FROM "User" WHERE name = 'wangpengliang' OR "User"."Age" = 18

	// Map
	db.Where("name = 'wangpengliang'").Or(map[string]interface{}{"age": 18}).Debug().Find(new([]User))
	// SELECT * FROM "User" WHERE name = 'wangpengliang' OR "age" = 18
}

// FixedFieldSelect 查询指定字段
func FixedFieldSelect() {
	db.Select("name", "age").Debug().Find(new([]User))
	// SELECT name,age FROM "User"

	db.Select([]string{"name", "age"}).Debug().Find(new([]User))
	// SELECT name,age FROM "User"

	// COALESCE函数当age=null时返回18否则返回真实值
	db.Table("User").Select("COALESCE(age,?)", 18).Debug().Rows()
	//  SELECT COALESCE(age,18) FROM "User"
}

// Order Order排序
func Order() {
	db.Order("age desc, name").Debug().Find(new([]User))
	// SELECT * FROM "User" ORDER BY age desc, name

	// 多个 order
	db.Order("age desc").Order("name").Debug().Find(new([]User))
	// SELECT * FROM "User" ORDER BY age desc,name

	// 注意：mssql: 'FIELD' 不是可以识别的内置函数名称
	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Debug().Find(new([]User))
	// SELECT * FROM "User" ORDER BY FIELD(id,1,2,3)
}

// LimitAndOffset SQLServer中的TOP
func LimitAndOffset() {
	db.Limit(1).Debug().Find(new([]User))
	// SELECT * FROM "User" ORDER BY "Id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY

	// 通过 -1 消除 Limit 条件
	db.Limit(1).Find(new([]User)).Limit(-1).Debug().Find(new([]User))
	// SELECT * FROM User LIMIT 10; (users1)
	// SELECT * FROM User; (users2)

	// offset表示跳过几行
	db.Offset(1).Debug().Find(new([]User))
	// SELECT * FROM "User" ORDER BY "Id" OFFSET 1 ROWS

	// 跳过1行获取第一条
	db.Limit(1).Offset(1).Debug().Find(new([]User))
	// SELECT * FROM "User" ORDER BY "Id" OFFSET 1 ROWS FETCH NEXT 1 ROWS ONLY

	// 通过 -1 消除 Offset 条件,跟上面同理
	db.Offset(1).Find(new([]User)).Offset(-1).Debug().Find(new([]User))
	// SELECT * FROM User OFFSET 10; (users1)
	// SELECT * FROM User; (users2)
}

// GroupAndHaving Group分组及Having
func GroupAndHaving() {
	type Result struct {
		Name  string
		Total int
	}
	var result []Result
	db.Model(new(User)).Select("Name, sum(age) as total").Group("name").Debug().Find(&result)
	// SELECT Name, sum(age) as total FROM "User" GROUP BY "name"
	fmt.Println(result)

	db.Model(new(User)).Select("name, sum(age) as total").Group("name").Having("name = ?", "wangpengliang").Debug().Find(&result)
	// SELECT name, sum(age) as total FROM "User" GROUP BY "name" HAVING name = 'wangpengliang'

	rows, _ := db.Table("User").Select("name, sum(age) as total").Group("name").Debug().Rows()
	// SELECT name, sum(age) as total FROM "user" GROUP BY "name"
	defer rows.Close()
	for rows.Next() {
		//
	}

	rows2, _ := db.Table("User").Select("name, sum(age) as total").Group("name").Having("sum(age) > ?", 10).Debug().Rows()
	// SELECT name, sum(age) as total FROM "User" GROUP BY "name" HAVING sum(age) > 10
	defer rows2.Close()
	for rows2.Next() {
		//
	}

	db.Table("User").Select("name, sum(age) as total").Group("name").Having("sum(age) > ?", 30).Debug().Scan(&result)
	// SELECT name, sum(age) as total FROM "user" GROUP BY "name" HAVING sum(age) > 30
	fmt.Println(result)
}

// Distinct 去重复
func Distinct() {
	db.Distinct("name,age").Order("name,age desc").Debug().Find(&user)
	// SELECT DISTINCT name,age FROM "User" ORDER BY name,age desc
}

// Join 联表操作
func Join() {
	type Result struct {
		Name        string
		Age         int
		Description string
	}
	var result []Result
	db.Model(new(User)).Select("[User].name,[User].age,UserInfo.Description").Joins("left join UserInfo on UserInfo.UserId = [User].Id").Debug().Scan(&result)
	// SELECT [User].name,[User].age,UserInfo.Description FROM "User" left join UserInfo on UserInfo.UserId = [User].Id
	for _, value := range result {
		fmt.Println(value.Name, value.Age, value.Description)
	}

	db.Table("User").Select("[User].name, UserInfo.Description").Joins("left join UserInfo on UserInfo.UserId = [User].Id").Debug().Rows()
	// SELECT [User].name, UserInfo.Description FROM "User" left join UserInfo on UserInfo.UserId = [User].Id

	var result2 []Result
	db.Table("User").Select("[User].name, UserInfo.Description").Joins("left join UserInfo on UserInfo.UserId = [User].Id").Debug().Scan(&result2)
	for _, value := range result2 {
		fmt.Println(value.Name, value.Age, value.Description)
	}

	// 带参数的多表连接
	db.Joins("JOIN UserInfo ON UserInfo.UserId = [User].Id AND [User].Name = ?", "wangpengliang").Debug().Find(new([]User))
	// SELECT "User"."Id","User"."Name","User"."Age","User"."PhoneNumber","User"."Address","User"."CreateTime" FROM "User" JOIN UserInfo ON UserInfo.UserId = [User].Id AND [User].Name = 'wangpengliang'

	// TODO:Joins预加载
}

// Scan 将结果到至 struct，用法与 Find 类似
func Scan() {
	type Result struct {
		Name string
		Age  int
	}

	var result Result
	db.Table("User").Select("name", "age").Where("name = ?", "wangpengliang").Scan(&result)
	fmt.Println(result)

	// 原生 SQL
	db.Raw("SELECT name, age FROM [User] WHERE name = ?", "wangpengliang").Scan(&result)
	fmt.Println(result)
}
