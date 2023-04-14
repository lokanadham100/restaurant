package models

type PriorityQueueItem interface {
	GetPriority() int
	GetValue() int
}

type DupRemover interface {
	MergeWithOutDuplicates([]string, []string, int) []string
}
