package weatherapi

type Response struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}
