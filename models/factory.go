package models

import (
	"math/rand"
	"strconv"
	"time"
)

func createCuisineTrackings() []CuisineTracking {
	res := []CuisineTracking{}
	for i := 0; i < 10; i++ {
		res = append(res, CuisineTracking{
			typ:        Cuisine(i),
			noOfOrders: rand.Intn(100),
		})
	}
	return res
}

func createCostTrackings() []CostTracking {
	res := []CostTracking{}
	for i := 0; i < 10; i++ {
		res = append(res, CostTracking{
			typ:        i,
			noOfOrders: rand.Intn(100),
		})
	}
	return res
}

func CreateUser(cuisines []CuisineTracking, costs []CostTracking) User {
	if cuisines == nil {
		cuisines = createCuisineTrackings()
	}
	if costs == nil {
		costs = createCostTrackings()
	}
	return User{cuisines: cuisines, costBrackets: costs}
}

func CreateRestaurants(count int) []Restaurant {
	resp := []Restaurant{}
	for i := 0; i < count; i++ {
		resp = append(resp, Restaurant{
			restaurantId:  "Restaurant-" + strconv.Itoa(i),
			cuisine:       Cuisine(rand.Intn(15)),
			costBracket:   rand.Intn(15),
			rating:        rand.Float64() * 5,
			isRecommended: rand.Intn(2) == 0,
			onboardedTime: time.Now().Add(time.Duration(rand.Intn(1000) * int(time.Minute))),
		})
	}
	return resp
}
