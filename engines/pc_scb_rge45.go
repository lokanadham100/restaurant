package engines

import (
	"fmt"

	"github.com/lokanadham100/restaurant/behaviour"
	"github.com/lokanadham100/restaurant/models"
)

type pcScbRge45 struct {
}

func init() {
	Register(&pcScbRge45{})
}

func (_ *pcScbRge45) GetOrder() int {
	return 300
}

func (_ *pcScbRge45) GetDisplayCount() int {
	return -1
}

func (s *pcScbRge45) GetRestaurantIds(tagger behaviour.RestaurantTagger, restaurants []models.Restaurant, respCh chan<- []string) {
	resp := []string{}
	for _, restaurant := range restaurants {
		if restaurant.Rating() >= 4.5 && tagger.IsPrimaryCuisine(restaurant.RestaurantId()) && tagger.IsSecondaryCostBracket(restaurant.RestaurantId()) {
			resp = append(resp, restaurant.RestaurantId())
		}
	}
	fmt.Println("pcScbRge45: ", resp)
	respCh <- resp
	close(respCh)
}
