package maps

// MergeStringMaps ... Merge string maps into one string map
func MergeStringMaps(maps ...map[string]interface{}) map[string]interface{} {
	res := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// MergeInt64Maps ... Merge int64 maps into one int64 map
func MergeInt64Maps(maps ...map[int64]interface{}) map[int64]interface{} {
	res := make(map[int64]interface{})
	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// MergeInt32Maps ... Merge int32 maps into one int32 map
func MergeInt32Maps(maps ...map[int32]interface{}) map[int32]interface{} {
	res := make(map[int32]interface{})
	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// MergeUint64Maps ... Merge uint64 maps into one uint64 map
func MergeUint64Maps(maps ...map[uint64]interface{}) map[uint64]interface{} {
	res := make(map[uint64]interface{})
	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}

// MergeUint32Maps ... Merge uint32 maps into one uint32 map
func MergeUint32Maps(maps ...map[uint32]interface{}) map[uint32]interface{} {
	res := make(map[uint32]interface{})
	for _, m := range maps {
		for k, v := range m {
			res[k] = v
		}
	}
	return res
}
