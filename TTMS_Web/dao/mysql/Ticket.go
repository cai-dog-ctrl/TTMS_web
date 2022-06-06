package mysql

import (
	"TTMS/models"
	"TTMS/pkg/snowflake"
	"fmt"
	"time"

	"go.uber.org/zap"
)

//有关票务的持久化代码

func SaleTicket(order *models.Order) (bool, error, float32) {
	day := time.Now().Day()
	month := int(time.Now().Month())
	year := time.Now().Year()
	date := fmt.Sprintf("%v-%02v-%02v", year, month, day)
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	now := fmt.Sprintf("%02v:%02v:%02v", hour, minute, second)
	order_id := snowflake.GenID()
	var TotalPrice float32

	tx, err := db.Begin()
	if err != nil {
		return false, err, 0
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	for _, ticket_id := range order.TicketList {
		var TicketPrice float32
		ticket := new(models.Tick)
		schdule := new(models.ScheduleOut)
		// user := new(models.)
		sqlStr1 := fmt.Sprintf("select status from ticket where id = %v", ticket_id)
		err1 := db.Get(&ticket.Status, sqlStr1)
		if err1 != nil || ticket.Status == 1 {
			zap.L().Error(sqlStr1)
			return false, err1, 0
		}
		sqlStr := "update ticket set status = ? where id = ? and status = -1"
		_, err := db.Exec(sqlStr, 1, ticket_id)
		if err != nil {
			zap.L().Error(sqlStr)
			return false, err, 0
		}
		sqlStr3 := fmt.Sprintf("select showschdule.price from showschdule, ticket where ticket.id = %v and ticket.schedule_id = showschdule.id", ticket_id)
		err3 := db.Get(&schdule.Price, sqlStr3)
		if err3 != nil {
			zap.L().Error(sqlStr3)
			return false, err3, 0
		}
		// discount
		// sqlStr4 := fmt.Sprintf("select discount from user where id = %v", order.UserID)
		// err4 := db.Get(&schdule.Price, sqlStr4)
		// if err4 != nil{
		// 	zap.L().Error(sqlStr4)
		// 	return false, err4, 0
		// }
		//
		//TicketPrice = schdule.Price
		TotalPrice += TicketPrice
		sqlStr2 := "insert into order_info (id, user_id, ticket_id, date, time, is_delete, price) values(?, ?, ?, ?, ?, ?, ?)"
		_, err2 := db.Exec(sqlStr2, order_id, order.UserID, ticket_id, date, now, -1, TicketPrice)
		if err2 != nil {
			zap.L().Error(sqlStr)
			return false, err2, 0
		}

	}

	return true, nil, TotalPrice
}

func GetTicketByScheduleId(id int64) (*models.Ticks, error) {
	sqlStr := "select id, schedule_id, cinema_id, movie_id, seat_id, status from ticket where is_delete = -1"
	p := new(models.Ticks)
	err := db.Select(&p.List, sqlStr, id)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func GetTicketByMovieIdAndDateDay(movie_id, date_day int64) ( *models.Ticks, error) {
	p := new(models.Sche)
	sqlStr := "select id from showschdule where movie_id = ? and date_day = ? and is_delete = -1"
	err := db.Select(&p.Id, sqlStr, movie_id, date_day)
	if err != nil {
		return nil, err
	}

	sqlStr1 := "select id, schedule_id, cinema_id, movie_id, seat_id, status from ticket where schedule_id = ? and is_delete = -1 "
	p1 := new(models.Ticks)
	err = db.Select(&p1.List, sqlStr1, p.Id)
	if err != nil {
		return nil, err
	}
	return p1, nil
}

func GetTicketByCinemaIdAndDateDay(cinema_id, date_day int64) ( *models.Ticks, error) {
	p := new(models.Sche)
	sqlStr := "select id from showschdule where cinema_id = ? and date_day = ? and is_delete = -1"
	err := db.Select(&p.Id, sqlStr, cinema_id, date_day)
	if err != nil {
		return nil, err
	}

	sqlStr1 := "select id, schedule_id, cinema_id, movie_id, seat_id, status from ticket where schedule_id = ? and is_delete = -1 "
	p1 := new(models.Ticks)
	err = db.Select(&p1.List, sqlStr1, p.Id)
	if err != nil {
		return nil, err
	}
	return p1, nil
}

