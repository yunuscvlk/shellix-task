package result

type Result struct {
	Status  bool   `json:"status"`
	Content string `json:"content"`
}