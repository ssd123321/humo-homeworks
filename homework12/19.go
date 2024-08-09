package main

type KeyValuePair struct {
	Key   string
	Value int
}

func Reverse(m map[string]int) []KeyValuePair {
	k := make([]KeyValuePair, 0)
	for key, value := range m {
		k = append(k, KeyValuePair{Key: key, Value: value})
	}
	return k
}
