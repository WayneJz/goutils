package slice

// InStrings ... Check if all strings in string slice
func InStrings(l []string, s ...string) bool {
	checkMap := make(map[string]int8)
	for _, n := range s {
		checkMap[n] = ToCheckTag
	}
	for _, element := range l {
		if checkMap[element] == ToCheckTag {
			checkMap[element] = FoundTag
		}
	}
	for _, n := range s {
		if checkMap[n] != FoundTag {
			return false
		}
	}
	return true
}

// OneInStrings ... Check if at least one string in string slice
func OneInStrings(l []string, s ...string) bool {
	checkMap := make(map[string]int8)
	for _, n := range s {
		checkMap[n] = ToCheckTag
	}
	for _, element := range l {
		if _, ok := checkMap[element]; ok {
			return true
		}
	}
	return false
}

// InInt64s ... Check if all int64 num in int64 slice
func InInt64s(l []int64, nums ...int64) bool {
	checkMap := make(map[int64]int8)
	for _, num := range nums {
		checkMap[num] = ToCheckTag
	}
	for _, element := range l {
		if checkMap[element] == ToCheckTag {
			checkMap[element] = FoundTag
		}
	}
	for _, num := range nums {
		if checkMap[num] != FoundTag {
			return false
		}
	}
	return true
}

// OneInInt64s ... Check if at least one int64 num in int64 slice
func OneInInt64s(l []int64, nums ...int64) bool {
	checkMap := make(map[int64]int8)
	for _, n := range nums {
		checkMap[n] = ToCheckTag
	}
	for _, element := range l {
		if _, ok := checkMap[element]; ok {
			return true
		}
	}
	return false
}

// InInt32s ... Check if all int32 num in int32 slice
func InInt32s(l []int32, nums ...int32) bool {
	checkMap := make(map[int32]int8)
	for _, num := range nums {
		checkMap[num] = ToCheckTag
	}
	for _, element := range l {
		if checkMap[element] == ToCheckTag {
			checkMap[element] = FoundTag
		}
	}
	for _, num := range nums {
		if checkMap[num] != FoundTag {
			return false
		}
	}
	return true
}

// OneInInt32s ... Check if at least one int32 num in int32 slice
func OneInInt32s(l []int32, nums ...int32) bool {
	checkMap := make(map[int32]int8)
	for _, n := range nums {
		checkMap[n] = ToCheckTag
	}
	for _, element := range l {
		if _, ok := checkMap[element]; ok {
			return true
		}
	}
	return false
}

// InUint64s ... Check if all uint64 num in uint64 slice
func InUint64s(l []uint64, nums ...uint64) bool {
	checkMap := make(map[uint64]int8)
	for _, num := range nums {
		checkMap[num] = ToCheckTag
	}
	for _, element := range l {
		if checkMap[element] == ToCheckTag {
			checkMap[element] = FoundTag
		}
	}
	for _, num := range nums {
		if checkMap[num] != FoundTag {
			return false
		}
	}
	return true
}

// OneInUint64s ... Check if at least one uint64 num in uint64 slice
func OneInUint64s(l []uint64, nums ...uint64) bool {
	checkMap := make(map[uint64]int8)
	for _, n := range nums {
		checkMap[n] = ToCheckTag
	}
	for _, element := range l {
		if _, ok := checkMap[element]; ok {
			return true
		}
	}
	return false
}

// InUint32s ... Check if all uint32 num in uint32 slice
func InUint32s(l []uint32, nums ...uint32) bool {
	checkMap := make(map[uint32]int8)
	for _, num := range nums {
		checkMap[num] = ToCheckTag
	}
	for _, element := range l {
		if checkMap[element] == ToCheckTag {
			checkMap[element] = FoundTag
		}
	}
	for _, num := range nums {
		if checkMap[num] != FoundTag {
			return false
		}
	}
	return true
}

// OneInUint32s ... Check if at least one uint32 num in uint32 slice
func OneInUint32s(l []uint32, nums ...uint32) bool {
	checkMap := make(map[uint32]int8)
	for _, n := range nums {
		checkMap[n] = ToCheckTag
	}
	for _, element := range l {
		if _, ok := checkMap[element]; ok {
			return true
		}
	}
	return false
}
