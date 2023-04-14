package behaviour

import "github.com/lokanadham100/restaurant/models"

type RestaurantTagger interface {
	IsPrimaryCuisine(string) bool
	IsSecondaryCuisine(string) bool
	IsPrimaryCostBracket(string) bool
	IsSecondaryCostBracket(string) bool
}

type tagger struct {
	primaryCuisineMap       map[string]struct{}
	primaryCostBracketMap   map[string]struct{}
	secondaryCuisineMap     map[string]struct{}
	secondaryCostBracketMap map[string]struct{}
}

func NewRestaurantTagger(userTracker UserTracker, restaurants []models.Restaurant) RestaurantTagger {
	tagger := &tagger{
		primaryCuisineMap:       map[string]struct{}{},
		primaryCostBracketMap:   map[string]struct{}{},
		secondaryCuisineMap:     map[string]struct{}{},
		secondaryCostBracketMap: map[string]struct{}{},
	}
	for _, restaurant := range restaurants {
		if restaurant.Cuisine() == userTracker.GetPrimaryCuisine() {
			tagger.primaryCuisineMap[restaurant.RestaurantId()] = struct{}{}
		} else {
			for _, secCuisine := range userTracker.GetSecondaryCuisines() {
				if restaurant.Cuisine() == secCuisine {
					tagger.secondaryCuisineMap[restaurant.RestaurantId()] = struct{}{}
					break
				}
			}
		}

		if restaurant.CostBracket() == userTracker.GetPrimaryCostBracket() {
			tagger.primaryCostBracketMap[restaurant.RestaurantId()] = struct{}{}
		} else {
			for _, secCostBracket := range userTracker.GetSecondaryCostBrackets() {
				if restaurant.CostBracket() == secCostBracket {
					tagger.secondaryCostBracketMap[restaurant.RestaurantId()] = struct{}{}
					break
				}
			}
		}
	}
	return tagger
}

func (t *tagger) IsPrimaryCuisine(id string) bool {
	_, ok := t.primaryCuisineMap[id]
	return ok
}

func (t *tagger) IsSecondaryCuisine(id string) bool {
	_, ok := t.secondaryCuisineMap[id]
	return ok
}

func (t *tagger) IsPrimaryCostBracket(id string) bool {
	_, ok := t.primaryCostBracketMap[id]
	return ok
}

func (t *tagger) IsSecondaryCostBracket(id string) bool {
	_, ok := t.secondaryCostBracketMap[id]
	return ok
}
