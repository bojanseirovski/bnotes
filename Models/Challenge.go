package Models

import (
	"bojanseirovski/bnotes/Config"

	_ "github.com/go-sql-driver/mysql"
)

//CreateChallenge ... Insert New data
func CreateChallenge(challenge *Challenge) (err error) {
	if err = Config.DB.Create(challenge).Error; err != nil {
		return err
	}
	return nil
}

//GetAllChallenges ... Fetch all Challenges
func GetAllChallenges(challenge *[]Challenge) (err error) {
	if err = Config.DB.Find(challenge).Error; err != nil {
		return err
	}
	return nil
}

//GetChallengeByID ... Fetch a challenge by id
func GetChallengeByID(challenge *Challenge, id int) (err error) {
	if err = Config.DB.Where("id = ?", id).Find(challenge).Error; err != nil {
		return err
	}
	return nil
}
