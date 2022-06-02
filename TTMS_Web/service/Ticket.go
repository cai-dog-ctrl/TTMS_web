package service

import (
	"TTMS/dao/mysql"
	"TTMS/models"
	"TTMS/pkg/utils"
)

//有关票务的业务逻辑代码

func SaleTicket(p *models.ParamsSaleTicket) (bool, error, float32) {
	order := new(models.Order)
	order.UserID = utils.ShiftToNum64(p.UserID)
	order.TicketList = p.IDList
	return mysql.SaleTicket(order)
}
func GetTicketByScheduleId(id int64) ( *models.TickRets, error) {
	p1, err := mysql.GetTicketByScheduleId(id)
	if err != nil {
		return nil, err
	}

	p := new(models.TickRets)
	p.Total = p1.Total
	for _,it := range p1.List {
		p.List = append(p.List, models.TickRet{
			Id: utils.ShiftToStringFromInt64(it.Id),
			ScheduleId: utils.ShiftToStringFromInt64(it.ScheduleId),
			CinemaId: utils.ShiftToStringFromInt64(it.CinemaId),
			MovieId: utils.ShiftToStringFromInt64(it.MovieId),
			SeatId: utils.ShiftToStringFromInt64(it.SeatId),
			Status:  it.Status,
		})
	}

	return p, nil	
}


