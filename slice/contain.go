package slice

const (
	numToCheckTag = 1
	numFoundTag   = 2
)

// Check if all int64 num in int64 slice
func InInt64Slice(l []int64, nums ...int64) bool {
	checkMap := make(map[int64]uint8)
	for _, num := range nums {
		checkMap[num] = numToCheckTag
	}
	for _, element := range l {
		if checkMap[element] == numToCheckTag {
			checkMap[element] = numFoundTag
		}
	}
	for _, num := range nums {
		if checkMap[num] != numFoundTag {
			return false
		}
	}
	return true
}

// Check if all int32 num in int32 slice
func InInt32Slice(l []int32, nums ...int32) bool {
	checkMap := make(map[int32]uint8)
	for _, num := range nums {
		checkMap[num] = numToCheckTag
	}
	for _, element := range l {
		if checkMap[element] == numToCheckTag {
			checkMap[element] = numFoundTag
		}
	}
	for _, num := range nums {
		if checkMap[num] != numFoundTag {
			return false
		}
	}
	return true
}

// Check if all uint64 num in uint64 slice
func InUint64Slice(l []uint64, nums ...uint64) bool {
	checkMap := make(map[uint64]uint8)
	for _, num := range nums {
		checkMap[num] = numToCheckTag
	}
	for _, element := range l {
		if checkMap[element] == numToCheckTag {
			checkMap[element] = numFoundTag
		}
	}
	for _, num := range nums {
		if checkMap[num] != numFoundTag {
			return false
		}
	}
	return true
}

// Check if all uint32 num in uint32 slice
func InUint32Slice(l []uint32, nums ...uint32) bool {
	checkMap := make(map[uint32]uint8)
	for _, num := range nums {
		checkMap[num] = numToCheckTag
	}
	for _, element := range l {
		if checkMap[element] == numToCheckTag {
			checkMap[element] = numFoundTag
		}
	}
	for _, num := range nums {
		if checkMap[num] != numFoundTag {
			return false
		}
	}
	return true
}
