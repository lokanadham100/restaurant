package utils

import "github.com/lokanadham100/restaurant/models"

func NewSorter(items []models.PriorityQueueItem, numOfElementsToSort int) *Sorter {
	if numOfElementsToSort < len(items) {
		numOfElementsToSort = len(items)
	}
	sorter := &Sorter{
		items:               items,
		numOfElementsToSort: numOfElementsToSort,
	}
	sorter.sortRequiredElements()
	return sorter
}

type Sorter struct {
	items               []models.PriorityQueueItem
	numOfElementsToSort int
}

func (s *Sorter) sortRequiredElements() {
	for i := 0; i < int(s.numOfElementsToSort); i++ {
		maxIndex := i
		for j := i + 1; j < len(s.items); j++ {
			if s.items[j].GetPriority() > s.items[maxIndex].GetPriority() {
				maxIndex = j
			}
		}
		s.items[i], s.items[maxIndex] = s.items[maxIndex], s.items[i]
	}
}

func (s *Sorter) Items() []models.PriorityQueueItem {
	return s.items
}

// func (s *Sorter) Len() int {
// 	return len(s.items)
// }

// func (s *Sorter) Less(i, j int) bool {
// 	return s.items[i].GetPriority() > s.items[j].GetPriority()
// }

// func (s *Sorter) Swap(i, j int) {
// 	s.items[i], s.items[j] = s.items[j], s.items[i]
// }
