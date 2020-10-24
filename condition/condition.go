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

// Quaternary ... Return one result of: both conditions satisfy, cond1 satisfy, cond2 satisfy, neither satisfy
func Quaternary(condition1, condition2 bool, both, cond1, cond2, neither interface{}) interface{} {
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

// HasZeroValue ... Check if any value is the zero value of its type (e.g. nil, numeric 0, empty string)
// note: if value is invalid, it is regarded as zero
func HasZeroValue(values ...interface{}) bool {
	for _, v := range values {
		if !reflect.ValueOf(v).IsValid() {
			return true
		}
		if reflect.ValueOf(v).IsZero() {
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
