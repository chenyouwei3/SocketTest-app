package model

type Device struct {
	Id         int64  `json:"id" bson:"id" gorm:"column:id;type:bigint;primary_key;not null"`                 //Id
	Code       string `json:"code" bson:"code"gorm:"column:code;type:varchar(10);not null"`                   //编码
	Addr       string `json:"addr" bson:"addr"gorm:"column:addr;type:varchar(20);not null"`                   //addr
	Protocol   string `json:"protocol" bson:"protocol"gorm:"column:protocol;type:varchar(20);not null"`       //协议
	IP         string `json:"ip" bson:"ip"gorm:"column:ip;type:varchar(20);not null"`                         //ip地址
	Port       string `json:"port" bson:"port"gorm:"column:port;type:varchar(5);not null"`                    //端口
	IsListen   bool   `json:"isListen" bson:"isListen"gorm:"column:isListen;type:varchar(1);not null"`        //是否监听
	Data       string `json:"data" bson:"data"gorm:"column:data;type:bigint;"`                                //原始数据,16进制
	CreateTime string `json:"createTime" bson:"createTime"gorm:"column:createTime;type:varchar(20);not null"` //创建时间
	UpdateTime string `json:"updateTime" bson:"updateTime"gorm:"column:createTime;type:varchar(20);"`         //更新时间
}
