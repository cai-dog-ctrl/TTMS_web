package service

import (
	"TTMS/dao/mysql"
	"TTMS/models"
	"TTMS/pkg/utils"
)

//有关票务的业务逻辑代码

func SaleTicket(p *models.ParamsSaleTicket) error {
	order := new(models.Order)
	order.UserID = utils.ShiftToNum64(p.UserID)
	order.TicketList = p.IDList
	err := mysql.SaleTicket(order)
	if err != nil {
		return err
	}
	return nil
}
