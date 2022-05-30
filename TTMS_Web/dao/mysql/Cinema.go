package mysql

//有关影院的持久化代码

import (
	"TTMS/models"
	"fmt"
	"go.uber.org/zap"
)

func InsertCinema(p *models.CinemaInfo) error {
	sqlStr := "insert into cinema_info (row, col, tag, is_delete) values (?, ?, ?, -1)"
	_, err := db.Exec(sqlStr, p.MaxRow, p.MaxCol, p.Tag)
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	var id int64
	sqlStr1 := "select last_insert_id()"
	err = db.Get(&id, sqlStr1)
	if err != nil {
		zap.L().Error(sqlStr1)
		return err
	}
	for i := 0; i < p.MaxRow; i++{
		for j := 0; j < p.MaxCol; j++{
			seatinfo := new(models.SeatInfo)
			seatinfo.ID = id
			seatinfo.Row = i
			seatinfo.Col = j
			seatinfo.State = 1
			err1 := InsertSeat(seatinfo)
			if err1 != nil {
				zap.L().Error(sqlStr)
				return err1
			}
		}
	}
	return nil
}

func ModifyCinema(p *models.CinemaInfo) error {
	sqlStr := "update movie_info set row = ?, col = ?, tag = ?, is_delete = ? where id = ?"
	_, err := db.Exec(sqlStr, p.MaxRow, p.MaxCol, p.Tag, p.IsDelete, p.ID)
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	return nil
}

func InsertSeat(p *models.SeatInfo) error{
	sqlStr := "insert into seat_info (id, row, col, status) values (?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, p.ID, p.Row, p.Col, p.State)
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	return nil
}

func GetSeatByCinemaID(id int64) (*models.SeatList, error) {
	seatList := new(models.SeatList)
	sqlStr1 := fmt.Sprintf("select row from cinema_info where id = %v", id)
	var row int
	err := db.Get(&row, sqlStr1)
	if err != nil {
		zap.L().Error(sqlStr1)
		return nil, err
	}
	sqlStr2 := fmt.Sprintf("select col from cinema_info where id = %v", id)
	var col int
	err = db.Get(&col, sqlStr2)
	if err != nil {
		zap.L().Error(sqlStr2)
		return nil, err
	}
	
	n := make([][]int, row)
	for i := range n {
		n[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			sqlStr := fmt.Sprintf("select status from seat_info where id = %v and col = %v and row = %v", id, j, i)
			err = db.Get(&n[i][j], sqlStr)
			if err != nil {
				zap.L().Error(sqlStr2)
				return nil, err
			}
		}
	}
	seatList.ID = id
	seatList.SeatList = n
	return seatList, nil
}

func ModifySeat(p *models.SeatList) error{
	sqlStr1 := fmt.Sprintf("select row from cinema_info where id = %v", p.ID)
	var row int
	err := db.Get(&row, sqlStr1)
	if err != nil {
		zap.L().Error(sqlStr1)
		return err
	}
	sqlStr2 := fmt.Sprintf("select col from cinema_info where id = %v", p.ID)
	var col int
	err = db.Get(&col, sqlStr2)
	if err != nil {
		zap.L().Error(sqlStr2)
		return err
	}
	for i := 0; i < row; i++{
		for j := 0; j < col; j++{
			sqlStr := fmt.Sprintf("update seat_info set status = %v where id = %v, row = %v, col = %v", p.SeatList[i][j], p.ID, i, j)
			_, err := db.Exec(sqlStr)
			if err != nil {
				zap.L().Error(sqlStr)
				return err
			}
		}
	}
	return nil
}