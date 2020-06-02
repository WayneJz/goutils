package slice

// IndexStrings ... Find the position of string in string slice, return -1 if not found
func IndexStrings(l []string, s string) int {
	for i, n := range l {
		if n == s {
			return i
		}
	}
	return NumNotFoundTag
}

// IndexAllStrings ... Find all positions of int64 nums in int64 slice...
// returns a map (key: given nums, value: their positions, -1 if not found)
func IndexAllStrings(l []string, s ...string) map[string]int {
	indexMap := make(map[string]int)
	for _, n := range s {
		indexMap[n] = NumNotFoundTag
	}
	for i, element := range l {
		if indexMap[element] == NumNotFoundTag {
			indexMap[element] = i
		}
	}
	return indexMap
}

// IndexInt64s ... Find the position of int64 num in int64 slice, return -1 if not found
func IndexInt64s(l []int64, num int64) int {
	for i, n := range l {
		if n == num {
			return i
		}
	}
	return NumNotFoundTag
}

// IndexAllInt64s ... Find all positions of int64 nums in int64 slice and
// returns a map (key: given nums, value: their positions, -1 if not found)
func IndexAllInt64s(l []int64, nums ...int64) map[int64]int {
	indexMap := make(map[int64]int)
	for _, num := range nums {
		indexMap[num] = NumNotFoundTag
	}
	for i, element := range l {
		if indexMap[element] == NumNotFoundTag {
			indexMap[element] = i
		}
	}
	return indexMap
}

// IndexInt32s ... Find the position of int32 num in int32 slice, return -1 if not found
func IndexInt32s(l []int32, num int32) int {
	for i, n := range l {
		if n == num {
			return i
		}
	}
	return NumNotFoundTag
}

// IndexAllInt32s ... Find all positions of int32 nums in int32 slice and
// returns a map (key: given nums, value: their positions, -1 if not found)
func IndexAllInt32s(l []int32, nums ...int32) map[int32]int {
	indexMap := make(map[int32]int)
	for _, num := range nums {
		indexMap[num] = NumNotFoundTag
	}
	for i, element := range l {
		if indexMap[element] == NumNotFoundTag {
			indexMap[element] = i
		}
	}
	return indexMap
}

// IndexUint64s ... Find the position of uint64 num in uint64 slice, return -1 if not found
func IndexUint64s(l []uint64, num uint64) int {
	for i, n := range l {
		if n == num {
			return i
		}
	}
	return NumNotFoundTag
}

// IndexAllUint64s ... Find all positions of uint64 nums in uint64 slice and
// returns a map (key: given nums, value: their positions, -1 if not found)
func IndexAllUint64s(l []uint64, nums ...uint64) map[uint64]int {
	indexMap := make(map[uint64]int)
	for _, num := range nums {
		indexMap[num] = NumNotFoundTag
	}
	for i, element := range l {
		if indexMap[element] == NumNotFoundTag {
			indexMap[element] = i
		}
	}
	return indexMap
}

// IndexUint32s ... Find the position of uint32 num in uint32 slice, return -1 if not found
func IndexUint32s(l []uint32, num uint32) int {
	for i, n := range l {
		if n == num {
			return i
		}
	}
	return NumNotFoundTag
}

// IndexAllUint32s ... Find all positions of uint32 nums in uint32 slice and
// returns a map (key: given nums, value: their positions, -1 if not found)
func IndexAllUint32s(l []uint32, nums ...uint32) map[uint32]int {
	indexMap := make(map[uint32]int)
	for _, num := range nums {
		indexMap[num] = NumNotFoundTag
	}
	for i, element := range l {
		if indexMap[element] == NumNotFoundTag {
			indexMap[element] = i
		}
	}
	return indexMap
}
