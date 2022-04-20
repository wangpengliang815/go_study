// 参考：https://www.bookstack.cn/read/gorm-2.0/docs-delete.md
package main

import (
	. "goProject/gorm2.0/common"
)

var db = CreateDbConn()

func main() {
	DefaultDel()
}

// DefaultDel 默认删除
func DefaultDel() {
	var user User
	db.First(&user, 3)

	// 删除一条已有的记录,Id=3
	db.Debug().Delete(&user)
	// DELETE FROM "User" WHERE "User"."Id" = 3

	// 通过内联条件删除记录
	db.Debug().Delete(&User{}, 4)
	// DELETE FROM "User" WHERE "User"."Id" = 4

	// 带上其它条件
	db.Where("name = ?", "test").Debug().Delete(&User{})
	// DELETE FROM "User" WHERE name = 'test'
}
