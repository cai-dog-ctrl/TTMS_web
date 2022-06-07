package service

import (
	"TTMS/dao/mysql"
	"TTMS/models"
	"TTMS/pkg/utils"
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
	if err != nil && len(p.List) != 0{
		return nil, err
	}
	return p, nil
}

func CountSalesByDay(day string) (*models.OrderDataList, error) {
	p, err := mysql.CountSalesByDay(day)
	if err != nil && len(p.List) != 0{
		return nil, err
	}
	return p, nil
}

func CountSalesByMonth(month string) (*models.OrderDataList, error) {
	p, err := mysql.CountSalesByMonth(month)
	if err != nil && len(p.List) != 0{
		return nil, err
	}
	return p, nil
}

func CountSalesByYear(year string) (*models.OrderDataList, error) {
	p, err := mysql.CountSalesByYear(year)
	if err != nil && len(p.List) != 0{
		return nil, err
	}
	return p, nil
} 