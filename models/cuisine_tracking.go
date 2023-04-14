package models

type CuisineTracking struct {
	typ        Cuisine
	noOfOrders int
}

func (ct CuisineTracking) GetPriority() int {
	return ct.noOfOrders
}

func (ct CuisineTracking) GetValue() int {
	return int(ct.typ)
}
