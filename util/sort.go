package util

type Order int

const (
	OrderAsc Order = iota
	OrderDesc
)

type OrderBy struct {
	Order Order
	value string
}

var (
	Orders = []OrderBy{
		{Order: OrderAsc, value: " ASC"},
		{Order: OrderDesc, value: " DESC"},
	}
)

func (c OrderBy) String() string {
	return c.value
}

type Sort struct {
	Field     string
	Direction Order // "asc" or "desc"
}
