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
	resp := make([]string, nr.GetDisplayCount())

	copy(temp, restaurants)

	for i := 0; i < nr.GetDisplayCount(); i++ {
		maxIndex := i
		for j := i + 1; j < len(temp); j++ {
			if temp[j].Rating() < temp[maxIndex].Rating() {
				continue
			}

			if (temp[j].Rating() == temp[maxIndex].Rating()) && temp[j].OnboardedTime().Before(temp[maxIndex].OnboardedTime()) {
				continue
			}

			maxIndex = j
		}
		temp[i], temp[maxIndex] = temp[maxIndex], temp[i]
		resp[i] = temp[i].RestaurantId()
	}

	respCh <- resp
	close(respCh)
}
