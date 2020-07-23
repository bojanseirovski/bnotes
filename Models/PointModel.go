package Models

import "time"

type Point struct {
	Id    int       `json:"id"`
	Point int       `json:"point"`
	Uid   int       `json:"uid"`
	Date  time.Time `json:"date"`
}

func (b *Point) TableName() string {
	return "point"
}
