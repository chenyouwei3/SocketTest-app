package model

type Api struct {
	Id     int64  `json:"id" bson:"id" gorm:"column:id;type:bigint;primary_key;not null"`
	Name   string `json:"name" bson:"name" gorm:"column:name;type:varchar(20);not null"`
	Path   string `json:"path" bson:"path" gorm:"column:path;type:varchar(20);not null"`
	Method string `json:"method" bson:"method" gorm:"column:,method;type:varchar(5);not null"`
}
