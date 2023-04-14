package models

import "time"

//go:generate accessory -type Restaurant -receiver restaurant -output restaurant_accessor.go

type Restaurant struct {
	restaurantId  string    `accessor:"getter"`
	cuisine       Cuisine   `accessor:"getter"`
	costBracket   int       `accessor:"getter"`
	rating        float64   `accessor:"getter"`
	isRecommended bool      `accessor:"getter"`
	onboardedTime time.Time `accessor:"getter"`
}
