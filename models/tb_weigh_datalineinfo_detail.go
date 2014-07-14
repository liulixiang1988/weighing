package models

import (
	"time"
)

//明细表
type Tb_weigh_datalineinfo_detail struct {
	Id          int64
	BatchNumber int64  `xorm:"varchar(64)"`
	CarNumber   string `xorm:"varchar(64)"`
	Weight      float64
	WeighTime   time.Time
	OperateBit  int8
	Remark      string
	Attribute1  string
	Attribute2  string
	Attribute3  string
	Attribute4  string
	Attribute5  string
	Attribute6  string
	Attribute7  string
	Attribute8  string
}
