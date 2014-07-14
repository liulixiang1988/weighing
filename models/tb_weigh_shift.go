package models

import (
	"time"
)

type Tb_weigh_shift struct {
	Id                   int64
	ShiftCode            string `xorm:"varchar(64)"`
	ShiftName            string `xorm:"varchar(64)"`
	StartTime            time.Time
	EndTime              time.Time
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
}
