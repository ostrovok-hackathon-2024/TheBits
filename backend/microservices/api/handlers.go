package api

import (
	"fmt"
	"math/rand"
	"net/http"

	"github.com/labstack/echo"
	"github.com/ostrovok-hackathon-2024/The-Bits/backend/microservices/api/utils"
)

func paramsHandler(c echo.Context) error {
	p := new(ReqParams)
	if err := c.Bind(p); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	fmt.Printf("Got: %+v \n", p)

	if err := utils.ValidateRegion(p.Region); err != nil {
		e := fmt.Sprintf("bad request: %e", err)
		return c.String(http.StatusBadRequest, e)
	}

	if p.Checkin == "" || p.Checkout == "" {
		return c.String(http.StatusBadRequest, "bad request: bag checkin or checkout date")
	}

	params := SearchParams{
		Region: p.Region,
	}

	if t, err := utils.Stod(p.Checkin); err != nil {
		e := fmt.Sprintf("bad request: bad checkin date %+v", err)
		return c.String(http.StatusBadRequest, e)
	} else {
		params.CheckinDate = t
	}

	if t, err := utils.Stod(p.Checkout); err != nil {
		e := fmt.Sprintf("bad request: bad checkout date %+v", err)
		return c.String(http.StatusBadRequest, e)
	} else {
		params.CheckoutDate = t
	}

	if params.CheckinDate.After(params.CheckoutDate) {
		return c.String(http.StatusBadRequest, "bad request: wrong date interval")
	}

	//send data to redis
	hotels := sendData(&params)

	return c.JSON(http.StatusCreated, hotels)
}

func sendData(params *SearchParams) *[]HotelInfo {
	hotels := make([]HotelInfo, 0)
	for i := 0; i < 20; i++ {
		hotels = append(hotels, HotelInfo{
			Title:            fmt.Sprintf("Hotel %d", i),
			CurrentPrice:     int(rand.Int31n(12000)),
			YandexPrice:      int(rand.Int31n(12000)),
			YandexDiscount:   int(rand.Int31n(100)),
			RecommendedPrice: int(rand.Int31n(12000)),
		})
	}
	fmt.Println("Got data: ", params)
	return &hotels
}
