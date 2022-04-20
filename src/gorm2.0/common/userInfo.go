package common

type UserInfo struct {
	Id          int `gorm:"primaryKey; column:Id"`
	UserId      int `gorm:"column:UserId;not null"`
	Description int `gorm:"column:Description"`
}

func (UserInfo) TableName() string {
	return "UserInfo"
}
