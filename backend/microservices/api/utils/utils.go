package utils

import (
	"errors"
	"time"
)

type RegionType int

var REGIONS = [...]string{"Krasnodar", "Nizhny Novgorod", "Minsk", "Italy", "Istanbul", "Yerevan"}

func ValidateRegion(region string) error {
	for _, r := range REGIONS {
		if r == region {
			return nil
		}
	}
	return errors.New("no such region")
}

func Stod(d string) (time.Time, error) {
	layout := "02-01-2006"
	t, err := time.Parse(layout, d)

	return t, err
}
