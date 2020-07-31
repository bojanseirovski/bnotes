package Models

import (
	"bojanseirovski/bnotes/Config"

	_ "github.com/go-sql-driver/mysql"
)

//CreateChallengePoint ... Insert New data
func CreateChallengePoint(challengepoint *ChallengePoint) (err error) {
	if err = Config.DB.Create(challengepoint).Error; err != nil {
		return err
	}
	return nil
}

//GetAllChallengePoins ... Get a sum of all points for a challenge
func GetAllChallengePoins() {

}
