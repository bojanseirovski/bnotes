package Routes

import (
	"bojanseirovski/bnotes/Controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	r.Use(cors.New(config))

	grp1 := r.Group("/api")
	{
		grp1.GET("user", Controllers.GetUsers)
		//	Manage user
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetUserByID)
		grp1.PUT("user/:id", Controllers.UpdateUser)
		grp1.DELETE("user/:id", Controllers.DeleteUser)
		//	list/add user points
		grp1.GET("user/:id/points", Controllers.GetPoints)
		grp1.POST("user/:id/points", Controllers.AddPoint)
		//	get all user challenges
		grp1.GET("user/:id/challenge", Controllers.GetUserChallengesByUID)

		//	create a challenge
		grp1.POST("challenge", Controllers.AddChallenge)
		//	list all challenges
		grp1.GET("challenge", Controllers.GetChallenges)
		//	get a single challenge
		grp1.GET("challenge/:id", Controllers.GetChallengeByID)
		//	list all users per challenge
		grp1.GET("challenge/:id/user", Controllers.GetChallengeUsers)
		//	add user to challenge
		grp1.PUT("challenge/:cid/user/:uid", Controllers.AdduserToChallenge)
		//	list all user points for challenge
		grp1.GET("challenge/:id/user/:uid", Controllers.GetChallengeUserPoints)
		//	list total challenge points
		grp1.GET("challenge/:id/points", Controllers.GetPoints)
	}
	return r
}
