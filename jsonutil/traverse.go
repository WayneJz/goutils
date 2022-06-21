package jsonutil

import (
	"bytes"
	"strconv"
	"strings"

	"github.com/tidwall/gjson"
)

// Key json key
type Key struct {
	Key  interface{}
	Type KeyType
}

// KeyType the json key type
type KeyType int

const (
	// String regular key-value json key type
	String KeyType = iota + 1

	// Index is array key type
	Index
)

// TraverseCallback callback function, will be called during traverse
type TraverseCallback func(key *Key, value *gjson.Result, path []byte)

// Traverse traverse json object and call the callback function
func Traverse(b string, callback TraverseCallback) {
	if !gjson.Valid(b) {
		return
	}
	root := gjson.Parse(b)
	traverse(nil, &root, nil, callback)
}

// TraverseBytes is like Traverse but more efficient if use bytes
func TraverseBytes(b []byte, callback TraverseCallback) {
	if !gjson.ValidBytes(b) {
		return
	}
	root := gjson.ParseBytes(b)
	traverse(nil, &root, nil, callback)
}

// traverse recursively traverse json object
func traverse(key *Key, value *gjson.Result, path []byte, callback TraverseCallback) {
	path = addPath(path, key)

	switch value.Type {
	case gjson.String, gjson.Number, gjson.Null, gjson.True, gjson.False:
		callback(key, value, path)

	case gjson.JSON:
		if value.IsArray() {
			for i, obj := range value.Array() {
				traverse(&Key{Key: i, Type: Index}, &obj, path, callback)
			}
		}
		if value.IsObject() {
			for k, obj := range value.Map() {
				traverse(&Key{Key: k, Type: String}, &obj, path, callback)
			}
		}
	default:
	}
}

// addPath add json absoulte path, syntax based on sjson:  https://github.com/tidwall/sjson#path-syntax
func addPath(path []byte, key *Key) []byte {
	if key == nil {
		return path
	}
	buffer := bytes.NewBuffer(path)

	switch key.Type {
	case Index:
		if buffer.Len() > 0 {
			buffer.WriteRune('.')
		}
		buffer.WriteString(strconv.Itoa(key.Key.(int)))

	case String:
		if buffer.Len() > 0 {
			_, err := strconv.Atoi(key.Key.(string))
			if err != nil {
				buffer.WriteRune('.')
			} else {
				buffer.WriteString(".:") // support numeric object key: https://github.com/tidwall/sjson#path-syntax
			}
		}
		buffer.WriteString(strings.ReplaceAll(key.Key.(string), ".", "\\."))
	}
	return buffer.Bytes()
}
