package Models

import (
	"bojanseirovski/bnotes/Config"

	_ "github.com/go-sql-driver/mysql"
)

//AddUserChallenge ... Insert New data
func AddUserChallenge(challengeuser *ChallengeUser) (err error) {
	if err = Config.DB.Create(challengeuser).Error; err != nil {
		return err
	}
	return nil
}

//GetByChallengeID ... Fetch all users for challenge by cid
func GetByChallengeID(challengeuser *[]ChallengeUser, id int) (err error) {
	if err = Config.DB.Where("challenge_id = ?", id).Find(challengeuser).Error; err != nil {
		return err
	}
	return nil
}
