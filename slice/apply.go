package slice

/*
	Apply a function to slice
	Sample usage: multiply each element by 2

	l := []int64{0, 7, 8}
	slice.ApplyToInt64Slice(l, func(i int64) int64 { return i * 2 })
	fmt.Println(l)

	Expected out: [0, 14, 16]
*/

// ApplyStrings ... Apply a function to each element in string slice
func ApplyStrings(l []string, applyFunc func(n string) string) {
	for i, s := range l {
		l[i] = applyFunc(s)
	}
}

// ApplyInt64s ... Apply a function to each element in int64 slice
func ApplyInt64s(l []int64, applyFunc func(n int64) int64) {
	for i, num := range l {
		l[i] = applyFunc(num)
	}
}

// ApplyInt32s ... Apply a function to each element in int32 slice
func ApplyInt32s(l []int32, applyFunc func(n int32) int32) {
	for i, num := range l {
		l[i] = applyFunc(num)
	}
}

// ApplyUint64s ... Apply a function to each element in uint64 slice
func ApplyUint64s(l []uint64, applyFunc func(n uint64) uint64) {
	for i, num := range l {
		l[i] = applyFunc(num)
	}
}

// ApplyUint32s ... Apply a function to each element in uint32 slice
func ApplyUint32s(l []uint32, applyFunc func(n uint32) uint32) {
	for i, num := range l {
		l[i] = applyFunc(num)
	}
}

/*
	Apply a index function to slice
	Sample usage: multiply each element by 2 if it is in odd-value position

	l := []int64{0, 7, 8}
	slice.ApplyIdxToInt64Slice(l, func(i int, num int64) int64 {
		if i%2 == 0 {
			return num * 2
		}
		return num
	})
	fmt.Println(l)

	Expected out: [0, 7, 16]
*/

// IApplyStrings ... Apply an index-related function to each element in string slice
func IApplyStrings(l []string, applyFunc func(i int, n string) string) {
	for i, s := range l {
		l[i] = applyFunc(i, s)
	}
}

// IApplyInt64s ... Apply an index-related function to each element in int64 slice
func IApplyInt64s(l []int64, applyFunc func(i int, n int64) int64) {
	for i, num := range l {
		l[i] = applyFunc(i, num)
	}
}

// IApplyInt32s ... Apply an index-related function to each element in int32 slice
func IApplyInt32s(l []int32, applyFunc func(i int, n int32) int32) {
	for i, num := range l {
		l[i] = applyFunc(i, num)
	}
}

// IApplyUint64s ... Apply an index-related function to each element in uint64 slice
func IApplyUint64s(l []uint64, applyFunc func(i int, n uint64) uint64) {
	for i, num := range l {
		l[i] = applyFunc(i, num)
	}
}

// IApplyUint32s ... Apply an index-related function to each element in uint32 slice
func IApplyUint32s(l []uint32, applyFunc func(i int, n uint32) uint32) {
	for i, num := range l {
		l[i] = applyFunc(i, num)
	}
}
