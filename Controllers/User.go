package Controllers

import (
	"bojanseirovski/bnotes/Models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CResp struct {
	Challenges      []Models.ChallengeUser
	TotalChallenges int
}

type UResp struct {
	Id          uint   `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	TotalPoints int    `json:"totalpoints"`
}

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {

		var usersOut []UResp
		var oneUserOut UResp
		sumPoints := 0

		for cUs := 0; cUs < len(user); cUs++ {
			sumPoints = 0
			var allPoints []Models.Point
			Models.GetPointsByUID(&allPoints, int(user[cUs].Id))

			for countI := 0; countI < len(allPoints); countI++ {
				sumPoints = sumPoints + allPoints[countI].Point
			}

			oneUserOut.Id = user[cUs].Id
			oneUserOut.Name = user[cUs].Name
			oneUserOut.Email = user[cUs].Email
			oneUserOut.Phone = user[cUs].Phone
			oneUserOut.Address = user[cUs].Address
			oneUserOut.TotalPoints = sumPoints

			usersOut = append(usersOut, oneUserOut)
		}

		c.JSON(http.StatusOK, usersOut)
	}
}

//CreateUser ... Create User
func CreateUser(c *gin.Context) {
	var user Models.User
	c.BindJSON(&user)
	err := Models.CreateUser(&user)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//GetUserByID ... Get the user by id
func GetUserByID(c *gin.Context) {
	id, errUser := strconv.Atoi(c.Params.ByName("id"))
	if errUser != nil {
		fmt.Println(errUser.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	var user Models.User
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//UpdateUser ... Update the user information
func UpdateUser(c *gin.Context) {
	var user Models.User
	id, errUser := strconv.Atoi(c.Params.ByName("id"))
	if errUser != nil {
		fmt.Println(errUser.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	err := Models.GetUserByID(&user, id)
	if err != nil {
		c.JSON(http.StatusNotFound, user)
	}
	c.BindJSON(&user)
	err = Models.UpdateUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
	}
}

//DeleteUser ... Delete the user
func DeleteUser(c *gin.Context) {
	var user Models.User
	var id, errID = strconv.Atoi(c.Params.ByName("id"))
	if errID != nil {
		fmt.Println(errID.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	err := Models.DeleteUser(&user, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + strconv.Itoa(id): "is deleted"})
	}
}

//GetUserChallengesByUID ... Get all user challenges by UID
func GetUserChallengesByUID(c *gin.Context) {
	id, errUser := strconv.Atoi(c.Params.ByName("id"))
	if errUser != nil {
		fmt.Println(errUser.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	var challenges []Models.ChallengeUser
	err := Models.GetUserChallengesByUID(&challenges, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var outC CResp
		outC.TotalChallenges = len(challenges)
		outC.Challenges = challenges
		c.JSON(http.StatusOK, outC)
	}
}
