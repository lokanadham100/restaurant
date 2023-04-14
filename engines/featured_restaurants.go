package engines

import (
	"fmt"

	"github.com/lokanadham100/restaurant/behaviour"
	"github.com/lokanadham100/restaurant/models"
)

type featuredRestaurants struct {
}

func init() {
	Register(&featuredRestaurants{})
}

func (_ *featuredRestaurants) GetOrder() int {
	return 100
}

func (_ *featuredRestaurants) GetDisplayCount() int {
	return -1
}

func (fr *featuredRestaurants) GetRestaurantIds(tagger behaviour.RestaurantTagger, restaurants []models.Restaurant, respCh chan<- []string) {
	primaryRestos := []string{}
	secondaryRestos := []string{}
	for _, restaurant := range restaurants {
		if !restaurant.IsRecommended() {
			continue
		}

		if tagger.IsPrimaryCuisine(restaurant.RestaurantId()) && tagger.IsPrimaryCostBracket(restaurant.RestaurantId()) {
			primaryRestos = append(primaryRestos, restaurant.RestaurantId())
			continue
		}

		if tagger.IsPrimaryCuisine(restaurant.RestaurantId()) && tagger.IsSecondaryCostBracket(restaurant.RestaurantId()) ||
			tagger.IsPrimaryCostBracket(restaurant.RestaurantId()) && tagger.IsSecondaryCuisine(restaurant.RestaurantId()) {
			secondaryRestos = append(secondaryRestos, restaurant.RestaurantId())
			continue
		}
	}

	fmt.Println("Featured Primary: ", primaryRestos)
	fmt.Println("Featured Secondary: ", secondaryRestos)

	respCh <- append(primaryRestos, secondaryRestos...)
	close(respCh)
}
