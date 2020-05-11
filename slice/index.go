package slice

const (
	NumNotFoundTag = -1
)

// Find the position of int64 num in int64 slice, return -1 if not found
func IndexInt64Slice(l []int64, num int64) int {
	for i, n := range l {
		if n == num {
			return i
		}
	}
	return NumNotFoundTag
}

// Find all positions of int64 nums in int64 slice...
// returns a map (key: given nums, value: their positions, -1 if not found)
func IndexAllInt64Slice(l []int64, nums ...int64) map[int64]int {
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

// Find the position of int32 num in int32 slice, return -1 if not found
func IndexInt32Slice(l []int32, num int32) int {
	for i, n := range l {
		if n == num {
			return i
		}
	}
	return NumNotFoundTag
}

// Find all positions of int32 nums in int32 slice...
// returns a map (key: given nums, value: their positions, -1 if not found)
func IndexAllInt32Slice(l []int32, nums ...int32) map[int32]int {
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

// Find the position of uint64 num in uint64 slice, return -1 if not found
func IndexUint64Slice(l []uint64, num uint64) int {
	for i, n := range l {
		if n == num {
			return i
		}
	}
	return NumNotFoundTag
}

// Find all positions of uint64 nums in uint64 slice...
// returns a map (key: given nums, value: their positions, -1 if not found)
func IndexAllUint64Slice(l []uint64, nums ...uint64) map[uint64]int {
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

// Find the position of uint32 num in uint32 slice, return -1 if not found
func IndexUint32Slice(l []uint32, num uint32) int {
	for i, n := range l {
		if n == num {
			return i
		}
	}
	return NumNotFoundTag
}

// Find all positions of uint32 nums in uint32 slice...
// returns a map (key: given nums, value: their positions, -1 if not found)
func IndexAllUint32Slice(l []uint32, nums ...uint32) map[uint32]int {
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
