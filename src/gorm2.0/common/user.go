/*
 gorm结构体定义：https://gorm.io/zh_CN/docs/models.html
 GORM倾向于约定而不是配置。
1.默认情况下使用 ID 作为主键
2.使用结构体名的 蛇形复数 作为表名
3.字段名的蛇形作为列名，
4.使用 CreatedAt、UpdatedAt 字段追踪创建、更新时间
*/

package common

import "time"

/*
	1."primaryKey"`: 表示设置某列为主键
	2."default:'test'"`: 表示为某字段设置默认值;插入记录到数据库时，默认值会被用于填充值为零值的字段
    3.column:Name:设置列名称
    4.type:nvarchar(64):指定列类型
    5.not null
	需要注意 对于声明了默认值的字段，像 0、''、false 等零值是不会保存到数据库。需要使用指针类型或 Scanner/Valuer 来避免这个问题
*/
type User struct {
	// gorm.Model        // gorm里自定义的结构包含ID/CreatedAt/UpdatedAt/DeletedAt四个字段,其中DeletedAt标记当前数据是否被删除(gorm中默认使用逻辑删除)如果使用的话对应的Sql语句会自动添加相关筛选语句
	Id          int       `gorm:"primaryKey; column:Id"`
	Name        string    `gorm:"default:'test'; column:Name; type:nvarchar(64);not null;comment:用户姓名"`
	Age         int       `gorm:"column:Age"`
	PhoneNumber string    `gorm:"column:PhoneNumber"`
	Address     string    `gorm:"column:Address"`
	CreateTime  time.Time `gorm:"column:CreateTime"`
}

// TableName 重命名表名User否则为小写的user+复数s也就是users
func (User) TableName() string {
	return "User"
}
