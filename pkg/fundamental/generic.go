package fundamental

// Generic represents a generic type
type Generic interface{}

// Comparable represents a comparable type
type Comparable interface {
	LessThan(other Comparable) bool
}
