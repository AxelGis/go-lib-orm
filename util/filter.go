package util

type Op int

const (
	OpEqual Op = iota
	OpLike
	OpIn
	OpGreaterThan
	OpGreaterThanOrEqual
	OpLessThan
	OpLessThanOrEqual
)

type Condition struct {
	Op    Op
	value string
}

var (
	Conditions = []Condition{
		{Op: OpEqual, value: "? = ?"},
		{Op: OpLike, value: "? LIKE ?"},
		{Op: OpIn, value: "? IN (?)"},
		{Op: OpGreaterThan, value: "? > ?"},
		{Op: OpGreaterThanOrEqual, value: "? >= ?"},
		{Op: OpLessThan, value: "? < ?"},
		{Op: OpLessThanOrEqual, value: "? <= ?"},
	}
)

func (c Condition) String() string {
	return c.value
}

type Filter struct {
	Field       string
	Value       interface{}
	Op          Op     // "=", "LIKE", "IN", ">", ">=", "<", "<="
	CustomQuery string // Custom SQL query
}
