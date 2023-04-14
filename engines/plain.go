package engines

import (
	"github.com/lokanadham100/restaurant/behaviour"
	"github.com/lokanadham100/restaurant/models"
)

type plainRecommender struct {
}

func init() {
	Register(&plainRecommender{})
}

func (_ *plainRecommender) GetOrder() int {
	return 900
}

func (_ *plainRecommender) GetDisplayCount() int {
	return -1
}

func (p *plainRecommender) GetRestaurantIds(_ behaviour.RestaurantTagger, restaurants []models.Restaurant, respCh chan<- []string) {
	resp := make([]string, len(restaurants))
	for i := 0; i < len(restaurants); i++ {
		resp[i] = restaurants[i].RestaurantId()
	}
	respCh <- resp
	close(respCh)
}
