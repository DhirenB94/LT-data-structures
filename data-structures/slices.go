package data_structures

type keyValue struct {
	key   int
	value string
}

type SliceKeyValues struct {
	KeyValues []keyValue
}

func NewSliceKeyValues(length int) *SliceKeyValues {
	return &SliceKeyValues{KeyValues: make([]keyValue, length)}
}

func (s *SliceKeyValues) Get(key int) string {
	for _, k := range s.KeyValues {
		if k.key == key {
			return k.value
		}
	}
	return "element not found"
}

func (s *SliceKeyValues) Set(key int, value string) string {
	keyIndex := s.indexOf(key)
	if keyIndex == -1 {
		s.KeyValues = append(s.KeyValues, keyValue{
			key:   key,
			value: value,
		})
		return value
	}
	s.KeyValues[keyIndex].value = value
	return value
}

func (s *SliceKeyValues) indexOf(key int) int {
	for i, k := range s.KeyValues {
		if k.key == key {
			return i
		}
	}
	return -1
}
