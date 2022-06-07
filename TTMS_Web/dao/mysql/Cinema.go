package mysql

//有关影院的持久化代码

import (
	"TTMS/models"
	"TTMS/pkg/snowflake"
	"TTMS/pkg/utils"
	"fmt"

	"go.uber.org/zap"
)

func GetAllCinemas(num int, page_num int) (*models.CinemaList, error) {
	cinemaList := new(models.CinemaList)
	sqlStr := "select count(*) from cinema_info"
	err := db.Get(&cinemaList.Number, sqlStr)
	if err != nil {
		zap.L().Error(sqlStr)
		return nil, err
	}
	sqlStr1 := "select * from cinema_info"
	err = db.Select(&cinemaList.CinemaList, sqlStr1)
	if err != nil {
		zap.L().Error(sqlStr)
		return nil, err
	}
	return cinemaList, nil
}

func GetCinemaByID(id int64) (*models.CinemaInfo, error) {
	cinema := new(models.CinemaInfo)
	sqlStr := fmt.Sprintf("select * from cinema_info where id = %v", id)
	err := db.Get(cinema, sqlStr)
	if err != nil {
		zap.L().Error(sqlStr)
		return nil, err
	}
	return cinema, nil
}

func InsertCinema(p *models.CinemaInfo) error {
	sqlStr := "insert into cinema_info (id, roww, coll, tag, is_delete, cinema_name) values (?, ?, ?, ?, -1, ?)"
	_, err := db.Exec(sqlStr, p.ID, p.MaxRow, p.MaxCol, p.Tag, p.Name)
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	for i := 0; i < p.MaxRow; i++ {
		for j := 0; j < p.MaxCol; j++ {
			seatinfo := new(models.SeatInfo)
			seatinfo.ID = snowflake.GenID()
			seatinfo.CinemaID = p.ID
			seatinfo.Row = i
			seatinfo.Col = j
			seatinfo.Status = 1
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
	sqlStr := "update cinema_info set roww = ?, coll = ?, tag = ?, is_delete = ? ,cinema_name = ? where id = ?"
	_, err := db.Exec(sqlStr, p.MaxRow, p.MaxCol, p.Tag, p.IsDelete, p.Name, p.ID)
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	return nil
}

func DeleteCinemaByID(id int64) error {
	sqlStr := "delete from cinema_info where id = ?"
	_, err := db.Exec(sqlStr, id)
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	sqlStr1 := "delete from seat_info where cinema_id = ?"
	_, err1 := db.Exec(sqlStr1, id)
	if err1 != nil {
		zap.L().Error(sqlStr1)
		return err1
	}
	return nil
}

func InsertSeat(p *models.SeatInfo) error {
	sqlStr := "insert into seat_info (id, cinema_id, roww, coll, status, flag) values (?, ?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, p.ID, p.CinemaID, p.Row, p.Col, p.Status, 3)
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	return nil
}

func GetSeatByCinemaID(id int64) (*models.SeatList, error) {
	seatList := new(models.SeatList)
	cinema := new(models.CinemaInfo)
	sqlStr1 := fmt.Sprintf("select roww from cinema_info where id = %v", id)
	err := db.Get(&cinema.MaxRow, sqlStr1)
	if err != nil {
		zap.L().Error(sqlStr1)
		return nil, err
	}
	sqlStr2 := fmt.Sprintf("select coll from cinema_info where id = %v", id)
	err = db.Get(&cinema.MaxCol, sqlStr2)
	if err != nil {
		zap.L().Error(sqlStr2)
		return nil, err
	}

	n := make([][]models.Seat, cinema.MaxRow)
	for i := range n {
		n[i] = make([]models.Seat, cinema.MaxCol)
	}
	for i := 0; i < cinema.MaxRow; i++ {
		sqlStr3 := fmt.Sprintf("select id, status from seat_info where cinema_id = %v and roww = %v", id, i)
		err = db.Select(&n[i], sqlStr3)
		if err != nil {
			zap.L().Error(sqlStr3)
			return nil, err
		}
	}
	seatList.ID = id
	seatList.SeatList = n
	return seatList, nil
}

func ModifySeat(p *models.ParamsModifySeat) error {
	sqlStr := ("update seat_info set status = ?, flag = ? where id = ?")
	_, err := db.Exec(sqlStr, p.Status, p.Flag, utils.ShiftToNum64(p.ID))
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	if p.Status == 2 {
		sqlStr1 := "update ticket set status = ? where ticket.seat_id = ?"
		_, err1 := db.Exec(sqlStr1, p.Status, utils.ShiftToNum64(p.ID))
		if err != nil {
			zap.L().Error(sqlStr1)
			return err1
		}
	}
	return nil
}
