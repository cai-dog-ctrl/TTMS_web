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
