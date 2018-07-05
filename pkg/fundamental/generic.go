package fundamental

// Generic represents a generic type
type Generic interface{}

// Comparable represents a comparable type
type Comparable interface {
	LessThan(other Comparable) bool
}

// ComparableInt represents an integer that implements Comparable interface
type ComparableInt int

// LessThan compares if a ComparableInt is less than the other ComparableInt
func (ci ComparableInt) LessThan(other Comparable) bool {
	return ci < other.(ComparableInt)
}
