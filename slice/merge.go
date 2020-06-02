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

// UniqueMergeStrings ... Merge string slices into one string slice, excluding duplicated strings
func UniqueMergeStrings(l ...[]string) []string {
	var res []string
	checkMap := make(map[string]int8)
	for _, s := range l {
		for _, str := range s {
			if checkMap[str] == numFoundTag {
				continue
			}
			checkMap[str] = numFoundTag
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

// Merge int64 slices into one int64 slice, excluding duplicated int64
func UniqueMergeInt64s(l ...[]int64) []int64 {
	var res []int64
	checkMap := make(map[int64]int8)
	for _, n := range l {
		for _, num := range n {
			if checkMap[num] == numFoundTag {
				continue
			}
			checkMap[num] = numFoundTag
			res = append(res, num)
		}
	}
	return res
}

// Merge int32 slices into one int32 slice
func MergeInt32s(l ...[]int32) []int32 {
	var res []int32
	for _, n := range l {
		for _, num := range n {
			res = append(res, num)
		}
	}
	return res
}

// Merge int32 slices into one int32 slice, excluding duplicated int32
func UniqueMergeInt32s(l ...[]int32) []int32 {
	var res []int32
	checkMap := make(map[int32]int8)
	for _, n := range l {
		for _, num := range n {
			if checkMap[num] == numFoundTag {
				continue
			}
			checkMap[num] = numFoundTag
			res = append(res, num)
		}
	}
	return res
}

// Merge uint64 slices into one uint64 slice
func MergeUint64s(l ...[]uint64) []uint64 {
	var res []uint64
	for _, n := range l {
		for _, num := range n {
			res = append(res, num)
		}
	}
	return res
}

// Merge uint64 slices into one uint64 slice, excluding duplicated uint64
func UniqueMergeUint64s(l ...[]uint64) []uint64 {
	var res []uint64
	checkMap := make(map[uint64]int8)
	for _, n := range l {
		for _, num := range n {
			if checkMap[num] == numFoundTag {
				continue
			}
			checkMap[num] = numFoundTag
			res = append(res, num)
		}
	}
	return res
}

// Merge uint32 slices into one uint32 slice
func MergeUint32s(l ...[]uint32) []uint32 {
	var res []uint32
	for _, n := range l {
		for _, num := range n {
			res = append(res, num)
		}
	}
	return res
}

// Merge uint32 slices into one uint32 slice, excluding duplicated uint32
func UniqueMergeUint32s(l ...[]uint32) []uint32 {
	var res []uint32
	checkMap := make(map[uint32]int8)
	for _, n := range l {
		for _, num := range n {
			if checkMap[num] == numFoundTag {
				continue
			}
			checkMap[num] = numFoundTag
			res = append(res, num)
		}
	}
	return res
}
