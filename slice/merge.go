package slice

// MergeStrings ... Merge string slices into one string slice
func MergeStrings(l ...[]string) []string {
	var res []string
	for _, s := range l {
		for _, str := range s {
			res = append(res, str)
		}
	}
	return res
}

// UniqueMergeStrings ... Merge string slices into one string slice, excluding duplicated elements
func UniqueMergeStrings(l ...[]string) []string {
	var res []string
	checkMap := make(map[string]int8)
	for _, s := range l {
		for _, str := range s {
			if checkMap[str] == FoundTag {
				continue
			}
			checkMap[str] = FoundTag
			res = append(res, str)
		}
	}
	return res
}

// MergeInt64s ... Merge int64 slices into one int64 slice
func MergeInt64s(l ...[]int64) []int64 {
	var res []int64
	for _, n := range l {
		for _, num := range n {
			res = append(res, num)
		}
	}
	return res
}

// UniqueMergeInt64s ... Merge int64 slices into one int64 slice, excluding duplicated elements
func UniqueMergeInt64s(l ...[]int64) []int64 {
	var res []int64
	checkMap := make(map[int64]int8)
	for _, n := range l {
		for _, num := range n {
			if checkMap[num] == FoundTag {
				continue
			}
			checkMap[num] = FoundTag
			res = append(res, num)
		}
	}
	return res
}

// MergeInt32s ... Merge int32 slices into one int32 slice
func MergeInt32s(l ...[]int32) []int32 {
	var res []int32
	for _, n := range l {
		for _, num := range n {
			res = append(res, num)
		}
	}
	return res
}

// UniqueMergeInt32s ... Merge int32 slices into one int32 slice, excluding duplicated elements
func UniqueMergeInt32s(l ...[]int32) []int32 {
	var res []int32
	checkMap := make(map[int32]int8)
	for _, n := range l {
		for _, num := range n {
			if checkMap[num] == FoundTag {
				continue
			}
			checkMap[num] = FoundTag
			res = append(res, num)
		}
	}
	return res
}

// MergeUint64s ... Merge uint64 slices into one uint64 slice
func MergeUint64s(l ...[]uint64) []uint64 {
	var res []uint64
	for _, n := range l {
		for _, num := range n {
			res = append(res, num)
		}
	}
	return res
}

// UniqueMergeUint64s ... Merge uint64 slices into one uint64 slice, excluding duplicated elements
func UniqueMergeUint64s(l ...[]uint64) []uint64 {
	var res []uint64
	checkMap := make(map[uint64]int8)
	for _, n := range l {
		for _, num := range n {
			if checkMap[num] == FoundTag {
				continue
			}
			checkMap[num] = FoundTag
			res = append(res, num)
		}
	}
	return res
}

// MergeUint32s ... Merge uint32 slices into one uint32 slice
func MergeUint32s(l ...[]uint32) []uint32 {
	var res []uint32
	for _, n := range l {
		for _, num := range n {
			res = append(res, num)
		}
	}
	return res
}

// UniqueMergeUint32s ... Merge uint32 slices into one uint32 slice, excluding duplicated elements
func UniqueMergeUint32s(l ...[]uint32) []uint32 {
	var res []uint32
	checkMap := make(map[uint32]int8)
	for _, n := range l {
		for _, num := range n {
			if checkMap[num] == FoundTag {
				continue
			}
			checkMap[num] = FoundTag
			res = append(res, num)
		}
	}
	return res
}
