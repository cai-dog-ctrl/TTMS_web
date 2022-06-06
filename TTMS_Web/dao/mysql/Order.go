package mysql

import (
	"TTMS/models"
	"TTMS/pkg/utils"
	"fmt"
	"go.uber.org/zap"
)

func PayMoneyByOrderID(id int64) error {
	sqlStr := "update order_info set status = ? where id = ?"
	_, err := db.Exec(sqlStr, 1, id)
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	return nil
}

func GetOrderByUserID(p *models.ParamsGetOrderByUserID) (*models.OrderFrontListRet, error) {
	id := utils.ShiftToNum64(p.ID)
	startIdx := p.Num * (p.Page_num -1)
	orderList := new(models.OrderFrontList)
	orderListRet := new(models.OrderFrontListRet)

	sqlStr0 := fmt.Sprintf("select count(*) from order_info where user_id = %v", id)
	err0 := db.Get(&orderListRet.Total, sqlStr0)
	if err0 != nil {
		zap.L().Error(sqlStr0)
		return nil, err0
	}
	sqlStr := fmt.Sprintf("select order_info.id, order_info.date, order_info.status, movie.name, movie.cover_img, cinema_info.name, showschedule.date_day, showschedule.start_time from order_info, movie_info, showschedule where order_info.user_id = %v and order_info.ticket_id = ticket.id and ticket.movie_id = movie_info.id and ticket.schedule_id = showschedule.id and ticket.cinema_id = cinema_info.id limit %v, %v", id, startIdx, p.Num)
	err := db.Select(&orderList.OrderFrontList, sqlStr)
	if err != nil {
		zap.L().Error(sqlStr)
		return nil, err
	}
	for i, entry := range orderList.OrderFrontList {
		sqlStr := fmt.Sprintf("select roww, coll, price from order_info, ticket, seat_info where order_info.id = %v and ticket.id = order_info.ticket_id and seat_info.id = ticket.seat_id", id)
		err := db.Select(&orderList.OrderFrontList[i].SeatList, sqlStr)
		if err != nil {
			zap.L().Error(sqlStr)
			return nil, err
		}
		for _, en := range orderList.OrderFrontList[i].SeatList {
			orderList.OrderFrontList[i].Price += en.Price
		}
		date_day := ""
		date_day += utils.ShiftToStringFromInt64(int64(entry.ScheduleDate / 10000))
		date_day += "-"
		day_time := entry.ScheduleDate % 10000
		date_day += utils.ShiftToStringFromInt64(int64(day_time / 100))
		date_day += "-"
		date_day += utils.ShiftToStringFromInt64(int64(day_time % 100))

		start_time := ""
		start_time += utils.ShiftToStringFromInt64(int64(entry.ScheduleTime / 100))
		start_time += ":"
		start_time += utils.ShiftToStringFromInt64(int64(entry.ScheduleTime % 100))
		tmp := models.OrderFrontRet{
			OrderID:      entry.OrderID,
			Date:         entry.Date,
			Price:        entry.Price,
			Status:       entry.Status,
			MovieName:    entry.MovieName,
			CoverImgPath: entry.CoverImgPath,
			CinemaName:   entry.CinemaName,
			SeatList:     entry.SeatList,
			ScheduleDate: date_day,
			ScheduleTime: start_time,
		}
		orderListRet.OrderFrontList = append(orderListRet.OrderFrontList, &tmp)
	}
	return orderListRet, nil
}

func GetOrderByID(id int64) (*models.OrderFrontRet, error) {
	order := new(models.OrderFront)
	sqlStr := fmt.Sprintf("select order_info.id, order_info.date, order_info.status, movie.name, movie.cover_img, cinema_info.name, showschedule.date_day, showschedule.start_time from order_info, movie_info, showschedule where order_info.id = %v and order_info.ticket_id = ticket.id and ticket.movie_id = movie_info.id and ticket.schedule_id = showschedule.id and ticket.cinema_id = cinema_info.id", id)
	err := db.Get(&order, sqlStr)
	if err != nil {
		zap.L().Error(sqlStr)
		return nil, err
	}
	sqlStr1 := fmt.Sprintf("select roww, coll, price from order_info, ticket, seat_info where order_info.id = %v and ticket.id = order_info.ticket_id and seat_info.id = ticket.seat_id", id)
	err1 := db.Select(&order.SeatList, sqlStr1)
	if err1 != nil {
		zap.L().Error(sqlStr1)
		return nil, err1
	}
	for _, en := range order.SeatList {
		order.Price += en.Price
	}
	date_day := ""
	date_day += utils.ShiftToStringFromInt64(int64(order.ScheduleDate / 10000))
	date_day += "-"
	day_time := order.ScheduleDate % 10000
	date_day += utils.ShiftToStringFromInt64(int64(day_time / 100))
	date_day += "-"
	date_day += utils.ShiftToStringFromInt64(int64(day_time % 100))

	start_time := ""
	start_time += utils.ShiftToStringFromInt64(int64(order.ScheduleTime / 100))
	start_time += ":"
	start_time += utils.ShiftToStringFromInt64(int64(order.ScheduleTime % 100))
	orderRet := models.OrderFrontRet{
		OrderID:      order.OrderID,
		Date:         order.Date,
		Price:        order.Price,
		Status:       order.Status,
		MovieName:    order.MovieName,
		CoverImgPath: order.CoverImgPath,
		CinemaName:   order.CinemaName,
		SeatList:     order.SeatList,
		ScheduleDate: date_day,
		ScheduleTime: start_time,
	}
	return &orderRet, nil
}
