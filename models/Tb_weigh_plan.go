package models

import (
	"time"
)

type Tb_weigh_plan struct {
	Id                   int64  `xorm:"autoincr notnull"`
	PlanCode             string `xorm:"pk varchar(64)"`
	PlanNumber           string `xorm:"varchar(64)"`
	OrganizationCode     string `xorm:"varchar(64)"`
	WeighHouseCodes      string `xorm:"varchar(256)"`
	SupplierCode         string `xorm:"varchar(64)"`
	SupplierName         string `xorm:"varchar(128)"`
	RecipientCode        string `xorm:"varchar(64)"`
	RecipientName        string `xorm:"varchar(128)"`
	MaterialCode         string `xorm:"varchar(64)"`
	MaterialName         string `xorm:"varchar(128)"`
	WeighType            string `xorm:"varchar(64)"`
	MeasureUnit          string `xorm:"varchar(64)"`
	PlanBeginTime        time.Time
	PlanEndTime          time.Time
	PlanExecuteTime      time.Time
	PlanAmount           float64
	PlanExecuteAmount    float64
	IsTotalControl       bool
	TotalFloatRange      float64
	TotalFloatPerRange   float64
	IsCarControl         bool
	IsFromErp            bool
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
