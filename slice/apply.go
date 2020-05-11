package slice

/*
	Apply a function to slice
	Sample usage: multiply each element by 2

	l := []int64{0, 7, 8}
	slice.ApplyToInt64Slice(l, func(i int64) int64 { return i * 2 })
	fmt.Println(l)

	Expected out: [0, 14, 16]
*/

// Apply a function to each element in int64 slice
func ApplyToInt64Slice(l []int64, applyFunc func(i int64) int64) {
	for i, num := range l {
		l[i] = applyFunc(num)
	}
}

// Apply a function to each element in int32 slice
func ApplyToInt32Slice(l []int32, applyFunc func(i int32) int32) {
	for i, num := range l {
		l[i] = applyFunc(num)
	}
}

// Apply a function to each element in uint64 slice
func ApplyToUint64Slice(l []uint64, applyFunc func(i uint64) uint64) {
	for i, num := range l {
		l[i] = applyFunc(num)
	}
}

// Apply a function to each element in uint32 slice
func ApplyToUint32Slice(l []uint32, applyFunc func(i uint32) uint32) {
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

// Apply a index function to each element in int64 slice
func ApplyIdxToInt64Slice(l []int64, applyFunc func(i int, num int64) int64) {
	for i, num := range l {
		l[i] = applyFunc(i, num)
	}
}

// Apply a index function to each element in int32 slice
func ApplyIdxToInt32Slice(l []int32, applyFunc func(i int, num int32) int32) {
	for i, num := range l {
		l[i] = applyFunc(i, num)
	}
}

// Apply a index function to each element in uint64 slice
func ApplyIdxToUint64Slice(l []uint64, applyFunc func(i int, num uint64) uint64) {
	for i, num := range l {
		l[i] = applyFunc(i, num)
	}
}

// Apply a index function to each element in uint32 slice
func ApplyIdxToUint32Slice(l []uint32, applyFunc func(i int, num uint32) uint32) {
	for i, num := range l {
		l[i] = applyFunc(i, num)
	}
}
