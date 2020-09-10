package model

type Account struct {
	// 物理id，自增，主键
	Id int `gorm:"column:id;primaryKey;autoIncrement;type:varchar(11)"`
	// 用户昵称,唯一键
	Nickname string `gorm:"column:nickname;unique;type:varchar(50)"`
	// 密码
	Password string `gorm:"column:password;type:varchar(50)"`
	// 手机号
	Phone string `gorm:"column:phone;type:varchar(11)"`

	Model
}
