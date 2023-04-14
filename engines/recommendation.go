package engines

import (
	"sort"
	"sync"

	"github.com/lokanadham100/restaurant/behaviour"
	"github.com/lokanadham100/restaurant/models"
	"github.com/lokanadham100/restaurant/utils"
)

var once sync.Once
var recommendationEngine *RecommendationEngine

const RestaurantsCount = 100

type RecommendationEngine struct {
	subEngines []SubEngine
}

type SubEngine interface {
	GetOrder() int
	GetDisplayCount() int
	GetRestaurantIds(behaviour.RestaurantTagger, []models.Restaurant, chan<- []string)
}

func Register(subEngine SubEngine) {
	once.Do(func() {
		recommendationEngine = &RecommendationEngine{
			subEngines: []SubEngine{},
		}
	})
	subEngines := recommendationEngine.subEngines

	i := sort.Search(len(subEngines), func(i int) bool { return subEngines[i].GetOrder() >= subEngine.GetOrder() })
	if i < len(subEngines) && subEngines[i].GetOrder() == subEngine.GetOrder() {
		panic("Two Recommendation SubEngines Have Same Priority")
	}

	if len(subEngines) <= i {
		subEngines = append(subEngines, subEngine)
	} else {
		subEngines = append(subEngines[:i+1], subEngines[i:]...)
		subEngines[i] = subEngine
	}

	recommendationEngine.subEngines = subEngines
}

func (recommendationEngine RecommendationEngine) GetRestaurantIds(tagger behaviour.RestaurantTagger, restaurants []models.Restaurant) []string {
	res := []string{}
	chs := []<-chan []string{}
	for _, subEngine := range recommendationEngine.subEngines {
		ch := make(chan []string)
		go subEngine.GetRestaurantIds(tagger, restaurants, ch)
		chs = append(chs, ch)
	}

	dr := utils.NewDupRemover()

	for i, ch := range chs {
		res = dr.MergeWithOutDuplicates(res, <-ch, recommendationEngine.subEngines[i].GetDisplayCount())
	}

	if len(res) <= RestaurantsCount {
		return res
	}
	return res[:RestaurantsCount]
}

func GetRestaurantIds(tagger behaviour.RestaurantTagger, restaurants []models.Restaurant) []string {
	return recommendationEngine.GetRestaurantIds(tagger, restaurants)
}
