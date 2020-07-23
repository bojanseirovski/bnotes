package Models

import (
	"bojanseirovski/bnotes/Config"

	_ "github.com/go-sql-driver/mysql"
)

//CreatePoint ... Insert New data
func CreatePoint(point *Point) (err error) {
	if err = Config.DB.Create(point).Error; err != nil {
		return err
	}
	return nil
}

//GetPointsByUID ... Fetch all points by uid
func GetPointsByUID(point *[]Point, id int) (err error) {
	if err = Config.DB.Where("uid = ?", id).Find(point).Error; err != nil {
		return err
	}
	return nil
}

//DeletePoint ... Delete user
func DeletePoint(point *Point, id int) (err error) {
	Config.DB.Where("id = ?", id).Delete(point)
	return nil
}
