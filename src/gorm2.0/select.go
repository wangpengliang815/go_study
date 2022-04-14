package main

import "fmt"

// 查询单个对象(主键升序)
func Select_First() {
	var user User
	db.Debug().First(&user) // SELECT * FROM "User" ORDER BY "User"."id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
	fmt.Printf("user:%#v", user)
	// 或者
	// user := new(User)
	// db.First(user)
}

// 查询单个对象(使用自然排序)
func Select_Take() {
	var user User
	db.Debug().Take(&user) // SELECT * FROM "User" ORDER BY "id" OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
	fmt.Printf("user:%#v", user)
}

// 查询单个对象(主键倒序)
func Select_Last() {
	var user User
	db.Debug().Last(&user) // SELECT * FROM "User" ORDER BY "User"."id" DESC OFFSET 0 ROW FETCH NEXT 1 ROWS ONLY
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
