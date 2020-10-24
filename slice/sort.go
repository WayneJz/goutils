package slice

import "sort"

// SortStringsByLen ... Stable sort string slice by length of each string
func SortStringsByLen(l []string) {
	sort.SliceStable(l, func(i int, j int) bool { return len(l[i]) < len(l[j]) })
}

// RSortStringsByLen ... Reverse stable sort string slice by length of each string
func RSortStringsByLen(l []string) {
	sort.SliceStable(l, func(i int, j int) bool { return len(l[i]) > len(l[j]) })
}

// SortBytesByLen ... Stable sort bytes slice by length of each bytes
func SortBytesByLen(l [][]byte) {
	sort.SliceStable(l, func(i int, j int) bool { return len(l[i]) < len(l[j]) })
}

// RSortBytesByLen ... Reverse stable sort bytes slice by length of each bytes
func RSortBytesByLen(l [][]byte) {
	sort.SliceStable(l, func(i int, j int) bool { return len(l[i]) > len(l[j]) })
}

// SortInt64s ... Stable sort int64 slice
func SortInt64s(l []int64) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] < l[j] })
}

// RSortInt64s ... Reverse stable sort int64 slice
func RSortInt64s(l []int64) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] > l[j] })
}

// SortInt32s ... Stable sort int32 slice
func SortInt32s(l []int32) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] < l[j] })
}

// RSortInt32s ... Reverse stable sort int32 slice
func RSortInt32s(l []int32) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] > l[j] })
}

// SortUint64s ... Stable sort uint64 slice
func SortUint64s(l []uint64) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] < l[j] })
}

// RSortUint64s ... Reverse stable sort uint64 slice
func RSortUint64s(l []uint64) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] > l[j] })
}

// SortUint32s ... Stable sort uint32 slice
func SortUint32s(l []uint32) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] < l[j] })
}

// RSortUint32s ... Reverse stable sort uint32 slice
func RSortUint32s(l []uint32) {
	sort.SliceStable(l, func(i int, j int) bool { return l[i] > l[j] })
}
