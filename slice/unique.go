package slice

// Return a unique copy of given string slice
func UniqueStrings(l []string) []string {
	checkMap := make(map[string]int8)
	var res []string
	for _, n := range l {
		if checkMap[n] == numFoundTag {
			continue
		}
		checkMap[n] = numFoundTag
		res = append(res, n)
	}
	return res
}

// Return a unique copy of given int64 slice
func UniqueInt64s(l []int64) []int64 {
	checkMap := make(map[int64]int8)
	var res []int64
	for _, n := range l {
		if checkMap[n] == numFoundTag {
			continue
		}
		checkMap[n] = numFoundTag
		res = append(res, n)
	}
	return res
}

// Return a unique copy of given int32 slice
func UniqueInt32s(l []int32) []int32 {
	checkMap := make(map[int32]int8)
	var res []int32
	for _, n := range l {
		if checkMap[n] == numFoundTag {
			continue
		}
		checkMap[n] = numFoundTag
		res = append(res, n)
	}
	return res
}

// Return a unique copy of given uint64 slice
func UniqueUint64s(l []uint64) []uint64 {
	checkMap := make(map[uint64]int8)
	var res []uint64
	for _, n := range l {
		if checkMap[n] == numFoundTag {
			continue
		}
		checkMap[n] = numFoundTag
		res = append(res, n)
	}
	return res
}

// Return a unique copy of given uint32 slice
func UniqueUint32s(l []uint32) []uint32 {
	checkMap := make(map[uint32]int8)
	var res []uint32
	for _, n := range l {
		if checkMap[n] == numFoundTag {
			continue
		}
		checkMap[n] = numFoundTag
		res = append(res, n)
	}
	return res
}
