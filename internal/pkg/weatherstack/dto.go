package weatherstack

type Response struct {
	Current struct {
		Temperature float64 `json:"temperature"`
	} `json:"current"`
}
