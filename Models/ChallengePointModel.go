package Models

type ChallengePoint struct {
	ID    int `json:"id"`
	CID   int `json:"cid"`
	Point int `json:"point"`
}

func (b *ChallengePoint) TableName() string {
	return "challenge_points"
}
