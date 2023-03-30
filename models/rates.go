package models

type Rate struct {
	Currency    string  `json:"coin"`
	BuyRate     float64 `json:"buy_rate"`  // how much DOP it takes to buy 1 COIN
	SellRate    float64 `json:"sell_rate"` // how much DOP it you make from selling 1 COIN
	UpdatedDate string  `json:"updated_date"`
	Source      string  `json:"source"`
}
