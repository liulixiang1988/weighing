package main

import (
	"fmt"
	"github.com/liulixiang1988/weighing/models"
)

func main() {
	fmt.Println("成功！")

	cartare := &models.Tb_weigh_cartare{}
	has, err := models.Pqorm.Where("Id=?", 1).Get(cartare)
	if err != nil {
		fmt.Println("查找cartare失败：", err)
		return
	}
	if has == false {
		count, err := models.Pqorm.Insert(&models.Tb_weigh_cartare{
			CarNumber:  "皖G12345",
			OperateBit: 1,
		})
		if err != nil {
			fmt.Println("插入cartare失败:", err)
			return
		}
		fmt.Println("插入", count, "个cartare")
	}

}
