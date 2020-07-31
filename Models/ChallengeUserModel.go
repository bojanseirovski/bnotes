package Models

type ChallengeUser struct {
	ID          int `json:"id"`
	UID         int `json:"uid"`
	ChallengeId int `json:"challenge_id"`
}

func (b *ChallengeUser) TableName() string {
	return "challenge_user"
}
