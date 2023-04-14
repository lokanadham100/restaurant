package behaviour

import (
	"github.com/lokanadham100/restaurant/models"
	"github.com/lokanadham100/restaurant/utils"
)

type UserTracker interface {
	GetPrimaryCuisine() models.Cuisine
	GetPrimaryCostBracket() int
	GetSecondaryCuisines() []models.Cuisine
	GetSecondaryCostBrackets() []int
}

type userTracker struct {
	cuisineSorter     *utils.Sorter
	costBracketSorter *utils.Sorter
}

func NewUserTracker(user models.User) UserTracker {
	ut := &userTracker{
		cuisineSorter:     utils.NewSorter(user.GetCuisinesForPQ(), 3),
		costBracketSorter: utils.NewSorter(user.GetCostBracketForPQ(), 3),
	}
	return ut
}

func (ut *userTracker) GetPrimaryCuisine() models.Cuisine {
	sortedItems := ut.cuisineSorter.Items()
	if len(sortedItems) == 0 {
		return -1
	}
	return models.Cuisine(sortedItems[0].GetValue())
}

func (ut *userTracker) GetPrimaryCostBracket() int {
	sortedItems := ut.costBracketSorter.Items()
	if len(sortedItems) == 0 {
		return -1
	}
	return int(sortedItems[0].GetValue())
}

func (ut *userTracker) GetSecondaryCuisines() []models.Cuisine {
	sortedItems := ut.cuisineSorter.Items()
	if len(sortedItems) <= 1 {
		return []models.Cuisine{-1}
	}
	return []models.Cuisine{models.Cuisine(sortedItems[1].GetValue()), models.Cuisine(sortedItems[2].GetValue())}
}

func (ut *userTracker) GetSecondaryCostBrackets() []int {
	sortedItems := ut.costBracketSorter.Items()
	if len(sortedItems) <= 1 {
		return []int{-1}
	}
	return []int{int(sortedItems[1].GetValue()), int(sortedItems[2].GetValue())}
}
