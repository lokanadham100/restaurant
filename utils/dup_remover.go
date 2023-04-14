package utils

import (
	"github.com/lokanadham100/restaurant/models"
)

type dupRemover struct {
	uniqMap map[string]struct{}
}

func (dr *dupRemover) MergeWithOutDuplicates(primary []string, secondary []string, count int) []string {
	if count == -1 || count > len(secondary) {
		count = len(secondary)
	}

	for i := 0; i < count; i++ {
		element := secondary[i]
		if _, ok := dr.uniqMap[element]; !ok {
			primary = append(primary, element)
			dr.uniqMap[element] = struct{}{}
		}
	}

	return primary
}

func NewDupRemover() models.DupRemover {
	return &dupRemover{
		uniqMap: map[string]struct{}{},
	}
}
