package engines

import (
	"fmt"

	"github.com/lokanadham100/restaurant/behaviour"
	"github.com/lokanadham100/restaurant/models"
)

type scPcbRl45 struct {
}

func init() {
	Register(&scPcbRl45{})
}

func (_ *scPcbRl45) GetOrder() int {
	return 800
}

func (_ *scPcbRl45) GetDisplayCount() int {
	return -1
}

func (s *scPcbRl45) GetRestaurantIds(tagger behaviour.RestaurantTagger, restaurants []models.Restaurant, respCh chan<- []string) {
	resp := []string{}
	for _, restaurant := range restaurants {
		if restaurant.Rating() < 4.5 && tagger.IsSecondaryCuisine(restaurant.RestaurantId()) && tagger.IsPrimaryCostBracket(restaurant.RestaurantId()) {
			resp = append(resp, restaurant.RestaurantId())
		}
	}
	fmt.Println("scPcbRl45: ", resp)
	respCh <- resp
	close(respCh)
}
