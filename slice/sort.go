package slice

import "sort"

// Stable sort string slice by length of each string
func SortStringsByLen(l []string) {
	sort.SliceStable(l, func(i int, j int) bool { return len(l[i]) < len(l[j]) })
}

// Stable sort string slice by length of each string
func RSortStringsByLen(l []string) {
	sort.SliceStable(l, func(i int, j int) bool { return len(l[i]) > len(l[j]) })
}

// Stable sort int64 slice
func SortInt64s(l []int64) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] < l[j] })
}

// Stable reverse sort int64 slice
func RSortInt64s(l []int64) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] > l[j] })
}

// Stable sort int32 slice
func SortInt32s(l []int32) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] < l[j] })
}

// Stable reverse sort int32 slice
func RSortInt32s(l []int32) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] > l[j] })
}

// Stable sort uint64 slice
func SortUint64s(l []uint64) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] < l[j] })
}

// Stable reverse sort uint64 slice
func RSortUint64s(l []uint64) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] > l[j] })
}

// Stable sort uint32 slice
func SortUint32s(l []uint32) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] < l[j] })
}

// Stable reverse sort uint32 slice
func RSortUint32s(l []uint32) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] > l[j] })
}
