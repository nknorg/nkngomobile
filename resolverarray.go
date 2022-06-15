package nkngomobile

import (
	mathRand "math/rand"
)

type Resolver interface {
	Resolve(address string) (string, error)
}

// ResolverArray is a wrapper type for gomobile compatibility. ResolverArray is not
// protected by lock and should not be read and write at the same time.
type ResolverArray struct{ elems []Resolver }

// NewResolverArray creates a ResolverArray from a list of string elements.
func NewResolverArray(elems ...Resolver) *ResolverArray {
	return &ResolverArray{elems}
}

// NewResolverArrayFromResolver creates a ResolverArray from a single string input.
// The input string will be split to string array by whitespace.
func NewResolverArrayFromResolver(e Resolver) *ResolverArray {
	return &ResolverArray{[]Resolver{e}}
}

// Elems returns the string array elements.
func (ia *ResolverArray) Elems() []Resolver {
	if ia == nil {
		return nil
	}
	return ia.elems
}

// Len returns the string array length.
func (ia *ResolverArray) Len() int {
	return len(ia.Elems())
}

// Append adds an element to the string array.
func (ia *ResolverArray) Append(a Resolver) {
	ia.elems = append(ia.elems, a)
}

// Get gets an element to the string array.
func (ia *ResolverArray) Get(i int) Resolver {
	return ia.Elems()[i]
}

// RandomElem returns a randome element from the string array. The random number
// is generated using math/rand and thus not cryptographically secure.
func (ia *ResolverArray) RandomElem() Resolver {
	if ia.Len() == 0 {
		return nil
	}
	return ia.Elems()[mathRand.Intn(ia.Len())]
}
