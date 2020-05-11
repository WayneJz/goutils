package slice

import "sort"

// Stable sort int64 slice
func SortInt64Slice(l []int64) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] < l[j] })
}

// Stable reverse sort int64 slice
func RSortInt64Slice(l []int64) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] > l[j] })
}

// Stable sort int32 slice
func SortInt32Slice(l []int32) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] < l[j] })
}

// Stable reverse sort int32 slice
func RSortInt32Slice(l []int32) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] > l[j] })
}

// Stable sort uint64 slice
func SortUint64Slice(l []uint64) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] < l[j] })
}

// Stable reverse sort uint64 slice
func RSortUint64Slice(l []uint64) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] > l[j] })
}

// Stable sort uint32 slice
func SortUint32Slice(l []uint32) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] < l[j] })
}

// Stable reverse sort uint32 slice
func RSortUint32Slice(l []uint32) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] > l[j] })
}
