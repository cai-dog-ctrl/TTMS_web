package service

//有关影院的业务逻辑代码

import (
	"TTMS/dao/mysql"
	"TTMS/models"
	"TTMS/pkg/utils"
)

func GetAllCinemas(p *models.ParamsGetAllCinemas) (*models.CinemaList, error) {
	return mysql.GetAllCinemas(p.Num, p.Page_num-1)
}

func GetCinemaByID(p *models.ParamsGetCinemaByID) (*models.CinemaInfo, error) {
	return mysql.GetCinemaByID(utils.ShiftToNum64(p.ID))

}

func AddNewCinema(p *models.ParamsAddNewCinema) error {
	Cinema := new(models.CinemaInfo)
	Cinema.Name = p.Name
	Cinema.Tag = p.Tag
	Cinema.MaxCol = p.MaxCol
	Cinema.MaxRow = p.MaxRow
	return mysql.InsertCinema(Cinema)
}

func ModifyCinemaByID(p *models.ParamsModifyCinema) error {
	Cinema := new(models.CinemaInfo)
	Cinema.ID = p.ID
	Cinema.Name = p.Name
	Cinema.Tag = p.Tag
	Cinema.MaxCol = p.MaxCol
	Cinema.MaxRow = p.MaxRow
	Cinema.IsDelete = p.IsDelete
	return mysql.ModifyCinema(Cinema)
}

func GetSeatByCinemaID(p *models.ParamsGetSeatByCinemaID) (*models.SeatList, error) {
	seatList, err := mysql.GetSeatByCinemaID(p.ID)
	if err != nil {
		return nil, err
	}
	return seatList, nil
}

func ModifySeat(p *models.ParamsModifySeat) error {
	seatlist := new(models.SeatList)
	seatlist.ID = p.ID
	seatlist.SeatList = p.SeatList
	return mysql.ModifySeat(seatlist)
}
