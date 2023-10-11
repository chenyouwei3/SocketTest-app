package model

type User struct {
	Id         int64   `db:"id" json:"id" gorm:"column:id;type:bigint;primary_key;not null"`
	Name       string  `db:"name" json:"name" gorm:"column:name;type:varchar(20);not null"`
	Account    string  `db:"account" json:"account" gorm:"column:account;type:varchar(20);not null"`
	Password   string  `db:"password" json:"password" gorm:"column:password;type:varchar(60);not null"`
	IsValid    bool    `db:"isValid" json:"isValid" gorm:"column:isValid;type:tinyint(1);not null;default:1"`
	Salt       string  `db:"salt" json:"salt" gorm:"column:salt;type:varchar(60);not null"`
	Auth       []int64 `db:"auth" json:"auth" gorm:"column:auth;"`
	CreateTime string  `db:"createTime" json:"createTime" gorm:"column:createTime;type:varchar(20);not null"`
	UpdateTime string  `db:"updateTime" json:"updateTime" gorm:"column:updateTime;type:varchar(20)"`
}
