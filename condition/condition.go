package condition

import (
	"reflect"
	"unicode"
)

// Any ... Check if any boolean conditions satisfy
func Any(conditions ...bool) bool {
	for _, c := range conditions {
		if c {
			return true
		}
	}
	return false
}

// All ... Check if all boolean conditions satisfy
func All(conditions ...bool) bool {
	for _, c := range conditions {
		if !c {
			return false
		}
	}
	return true
}

// Ternary ... If condition satisfies, return a, else return b
func Ternary(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}
	return b
}

// Quaternion ... Return one result of: both conditions satisfy, cond1 satisfy, cond2 satisfy, neither satisfy
func Quaternion(condition1, condition2 bool, both, cond1, cond2, neither interface{}) interface{} {
	if condition1 && condition2 {
		return both
	}
	if condition1 {
		return cond1
	}
	if condition2 {
		return cond2
	}
	return neither
}

// AnyValid ... Check if any data valid (e.g. not nil, non-zero value)
func AnyValid(data ...interface{}) bool {
	for _, d := range data {
		if reflect.ValueOf(d).IsValid() {
			return true
		}
	}
	return false
}

// AllValid ... Check if all data valid (e.g. not nil, non-zero value)
func AllValid(data ...interface{}) bool {
	for _, d := range data {
		if !reflect.ValueOf(d).IsValid() {
			return false
		}
	}
	return true
}

// HasNilPtr ... Check if any data is a nil pointer
func HasNilPtr(data ...interface{}) bool {
	for _, d := range data {
		if reflect.TypeOf(d).Kind() == reflect.Ptr && reflect.ValueOf(d).IsNil() {
			return true
		}
	}
	return false
}

// HasNumber ... Check if a string contains number
func HasNumber(s string) bool {
	for _, r := range s {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

// HasLetter ... Check if a string contains letter
func HasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) {
			return true
		}
	}
	return false
}
