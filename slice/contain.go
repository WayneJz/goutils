package slice

// Check if all strings in string slice
func InStrings(l []string, s ...string) bool {
	checkMap := make(map[string]int8)
	for _, n := range s {
		checkMap[n] = numToCheckTag
	}
	for _, element := range l {
		if checkMap[element] == numToCheckTag {
			checkMap[element] = numFoundTag
		}
	}
	for _, n := range s {
		if checkMap[n] != numFoundTag {
			return false
		}
	}
	return true
}

// Check if all int64 num in int64 slice
func InInt64s(l []int64, nums ...int64) bool {
	checkMap := make(map[int64]int8)
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
func InInt32s(l []int32, nums ...int32) bool {
	checkMap := make(map[int32]int8)
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
func InUint64s(l []uint64, nums ...uint64) bool {
	checkMap := make(map[uint64]int8)
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
func InUint32s(l []uint32, nums ...uint32) bool {
	checkMap := make(map[uint32]int8)
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
