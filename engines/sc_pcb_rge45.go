package engines

import (
	"fmt"

	"github.com/lokanadham100/restaurant/behaviour"
	"github.com/lokanadham100/restaurant/models"
)

type scPcbRge45 struct {
}

func init() {
	Register(&scPcbRge45{})
}

func (_ *scPcbRge45) GetOrder() int {
	return 400
}

func (_ *scPcbRge45) GetDisplayCount() int {
	return -1
}

func (s *scPcbRge45) GetRestaurantIds(tagger behaviour.RestaurantTagger, restaurants []models.Restaurant, respCh chan<- []string) {
	resp := []string{}
	for _, restaurant := range restaurants {
		if restaurant.Rating() >= 4.5 && tagger.IsSecondaryCuisine(restaurant.RestaurantId()) && tagger.IsPrimaryCostBracket(restaurant.RestaurantId()) {
			resp = append(resp, restaurant.RestaurantId())
		}
	}
	fmt.Println("scPcbRge45: ", resp)
	respCh <- resp
	close(respCh)
}
