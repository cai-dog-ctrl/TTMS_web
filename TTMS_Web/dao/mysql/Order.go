package mysql

import (
	"TTMS/models"
	"TTMS/pkg/utils"
	"errors"
	"fmt"
	"strings"
	"time"
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
	startIdx := p.Num * (p.Page_num - 1)
	orderList := new(models.OrderFrontList)
	orderListRet := new(models.OrderFrontListRet)

	sqlStr0 := fmt.Sprintf("select count(distinct id) from order_info where user_id = %v", id)
	err0 := db.Get(&orderListRet.Total, sqlStr0)
	if err0 != nil {
		zap.L().Error(sqlStr0)
		return nil, err0
	}
	sqlStr := fmt.Sprintf("select distinct order_info.id, order_info.date, order_info.status, movie_info.name, movie_info.cover_img, cinema_info.cinema_name, showschdule.date_day, showschdule.start_time from order_info, movie_info, cinema_info,showschdule, ticket where order_info.user_id = %v and order_info.ticket_id = ticket.id and ticket.movie_id = movie_info.id and ticket.schedule_id = showschdule.id and ticket.cinema_id = cinema_info.id order by id desc limit %v, %v", id, startIdx, p.Num)
	err := db.Select(&orderList.OrderFrontList, sqlStr)
	if err != nil {
		zap.L().Error(sqlStr)
		return nil, err
	}
	for i, entry := range orderList.OrderFrontList {
		sqlStr := fmt.Sprintf("select roww, coll, price from order_info, ticket, seat_info where order_info.id = %v and ticket.id = order_info.ticket_id and seat_info.id = ticket.seat_id", entry.OrderID)
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
	sqlStr := fmt.Sprintf("select order_info.id, order_info.date, order_info.status, movie_info.name, movie_info.cover_img, cinema_info.cinema_name, showschdule.date_day, showschdule.start_time from order_info, movie_info, cinema_info, showschdule, ticket where order_info.id = %v and order_info.ticket_id = ticket.id and ticket.movie_id = movie_info.id and ticket.schedule_id = showschdule.id and ticket.cinema_id = cinema_info.id", id)
	err := db.Get(order, sqlStr)
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

func CountAllSales() (*models.OrderDataList, error) {
	sqlStr1 := "select id, name, img from movie_info order by id"
	p1 := new(models.MovieIds)
	err := db.Select(&p1.IDS, sqlStr1)
	if err != nil && len(p1.IDS) == 0 {
		fmt.Println("don't have movie in table")
		return nil, err
	}
	p2 := new(models.OrderDataList)
	for _, it := range p1.IDS {
		sqlStr2 := "select sum(price) from order_info where movie_id = ? and is_delete = -1 and status = 1"
		p := new(models.OrderData)
		p.TotalPrice = 0
		err = db.Get(&p.TotalPrice, sqlStr2, it.Id)
		if err != nil && p.TotalPrice != 0 {
			fmt.Println(it.Id, "is", p.TotalPrice)
			return p2, err
		}
		p.CoverImgPath = it.CoverImgPath
		p.MovieName = it.MovieName
		p2.List = append(p2.List, *p)
		p2.Total += p.TotalPrice
	}
	return p2, nil
}

func CountSalesByDay(day string) (*models.OrderDataList, error) {
	sqlStr1 := "select id, name, img from movie_info order by id"
	p1 := new(models.MovieIds)
	err := db.Select(&p1.IDS, sqlStr1)
	if err != nil && len(p1.IDS) != 0 {
		return nil, err
	}

	p2 := new(models.OrderDataList)
	for _, it := range p1.IDS {
		sqlStr2 := "select sum(price) from order_info where movie_id = ? and date = ? and status = 1 and is_delete = -1"
		p := new(models.OrderData)
		p.TotalPrice = 0
		err = db.Get(&p.TotalPrice, sqlStr2, it.Id, day)
		if err != nil && p.TotalPrice != 0 {
			return p2, err
		}
		p.CoverImgPath = it.CoverImgPath
		p.MovieName = it.MovieName
		p2.List = append(p2.List, *p)
		p2.Total += p.TotalPrice
	}
	return p2, nil
}

func CountSalesByMonth(month string) (*models.OrderDataList, error) {
	sqlStr1 := "select id, name, img from movie_info order by id"
	p1 := new(models.MovieIds)
	err := db.Select(&p1.IDS, sqlStr1)
	if err != nil && len(p1.IDS) != 0 {
		return nil, err
	}

	strs := strings.Split(month, "-")
	start := utils.ShiftToNum(strs[1])
	end := utils.ShiftToNum(strs[1]) + 1
	// day_start := strs[0] + "-" + utils.ShiftToStringFromInt64(int64(start)) + "-00"
	// day_end	  := strs[0] + "-" + utils.ShiftToStringFromInt64(int64(end)) 	+ "-00"

	day_start := fmt.Sprintf("%v-%02v-%02v", strs[0], utils.ShiftToStringFromInt64(int64(start)), "00")
	day_end := fmt.Sprintf("%v-%02v-%02v", strs[0], utils.ShiftToStringFromInt64(int64(end)), "00")

	p2 := new(models.OrderDataList)
	for _, it := range p1.IDS {
		//sqlStr2 := "select sum(price) from order_info where movie_id = ? and date > '?' and date < '?' and status = 1 and is_delete = -1"
		sqlStr2 := fmt.Sprintf("select sum(price) from order_info where movie_id = %v and date > '%v' and date < '%v' and status = 1 and is_delete = -1", it.Id, day_start, day_end)
		p := new(models.OrderData)
		// fmt.Println(sqlStr2)
		p.TotalPrice = 0
		//err = db.Get(&p.TotalPrice, sqlStr2, it.Id, day_start, day_end)
		err = db.Get(&p.TotalPrice, sqlStr2)
		if err != nil && p.TotalPrice != 0 {
			return p2, err
		}
		p.CoverImgPath = it.CoverImgPath
		p.MovieName = it.MovieName
		p2.List = append(p2.List, *p)
		p2.Total += p.TotalPrice
	}
	return p2, nil
}

func CountSalesByYear(year string) (*models.OrderDataList, error) {
	sqlStr1 := "select id, name, img from movie_info order by id"
	p1 := new(models.MovieIds)
	err := db.Select(&p1.IDS, sqlStr1)
	if err != nil && len(p1.IDS) != 0 {
		return nil, err
	}

	strs := strings.Split(year, "-")
	start := utils.ShiftToNum(strs[0])
	end := utils.ShiftToNum(strs[0]) + 1
	// day_start := utils.ShiftToStringFromInt64(int64(start)) + "-00" + "-00"
	// day_end	  := utils.ShiftToStringFromInt64(int64(end)) 	+ "-00" + "-00"

	day_start := fmt.Sprintf("%v-%02v-%02v", utils.ShiftToStringFromInt64(int64(start)), "00", "00")
	day_end := fmt.Sprintf("%v-%02v-%02v", utils.ShiftToStringFromInt64(int64(end)), "00", "00")

	p2 := new(models.OrderDataList)
	for _, it := range p1.IDS {
		// sqlStr2 := "select sum(price) from order_info where movie_id = ? and date > ? and date < ? and status = 1 and is_delete = -1"
		sqlStr2 := fmt.Sprintf("select sum(price) from order_info where movie_id = %v and date > '%v' and date < '%v' and status = 1 and is_delete = -1", it.Id, day_start, day_end)
		p := new(models.OrderData)
		p.TotalPrice = 0
		err = db.Get(&p.TotalPrice, sqlStr2)
		if err != nil && p.TotalPrice != 0 {
			return p2, err
		}
		p.CoverImgPath = it.CoverImgPath
		p.MovieName = it.MovieName
		p2.List = append(p2.List, *p)
		p2.Total += p.TotalPrice
	}
	return p2, nil
}

func RefundOrder(Id int64) (bool, error) {
	p1 := new(models.ROrderList)

	sqlStr1 := "select ticket_id, user_id from order_info where id = ? and is_delete = -1 and status = 1"
	db.Select(&p1.OrderList, sqlStr1, Id)

	sqlStr2 := "select schedule_id from ticket where id = ? and is_delete = -1 and status = 1"
	db.Get(&p1.ScheduleId, sqlStr2)

	sqlStr3 := "select end_time, date_day from showschdule where id = ? and is_delete = -1"
	db.Get(&p1, sqlStr3)

	day := time.Now().Day()
	month := int(time.Now().Month())
	year := time.Now().Year()
	date := int64(year)*10000 + int64(month)*100 + int64(day)
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	now := int64(hour)*100 + int64(minute)

	if p1.DateDay < date {
		return false, errors.New("date 超时")
	}
	if p1.DateDay == date {
		if p1.EndTime < now || p1.EndTime == now {
			return false, errors.New("date 超时")
		}
	}

	sqlStr4 := "update order_info set is_delete = 1 where id = ? and status = 1"
	_, err := db.Exec(sqlStr4, Id)
	if err != nil {
		return false, errors.New("update order fail")
	}
	for _, it := range p1.OrderList {
		sqlStr5 := "update ticket set status = -1 where id = ? and is_delete = -1"
		_, err = db.Exec(sqlStr5, it.TicketID)
		if err != nil {
			return false, errors.New("update ticket fail")
		}
	}

	return true, nil
}