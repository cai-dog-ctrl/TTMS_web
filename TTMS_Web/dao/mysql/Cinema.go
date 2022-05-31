package mysql

//有关影院的持久化代码

import (
	"TTMS/models"
	"fmt"
	"go.uber.org/zap"
)

func GetAllCinemas(num int, page_num int) (*models.CinemaList, error){
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

func GetCinemaByID(id int64)(*models.CinemaInfo, error) {
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
	sqlStr := "insert into cinema_info (roww, coll, tag, is_delete, cinema_name) values (?, ?, ?, -1, ?)"
	_, err := db.Exec(sqlStr, p.MaxRow, p.MaxCol, p.Tag, p.Name)
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
	var num int64
	for i := 0; i < p.MaxRow; i++{
		for j := 0; j < p.MaxCol; j++{
			seatinfo := new(models.SeatInfo)
			seatinfo.ID = num
			seatinfo.CinemaID = id
			seatinfo.Row = i
			seatinfo.Col = j
			seatinfo.Status = 1
			err1 := InsertSeat(seatinfo)
			if err1 != nil {
				zap.L().Error(sqlStr)
				return err1
			}
			num++
		}
	}
	return nil
}

func ModifyCinema(p *models.CinemaInfo) error {
	sqlStr := "update movie_info set roww = ?, coll = ?, tag = ?, is_delete = ? cinema_name = ? where id = ?"
	_, err := db.Exec(sqlStr, p.MaxRow, p.MaxCol, p.Tag, p.IsDelete, p.Name, p.ID)
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	return nil
}

func InsertSeat(p *models.SeatInfo) error{
	sqlStr := "insert into seat_info (id, cinema_id, roww, coll, status) values (?, ?, ?, ?, ?)"
	_, err := db.Exec(sqlStr, p.ID, p.CinemaID, p.Row, p.Col, p.Status)
	if err != nil {
		zap.L().Error(sqlStr)
		return err
	}
	return nil
}

func GetSeatByCinemaID(id int64) (*models.SeatList, error) {
	seatList := new(models.SeatList)
	sqlStr1 := fmt.Sprintf("select roww from cinema_info where id = %v", id)
	var row int
	err := db.Get(&row, sqlStr1)
	if err != nil {
		zap.L().Error(sqlStr1)
		return nil, err
	}
	sqlStr2 := fmt.Sprintf("select coll from cinema_info where id = %v", id)
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
			sqlStr := fmt.Sprintf("select status from seat_info where id = %v and coll = %v and roww = %v", id, j, i)
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
	sqlStr1 := fmt.Sprintf("select roww from cinema_info where id = %v", p.ID)
	var row int
	err := db.Get(&row, sqlStr1)
	if err != nil {
		zap.L().Error(sqlStr1)
		return err
	}
	sqlStr2 := fmt.Sprintf("select coll from cinema_info where id = %v", p.ID)
	var col int
	err = db.Get(&col, sqlStr2)
	if err != nil {
		zap.L().Error(sqlStr2)
		return err
	}
	for i := 0; i < row; i++{
		for j := 0; j < col; j++{
			sqlStr := fmt.Sprintf("update seat_info set status = %v where id = %v, roww = %v, coll = %v", p.SeatList[i][j], p.ID, i, j)
			_, err := db.Exec(sqlStr)
			if err != nil {
				zap.L().Error(sqlStr)
				return err
			}
		}
	}
	return nil
}