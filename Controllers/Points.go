package Controllers

import (
	"bojanseirovski/bnotes/Models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//GetPoints ... Get all points for user
func GetPoints(c *gin.Context) {
	var id, err = strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	var points []Models.Point
	err1 := Models.GetPointsByUID(&points, id)
	if err1 != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, points)
	}
}

//AddPoint ... Create Point
func AddPoint(c *gin.Context) {
	var id, err = strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	var point Models.Point
	point.Uid = id
	point.Date = time.Now()
	point.Point = 1

	c.BindJSON(&point)
	err1 := Models.CreatePoint(&point)
	if err1 != nil {
		fmt.Println(err1.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, point)
	}
}
