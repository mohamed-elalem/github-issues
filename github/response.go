package github

type Response struct {
	TotalCount int `json:"total_count"`
	Issues     `json:"items"`
}
