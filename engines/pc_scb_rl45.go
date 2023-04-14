package engines

import (
	"fmt"

	"github.com/lokanadham100/restaurant/behaviour"
	"github.com/lokanadham100/restaurant/models"
)

type pcScbRl45 struct {
}

func init() {
	Register(&pcScbRl45{})
}

func (_ *pcScbRl45) GetOrder() int {
	return 700
}
func (_ *pcScbRl45) GetDisplayCount() int {
	return -1
}

func (s *pcScbRl45) GetRestaurantIds(tagger behaviour.RestaurantTagger, restaurants []models.Restaurant, respCh chan<- []string) {
	resp := []string{}
	for _, restaurant := range restaurants {
		if restaurant.Rating() < 4.5 && tagger.IsPrimaryCuisine(restaurant.RestaurantId()) && tagger.IsSecondaryCostBracket(restaurant.RestaurantId()) {
			resp = append(resp, restaurant.RestaurantId())
		}
	}
	fmt.Println("pcScbRl45: ", resp)
	respCh <- resp
	close(respCh)
}
