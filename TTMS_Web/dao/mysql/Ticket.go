package mysql

import (
	"TTMS/models"
	"TTMS/pkg/snowflake"
	"fmt"
	"time"

	"go.uber.org/zap"
	"golang.org/x/tools/go/analysis/passes/nilfunc"
)

//有关票务的持久化代码

func SaleTicket(order *models.Order) error {
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
		return err
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
		ticket := new(models)
		sqlStr1 := fmt.Sprintf("select status from ticket where id = %v", ticket_id)
		err := db.Get(&)
		sqlStr := "update ticket set status = ? where id = ? and status = -1"
		_, err := db.Exec(sqlStr, 1, ticket_id)
		if err != nil {
			zap.L().Error(sqlStr)
			return err
		}
	}

	for _, ticket_id := range order.TicketList {
		sqlStr := "insert into order_info (id, user_id, ticket_id, date, time, is_delete) values(?, ?, ?, ?, ?, ?)"
		_, err := db.Exec(sqlStr, order_id, order.UserID, ticket_id, date, now, -1)
		if err != nil {
			zap.L().Error(sqlStr)
			return err
		}
	}
	return nil
}
