package models

import (
	"time"
)

type Tb_weigh_batch struct {
	Id                   int64
	BatchCode            string `xorm:"varchar(64)"`
	BatchNumber          string `xorm:"varchar(64)"`
	BatchName            string `xorm:"varchar(64)"`
	WeighHouseCodes      string `xorm:"varchar(256)"`
	Status               bool
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
