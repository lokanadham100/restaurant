package engines

import (
	"fmt"

	"github.com/lokanadham100/restaurant/behaviour"
	"github.com/lokanadham100/restaurant/models"
)

type pcPcbRge4 struct {
}

func init() {
	Register(&pcPcbRge4{})
}

func (_ *pcPcbRge4) GetOrder() int {
	return 200
}

func (_ *pcPcbRge4) GetDisplayCount() int {
	return -1
}

func (s *pcPcbRge4) GetRestaurantIds(tagger behaviour.RestaurantTagger, restaurants []models.Restaurant, respCh chan<- []string) {
	resp := []string{}
	for _, restaurant := range restaurants {
		if restaurant.Rating() >= 4 && tagger.IsPrimaryCuisine(restaurant.RestaurantId()) && tagger.IsPrimaryCostBracket(restaurant.RestaurantId()) {
			resp = append(resp, restaurant.RestaurantId())
		}
	}
	fmt.Println("pcPcbRge4: ", resp)
	respCh <- resp
	close(respCh)
}
