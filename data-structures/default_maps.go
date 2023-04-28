package data_structures

type mapper map[int]string

type mapKeyValues struct {
	m mapper
}

func NewMapKeyValues(capacity int) *mapKeyValues {
	return &mapKeyValues{m: make(mapper, capacity)}
}

// Get retrieves the value associated with the given key.
// Returns an empty string if the key is not found.
func (m *mapKeyValues) Get(key int) string {
	return m.m[key]
}

// Set associates the given value with the given key.
// If the key already exists in the map, its value is updated.
func (m *mapKeyValues) Set(key int, value string) {
	m.m[key] = value
}
