package api

import "time"

type HotelInfo struct {
	Title            string `json:"title,omitempty"`
	CurrentPrice     int    `json:"current_price,omitempty"`
	YandexPrice      int    `json:"yandex_price,omitempty"`
	YandexDiscount   int    `json:"yandex_discount,omitempty"`
	RecommendedPrice int    `json:"recommended_price,omitempty"`
}

type ReqParams struct {
	Region   string `json:"region,omitempty"`
	Checkin  string `json:"checkin,omitempty"`
	Checkout string `json:"checkout,omitempty"`
}

type SearchParams struct {
	Region       string    `json:"region"`
	CheckinDate  time.Time `json:"checkin"`
	CheckoutDate time.Time `json:"checkout"`
}
