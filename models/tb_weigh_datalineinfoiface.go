package models

import (
	"time"
)

type Tb_weigh_datalineinfoiface struct {
	Id                       int64
	WeighRecordNumber        string `xorm:"varchar(64)"`
	PlanCode                 string `xorm:"varchar(64)"`
	PlanNumber               string `xorm:"varchar(64)"`
	BatchNumber              string `xorm:"varchar(64)"`
	SupplierCode             string `xorm:"varchar(64)"`
	SupplierName             string `xorm:"varchar(128)"`
	RecipientCode            string `xorm:"varchar(64)"`
	RecipientName            string `xorm:"varchar(128)"`
	MaterialCode             string `xorm:"varchar(64)"`
	MaterialName             string `xorm:"varchar(128)"`
	Specification            string `xorm:"varchar(64)"`
	Model                    string `xorm:"varchar(64)"`
	CarNumber                string `xorm:"varchar(64)"`
	GrossWeighHouseCode      string `xorm:"varchar(64)"`
	GrossWeighMachineCode    string `xorm:"varchar(64)"`
	GrossWeight              float64
	GrossWeighTime           time.Time
	GrossWeighManCode        string `xorm:"varchar(64)"`
	GrossWeighManName        string `xorm:"varchar(64)"`
	GrossWeighSupervisorCode string `xorm:"varchar(64)"`
	GrossWeighSupervisor     string `xorm:"varchar(64)"`
	GrossWeighShift          string `xorm:"varchar(64)"`
	TareWeighHouseCode       string `xorm:"varchar(64)"`
	TareWeighMachineCode     string `xorm:"varchar(64)"`
	TareWeight               float64
	TareWeighTime            time.Time
	TareWeighManCode         string `xorm:"varchar(64)"`
	TareWeighManName         string `xorm:"varchar(64)"`
	TareWeighSupervisorCode  string `xorm:"varchar(64)"`
	TareWeighSupervisor      string `xorm:"varchar(64)"`
	TareWeighShift           string `xorm:"varchar(64)"`
	IsManualInputTare        bool
	Suttle                   float64
	MeasureUnit              string `xorm:"varchar(64)"`
	WeighTime                time.Time
	Deduction                float64
	DeductionPercent         float64
	TwiceDeduction           float64
	TwiceDeductionPercent    float64
	MoistureContent          float64
	IsManualInput            bool
	OperateBit               int8
	UploadBit                int8
	UploadTime               time.Time
	IsSurplus                bool
	IsReturnPurchase         bool
	CreateUserCode           string `xorm:"varchar(64)"`
	CreateUserName           string `xorm:"varchar(64)"`
	CreateTime               time.Time
	LastModifiedUserCode     string `xorm:"varchar(64)"`
	LastModifiedUserName     string `xorm:"varchar(64)"`
	LastModifiedTime         time.Time
	Remark                   string
	Attribute1               string
	Attribute2               string
	Attribute3               string
	Attribute4               string
	Attribute5               string
	Attribute6               string
	Attribute7               string
	Attribute8               string
	Attribute9               string
	Attribute10              string
	Attribute11              string
	Attribute12              string
	Attribute13              string
	Attribute14              string
	Attribute15              string
	Attribute16              string
}
