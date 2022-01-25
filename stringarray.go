package nkngomobile

import (
	mathRand "math/rand"
	"strings"
)

// StringArray is a wrapper type for gomobile compatibility. StringArray is not
// protected by lock and should not be read and write at the same time.
type StringArray struct{ elems []string }

// NewStringArray creates a StringArray from a list of string elements.
func NewStringArray(elems ...string) *StringArray {
	return &StringArray{elems}
}

// NewStringArrayFromString creates a StringArray from a single string input.
// The input string will be split to string array by whitespace.
func NewStringArrayFromString(s string) *StringArray {
	return &StringArray{strings.Fields(s)}
}

// Elems returns the string array elements.
func (sa *StringArray) Elems() []string {
	if sa == nil {
		return nil
	}
	return sa.elems
}

// Len returns the string array length.
func (sa *StringArray) Len() int {
	return len(sa.Elems())
}

// Append adds an element to the string array.
func (sa *StringArray) Append(s string) {
	sa.elems = append(sa.elems, s)
}

// Get gets an element to the string array.
func (sa *StringArray) Get(i int) string {
	return sa.Elems()[i]
}

// RandomElem returns a randome element from the string array. The random number
// is generated using math/rand and thus not cryptographically secure.
func (sa *StringArray) RandomElem() string {
	if sa.Len() == 0 {
		return ""
	}
	return sa.Elems()[mathRand.Intn(sa.Len())]
}

// Join returns a single string by concatenates the elements
func (sa *StringArray) Join(separator string) string {
	if sa == nil {
		return ""
	}
	return strings.Join(sa.elems, separator)
}
