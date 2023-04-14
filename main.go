package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/lokanadham100/restaurant/behaviour"
	"github.com/lokanadham100/restaurant/engines"
	"github.com/lokanadham100/restaurant/models"
)

func GetRestaurantRecommendations(user models.User, availableRestaurants []models.Restaurant) []string {
	ut := behaviour.NewUserTracker(user)
	tagger := behaviour.NewRestaurantTagger(ut, availableRestaurants)
	return engines.GetRestaurantIds(tagger, availableRestaurants)
}

// Randomised Test Case
func main() {
	rand.Seed(time.Now().UnixNano())

	user := models.CreateUser(nil, nil)
	restos := models.CreateRestaurants(1000)
	ut := behaviour.NewUserTracker(user)
	tagger := behaviour.NewRestaurantTagger(ut, restos)

	result := GetRestaurantRecommendations(user, restos)

	// Formatting the result for readability
	restosMap := map[string]models.Restaurant{}
	for _, resto := range restos {
		restosMap[resto.RestaurantId()] = resto
	}

	fmt.Println("******************User Cuisines Trackings*********************")
	for _, cuisine := range user.GetCuisinesForPQ() {
		fmt.Println("Type: ", cuisine.GetValue(), "     No: ", cuisine.GetPriority())
	}

	fmt.Println("******************User Cost Trackings*********************")
	for _, cost := range user.GetCostBracketForPQ() {
		fmt.Println("Type: ", cost.GetValue(), "     No: ", cost.GetPriority())
	}

	fmt.Println("******************Restaurants*********************")
	for i, resto := range result {
		r := restosMap[resto]
		fmt.Println(i, " . ", " ", tagger.IsPrimaryCuisine(resto), " ", tagger.IsPrimaryCostBracket(resto), " ", tagger.IsSecondaryCuisine(resto), " ",
			tagger.IsSecondaryCostBracket(resto), " ID: ", r.RestaurantId(), "   Cuisine: ", r.Cuisine(), "   Cost: ", r.CostBracket(),
			"   rating: ", r.Rating(), "   Recommend: ", r.IsRecommended(), "   Date: ", r.OnboardedTime())
	}
}
