package engines

import (
	"fmt"

	"github.com/lokanadham100/restaurant/behaviour"
	"github.com/lokanadham100/restaurant/models"
)

type pcPcbRl4 struct {
}

func init() {
	Register(&pcPcbRl4{})
}

func (_ *pcPcbRl4) GetOrder() int {
	return 600
}

func (_ *pcPcbRl4) GetDisplayCount() int {
	return -1
}

func (s *pcPcbRl4) GetRestaurantIds(tagger behaviour.RestaurantTagger, restaurants []models.Restaurant, respCh chan<- []string) {
	resp := []string{}
	for _, restaurant := range restaurants {
		if restaurant.Rating() < 4 && tagger.IsPrimaryCuisine(restaurant.RestaurantId()) && tagger.IsPrimaryCostBracket(restaurant.RestaurantId()) {
			resp = append(resp, restaurant.RestaurantId())
		}
	}
	fmt.Println("pcPcbRl4: ", resp)
	respCh <- resp
	close(respCh)
}
