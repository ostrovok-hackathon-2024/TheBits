package ostrovok

type APIResponse struct {
	Data struct {
		Hotels []interface{} `json:"hotels"`
	} `json:"data"`
	//Debug  string `json:"debug"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

type Hotel struct {
}

type RequestBody struct {
	CheckIn     string  `json:"checkin"`
	CheckOut    string  `json:"checkout"`
	Residency   string  `json:"residency"`
	Language    string  `json:"language"`
	Guests      []Guest `json:"guests"`
	RegionId    int     `json:"region_id"`
	Currency    string  `json:"currency"`
	HotelsLimit int     `json:"hotels_limit"`
}

type Guest struct {
	Adults   int           `json:"adults"`
	Children []interface{} `json:"children"`
}

type HotelInfo struct {
		Title            string `json:"title"`
		CurrentPrice     int    `json:"currentPrice"`
		YandexPrice      int    `json:"YandexPrice"`
		YandexDiscount     int    `json:"YandexDiscount"`
		RecommendedPrice int    `json:"RecommendedPrice"`
}

