package Controllers

import (
	"bojanseirovski/bnotes/Models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//GetChallenges ... Get all Challenges
func GetChallenges(c *gin.Context) {
	var challenges []Models.Challenge
	err1 := Models.GetAllChallenges(&challenges)
	if err1 != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, challenges)
	}
}

//AddChallenge ... Create Challenge
func AddChallenge(c *gin.Context) {

	var challenge Models.Challenge
	c.BindJSON(&challenge)

	err1 := Models.CreateChallenge(&challenge)
	if err1 != nil {
		fmt.Println(err1.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, challenge)
	}
}

//GetChallengeUsers ... Get all users involved in a challenge
func GetChallengeUsers(c *gin.Context) {
	var id, err = strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	var users []Models.ChallengeUser
	var allChallengeUsers []Models.User

	err1 := Models.GetByChallengeID(&users, id)
	if err1 != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		for countI := 0; countI < len(users)-1; countI++ {
			var oneUser Models.User
			err2 := Models.GetUserByID(&oneUser, users[countI].UID)

			if err2 != nil {
				//	skip invalid
			} else {
				allChallengeUsers = append(allChallengeUsers, oneUser)
			}
		}
		c.JSON(http.StatusOK, allChallengeUsers)
	}
}

//AdduserToChallenge .. Add a single user to a challenge
func AdduserToChallenge(c *gin.Context) {
	var cid, err = strconv.Atoi(c.Params.ByName("cid"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}
	var uid, errUID = strconv.Atoi(c.Params.ByName("uid"))
	if errUID != nil {
		fmt.Println(errUID.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	var challengeuser Models.ChallengeUser
	challengeuser.ChallengeId = cid
	challengeuser.UID = uid

	c.BindJSON(&challengeuser)

	err1 := Models.AddUserChallenge(&challengeuser)
	if err1 != nil {
		fmt.Println(err1.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, challengeuser)
	}
}

//GetChallengeUserPoints ... Get all user points for a challenge
func GetChallengeUserPoints(c *gin.Context) {
	var id, err = strconv.Atoi(c.Params.ByName("id"))
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	var uid, erruid = strconv.Atoi(c.Params.ByName("uid"))
	if erruid != nil {
		fmt.Println(erruid.Error())
		c.AbortWithStatus(http.StatusNotFound)
	}

	var user Models.User
	var challenge Models.Challenge
	var points []Models.Point
	//	get and verify the user
	errUser := Models.GetUserByID(&user, uid)
	if errUser != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	//	get and veify the challenge
	errChallenge := Models.GetChallengeByID(&challenge, id)
	if errChallenge != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}

	//	get all points for the user between challenge dates
	errPoints := Models.GetPointsByUIDandInterval(&points, id, challenge.Start, challenge.End)
	if errPoints != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		var sumPoint int
		for countI := 0; countI < len(points)-1; countI++ {
			sumPoint += points[countI].Point
		}
		c.JSON(http.StatusOK, gin.H{"points": sumPoint})
	}

}
