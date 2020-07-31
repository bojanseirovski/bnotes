package Controllers

import (
	"bojanseirovski/bnotes/Models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetUsers ... Get all users
func GetUsers(c *gin.Context) {
	var user []Models.User
	err := Models.GetAllUsers(&user)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, user)
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
