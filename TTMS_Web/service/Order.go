package service

import (
	"TTMS/dao/mysql"
	"TTMS/models"
	"TTMS/pkg/utils"
	"fmt"
	"time"
)

func PayMoneyByOrderID(p *models.ParamsPayMoneyByOrderID) error {
	return mysql.PayMoneyByOrderID(utils.ShiftToNum64(p.ID))
}

func GetOrderByUserID(p *models.ParamsGetOrderByUserID) (*models.OrderFrontListRet, error) {
	return mysql.GetOrderByUserID(p)
}

func GetOrderByID(p *models.ParamsGetOrderByID) (*models.OrderFrontRet, error) {
	return mysql.GetOrderByID(utils.ShiftToNum64(p.ID))
}

func CountAllSales() (*models.OrderDataList, error) {
	p, err := mysql.CountAllSales()
	if err != nil && len(p.List) != 0 {
		return nil, err
	}
	year := time.Now().Year()
	day := time.Now().Day()
	month1 := int(time.Now().Month()) - 1
	if month1 == 0 {
		month1 = 12
		year -= 1
	}
	month2 := int(time.Now().Month())

	date1 := fmt.Sprintf("%v-%02v-%02v", year, month1, day)
	date2 := fmt.Sprintf("%v-%02v-%02v", year, month2, day)

	p1, err := mysql.CountSalesByMonth(date1)
	if err != nil && len(p1.List) != 0 {
		return nil, err
	}
	p2, err := mysql.CountSalesByMonth(date2)
	if err != nil && len(p2.List) != 0 {
		return nil, err
	}

	for i, _ := range p.List {
		var rate float32
		if p.List[i].TotalPrice != 0 {
			fmt.Println("s:", p1.List[i].TotalPrice, "x:", p2.List[i].TotalPrice)
		}
		p.List[i].Rate = 0
		if p1.List[i].TotalPrice > 0 {
			rate = (p2.List[i].TotalPrice - p1.List[i].TotalPrice) / p1.List[i].TotalPrice
		}
		if p1.List[i].TotalPrice == 0 && p2.List[i].TotalPrice > 0{
			rate = 1
		}
		p.List[i].Rate = int(rate * 100)
	}

	return p, nil
}

func CountSalesByDay(day string) (*models.OrderDataList, error) {
	p, err := mysql.CountSalesByDay(day)
	if err != nil && len(p.List) != 0 {
		return nil, err
	}
	return p, nil
}

func CountSalesByMonth(month string) (*models.OrderDataList, error) {
	p, err := mysql.CountSalesByMonth(month)
	if err != nil && len(p.List) != 0 {
		return nil, err
	}
	return p, nil
}

func CountSalesByYear(year string) (*models.OrderDataList, error) {
	p, err := mysql.CountSalesByYear(year)
	if err != nil && len(p.List) != 0 {
		return nil, err
	}
	return p, nil
}

func RefundOrder(Id string) (bool, error) {
	return mysql.RefundOrder(utils.ShiftToNum64(Id))	
}