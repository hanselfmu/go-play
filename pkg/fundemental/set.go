package fundemental

// Set is a simple set implementation, ignoring thread safety
type Set map[Generic]struct{}

// SetValStub is a stub that represents a 0-byte value for a map-implemented set
var SetValStub = struct{}{}

// Add creates i in set and returns true; returns false if i already exists
func (set *Set) Add(i Generic) bool {
	_, exists := (*set)[i]
	if exists {
		return false //False if it existed already
	}

	(*set)[i] = SetValStub
	return true
}

// Contains returns true if i exists in set; returns false otherwise;
// "Contains" follows Java's naming convention
func (set *Set) Contains(i Generic) bool {
	_, exists := (*set)[i]
	return exists
}
