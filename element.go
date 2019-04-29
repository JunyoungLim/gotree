package gotree

// Element is an abstraction of the Tree Node Element type, which supports
// comparison, including equality and precedence.
type Element interface {
	// Checks element equality and returns true if this element equals to the other
	Equals(Element) bool
	// Checks element precedence and returns true if element is less than the other
	Less(Element) bool
	// Returns a string representation of this element
	String() string
}
