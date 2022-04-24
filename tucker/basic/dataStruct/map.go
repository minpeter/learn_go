package dataStruct

func Hash(s string) int {
	h := 0
	A := 256  // size of alphabet (ASCII)
	B := 3571 // 적당히 큰 소수 값
	for i := 0; i < len(s); i++ {
		h = (h*A + int(s[i])) % B
	}
	return h
}

type KeyValue struct {
	key   string
	value string
}

type Map struct {
	keyArray [3571][]KeyValue
}

func CreatMap() *Map {
	return &Map{}
}

func (m *Map) Add(key, value string) {
	h := Hash(key)
	m.keyArray[h] = append(m.keyArray[h], KeyValue{key, value})
}

func (m *Map) Get(key string) string {
	h := Hash(key)
	for _, kv := range m.keyArray[h] {
		if kv.key == key {
			return kv.value
		}
	}
	return ""
}
