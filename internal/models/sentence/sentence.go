package sentence

import "time"

type Model struct {
	Id			int			`json:"id"`
	Sentence	string		`json:"sentence"`
	CreatedAt	*time.Time	`json:"created_at"`
}