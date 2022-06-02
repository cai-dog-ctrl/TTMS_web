package mysql

import (
	"TTMS/models"
	"TTMS/pkg/snowflake"
	"fmt"
	"time"

	"go.uber.org/zap"
)

//有关票务的持久化代码

func SaleTicket(order *models.Order) (bool, error) {
	day := time.Now().Day()
	month := int(time.Now().Month())
	year := time.Now().Year()
	date := fmt.Sprintf("%v-%02v-%02v", year, month, day)
	hour := time.Now().Hour()
	minute := time.Now().Minute()
	second := time.Now().Second()
	now := fmt.Sprintf("%02v:%02v:%02v", hour, minute, second)
	order_id := snowflake.GenID()
	
	tx, err := db.Begin()
	if err != nil {
		return false, err
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
		ticket := new(models.Tick)
		sqlStr1 := fmt.Sprintf("select status from ticket where id = %v", ticket_id)
		err1 := db.Get(&ticket.Status, sqlStr1)
		if err1 != nil || ticket.Status == 1{
			zap.L().Error(sqlStr1)
			return false, err1
		}
		sqlStr := "update ticket set status = ? where id = ? and status = -1"
		_, err := db.Exec(sqlStr, 1, ticket_id)
		if err != nil {
			zap.L().Error(sqlStr)
			return false, err
		}
	}

	for _, ticket_id := range order.TicketList {
		sqlStr := "insert into order_info (id, user_id, ticket_id, date, time, is_delete) values(?, ?, ?, ?, ?, ?)"
		_, err := db.Exec(sqlStr, order_id, order.UserID, ticket_id, date, now, -1)
		if err != nil {
			zap.L().Error(sqlStr)
			return false, err
		}
	}

	return true, nil
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

