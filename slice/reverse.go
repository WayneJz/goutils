package slice

// ReverseStrings ... Reverse string slice in-place
func ReverseStrings(l []string) {
	length := len(l)
	for i := 0; i < length/2; i++ {
		l[i], l[length-i-1] = l[length-i-1], l[i]
	}
}

// ReverseInt64s ... Reverse int64 slice in-place
func ReverseInt64s(l []int64) {
	length := len(l)
	for i := 0; i < length/2; i++ {
		l[i], l[length-i-1] = l[length-i-1], l[i]
	}
}

// ReverseInt32s ... Reverse int32 slice in-place
func ReverseInt32s(l []int32) {
	length := len(l)
	for i := 0; i < length/2; i++ {
		l[i], l[length-i-1] = l[length-i-1], l[i]
	}
}

// ReverseUint64s ... Reverse uint64 slice in-place
func ReverseUint64s(l []uint64) {
	length := len(l)
	for i := 0; i < length/2; i++ {
		l[i], l[length-i-1] = l[length-i-1], l[i]
	}
}

// ReverseUint32s ... Reverse uint32 slice in-place
func ReverseUint32s(l []uint32) {
	length := len(l)
	for i := 0; i < length/2; i++ {
		l[i], l[length-i-1] = l[length-i-1], l[i]
	}
}
