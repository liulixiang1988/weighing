package models

import (
	"log"

	"github.com/go-xorm/xorm"
	_ "github.com/lib/pq"
	_ "github.com/lunny/godbc"
)

var Mssql *xorm.Engine
var Pqorm *xorm.Engine

func init() {
	var err error
	Mssql, err = xorm.NewEngine("odbc", "driver={sql server};server=127.0.0.1;port=1433;uid=sa;pwd=tlys.oaxmz.5860247;database=tgzljl_czgl")
	if err != nil {
		log.Fatalf("err happened", err)
	}

	//创建pq orm引擎
	Pqorm, err = xorm.NewEngine("postgres", "host=127.0.0.1 port=5432 user=zljl password=123 dbname=tgzljl_czgl sslmode=disable")
	if err != nil {
		log.Fatalln("pq orm created error:", err)
	}

	//Tb_weigh_batch
	err = Pqorm.Sync(new(Tb_weigh_batch))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}

	//Tb_weigh_carcontrol
	err = Pqorm.Sync(new(Tb_weigh_carcontrol))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}

	//tb_weigh_cartare
	err = Pqorm.Sync(new(Tb_weigh_cartare))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}

	//tb_weigh_customer
	err = Pqorm.Sync(new(Tb_weigh_customer))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}

	//tb_weigh_supplier
	err = Pqorm.Sync(new(Tb_weigh_supplier))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}

	//tb_weigh_datalineinfo
	err = Pqorm.Sync(new(Tb_weigh_datalineinfo))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}

	err = Pqorm.Sync(new(Tb_weigh_datalineinfo_detail))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}

	err = Pqorm.Sync(new(Tb_weigh_datalineinfo_detail_cache))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}

	err = Pqorm.Sync(new(Tb_weigh_datalineinfobak))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}

	err = Pqorm.Sync(new(Tb_weigh_datalineinfoiface))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}

	err = Pqorm.Sync(new(Tb_weigh_plan))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}

	err = Pqorm.Sync(new(Tb_weigh_remark))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}

	err = Pqorm.Sync(new(Tb_weigh_shift))
	if err != nil {
		log.Fatalln("pq orm sync error:", err)
	}
}
