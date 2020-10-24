package slice

func ReverseStrings(l []string) {
	length := len(l)
	for i := 0; i < length/2; i++ {
		l[i], l[length-i] = l[length-i], l[i]
	}
}

func ReverseInt64s(l []int64) {
	length := len(l)
	for i := 0; i < length/2; i++ {
		l[i], l[length-i] = l[length-i], l[i]
	}
}

func ReverseInt32s(l []int32) {
	length := len(l)
	for i := 0; i < length/2; i++ {
		l[i], l[length-i] = l[length-i], l[i]
	}
}

func ReverseUint64s(l []uint64) {
	length := len(l)
	for i := 0; i < length/2; i++ {
		l[i], l[length-i] = l[length-i], l[i]
	}
}

func ReverseUint32s(l []uint32) {
	length := len(l)
	for i := 0; i < length/2; i++ {
		l[i], l[length-i] = l[length-i], l[i]
	}
}
