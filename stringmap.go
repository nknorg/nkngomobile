package nkngomobile

import "errors"

// StringMapFunc is a wrapper type for gomobile compatibility.
type StringMapFunc interface{ OnVisit(string, string) bool }

// StringMap is a wrapper type for gomobile compatibility. StringMap is not
// protected by lock and should not be read and write at the same time.
type StringMap struct{ Map map[string]string }

// NewStringMap creates a StringMap from a map.
func NewStringMap(m map[string]string) *StringMap {
	return &StringMap{m}
}

// NewStringMapWithSize creates an empty StringMap with a given size.
func NewStringMapWithSize(size int) *StringMap {
	return &StringMap{make(map[string]string, size)}
}

// Get returns the value of a key, or ErrKeyNotInMap if key does not exist.
func (sm *StringMap) Get(key string) (string, error) {
	if value, ok := sm.Map[key]; ok {
		return value, nil
	}
	return "", errors.New("key not in map")
}

// Set sets the value of a key to a value.
func (sm *StringMap) Set(key, value string) {
	sm.Map[key] = value
}

// Delete deletes a key and its value from the map.
func (sm *StringMap) Delete(key string) {
	delete(sm.Map, key)
}

// Len returns the number of elements in the map.
func (sm *StringMap) Len() int {
	return len(sm.Map)
}

// Range iterates over the StringMap and call the OnVisit callback function with
// each element in the map. If the OnVisit function returns false, the iterator
// will stop and no longer visit the rest elements.
func (sm *StringMap) Range(cb StringMapFunc) {
	if cb != nil {
		for key, value := range sm.Map {
			if !cb.OnVisit(key, value) {
				return
			}
		}
	}
}
