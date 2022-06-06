package service

import (
	"TTMS/dao/mysql"
	"TTMS/models"
	"TTMS/pkg/utils"
	"strings"
)

//有关票务的业务逻辑代码

func SaleTicket(p *models.ParamsSaleTicket) (bool, float32, error) {
	order := new(models.Order)
	order.UserID = utils.ShiftToNum64(p.UserID)
	order.TicketList = p.IDList
	return mysql.SaleTicket(order)
}
func GetTicketByScheduleId(id int64) ( *models.TickRetss, error) {
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

	p2 := new(models.TickRetss)
	len := len(p.List)

	for len >= p1.RowNum {
		arr := p.List[:p1.RowNum]
		p.List = p.List[p1.RowNum:]
		p2.List = append(p2.List, arr)
		len -= p1.RowNum
	}

	return p2, nil	
}

func GetTicketByMovieIdAndDateDay(movie_id, date_day string) (*models.TickRets, error) {
	
	id := utils.ShiftToNum64(movie_id)
	
	str_day  := strings.Split(date_day, "-")
	day := str_day[0] + str_day[1] + str_day[2] 
	
	p1, err := mysql.GetTicketByMovieIdAndDateDay(id, utils.ShiftToNum64(day))
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

func GetTicketByCinemaIdAndDateDay(cinema_id, date_day string) (*models.TickRets, error) {
	id := utils.ShiftToNum64(cinema_id)

	str_day  := strings.Split(date_day, "-")
	day := str_day[0] + str_day[1] + str_day[2]

	p1, err := mysql.GetTicketByCinemaIdAndDateDay(id, utils.ShiftToNum64(day))

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

func Refund(ticket_id, user_id int64) error {
	err := mysql.Refund(ticket_id, user_id)
	return err
}