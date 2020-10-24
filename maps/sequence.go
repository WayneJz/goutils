package maps

import (
	"sort"
)

// StringMapInOrder ... Extract string map keys in order, generates [[key1, value1], [key2, value2] ...]
func StringMapInOrder(m map[string]interface{}) [][]interface{} {
	res := make([][]interface{}, len(m))

	var keys []string
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for i, k := range keys {
		res[i] = []interface{}{k, m[k]}
	}
	return res
}

// Int64MapInOrder ... Extract int64 map keys in order, generates [[key1, value1], [key2, value2] ...]
func Int64MapInOrder(m map[int64]interface{}) [][]interface{} {
	res := make([][]interface{}, len(m))

	var keys []int64
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for i, k := range keys {
		res[i] = []interface{}{k, m[k]}
	}
	return res
}

// Int32MapInOrder ... Extract int32 map keys in order, generates [[key1, value1], [key2, value2] ...]
func Int32MapInOrder(m map[int32]interface{}) [][]interface{} {
	res := make([][]interface{}, len(m))

	var keys []int32
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for i, k := range keys {
		res[i] = []interface{}{k, m[k]}
	}
	return res
}

// Uint64MapInOrder ... Extract uint64 map keys in order, generates [[key1, value1], [key2, value2] ...]
func Uint64MapInOrder(m map[uint64]interface{}) [][]interface{} {
	res := make([][]interface{}, len(m))

	var keys []uint64
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for i, k := range keys {
		res[i] = []interface{}{k, m[k]}
	}
	return res
}

// Uint32MapInOrder ... Extract uint32 map keys in order, generates [[key1, value1], [key2, value2] ...]
func Uint32MapInOrder(m map[uint32]interface{}) [][]interface{} {
	res := make([][]interface{}, len(m))

	var keys []uint32
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })

	for i, k := range keys {
		res[i] = []interface{}{k, m[k]}
	}
	return res
}
