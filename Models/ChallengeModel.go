package Models

type Challenge struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Start  string `json:"start" time_format:"sql_date" time_utc:"true"`
	End    string `json:"end" time_format:"sql_date" time_utc:"true"`
	Points int    `json:"points"`
}

func (b *Challenge) TableName() string {
	return "challenge"
}
