package models

type User struct {
	cuisines     []CuisineTracking
	costBrackets []CostTracking
}

func (u User) GetCuisinesForPQ() []PriorityQueueItem {
	res := make([]PriorityQueueItem, len(u.cuisines))
	for i, cuisine := range u.cuisines {
		res[i] = cuisine
	}
	return res
}

func (u User) GetCostBracketForPQ() []PriorityQueueItem {
	res := make([]PriorityQueueItem, len(u.costBrackets))
	for i, costBracket := range u.costBrackets {
		res[i] = costBracket
	}
	return res
}
