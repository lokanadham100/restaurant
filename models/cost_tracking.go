package models

type CostTracking struct {
	typ        int
	noOfOrders int
}

func (ct CostTracking) GetPriority() int {
	return ct.noOfOrders
}

func (ct CostTracking) GetValue() int {
	return ct.typ
}
