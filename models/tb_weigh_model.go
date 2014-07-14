package models

import (
	"time"
)

type Tb_weigh_model struct {
	Id                   int64
	ModelCode            string `xorm:"varchar(64)"`
	ModelNumber          string `xorm:"varchar(64)"`
	ModelName            string `xorm:"varchar(128)"`
	Length               float64
	Width                float64
	Height               float64
	LengthUnit           string `xorm:"varchar(64)"`
	WidthUnit            string `xorm:"varchar(64)"`
	HeightUnit           string `xorm:"varchar(64)"`
	WeighHouseCodes      string `xorm:"varchar(256)"`
	Status               int8
	OperateBit           int8
	UploadBit            int8
	UploadTime           time.Time `xorm:"DateTime"`
	CreateUserCode       string    `xorm:"varchar(64)"`
	CreateUserName       string    `xorm:"varchar(64)"`
	CreateTime           time.Time
	LastModifiedUserCode string `xorm:"varchar(64)"`
	LastModifiedUserName string `xorm:"varchar(64)"`
	LastModifiedTime     time.Time
	Remark               string
	Attribute1           string
	Attribute2           string
	Attribute3           string
	Attribute4           string
	Attribute5           string
	Attribute6           string
	Attribute7           string
	Attribute8           string
	Attribute9           string
	Attribute10          string
	Attribute11          string
	Attribute12          string
	Attribute13          string
	Attribute14          string
	Attribute15          string
	Attribute16          string
}
