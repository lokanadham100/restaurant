package engines

import (
	"github.com/lokanadham100/restaurant/behaviour"
	"github.com/lokanadham100/restaurant/models"
)

type newRestaurants struct {
}

func init() {
	Register(&newRestaurants{})
}

func (_ *newRestaurants) GetOrder() int {
	return 500
}

func (_ *newRestaurants) GetDisplayCount() int {
	return 4
}

func (nr *newRestaurants) GetRestaurantIds(_ behaviour.RestaurantTagger, restaurants []models.Restaurant, respCh chan<- []string) {
	temp := make([]models.Restaurant, len(restaurants))
	resp := make([]string, len(restaurants))
	for i := 0; i < len(restaurants); i++ {
		temp[i] = restaurants[i]
	}
	for i := 0; i < len(temp); i++ {
		maxIndex := i
		for j := i + 1; j < len(temp); j++ {
			if temp[j].OnboardedTime().After(temp[maxIndex].OnboardedTime()) {
				maxIndex = j
			}
		}
		temp[i], temp[maxIndex] = temp[maxIndex], temp[i]
		resp[i] = temp[i].RestaurantId()
	}

	respCh <- resp
	close(respCh)
}
