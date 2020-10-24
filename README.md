# goutils [![Build Status](https://travis-ci.org/WayneJz/goutils.svg?branch=master)](https://travis-ci.org/WayneJz/goutils) [![Coverage Status](https://coveralls.io/repos/github/WayneJz/goutils/badge.svg?branch=master)](https://coveralls.io/github/WayneJz/goutils?branch=master)
Utility Code for Golang, for better map, slice manipulation and condition handling

Require Go 1.11 or newer

## maps

- **JSON Map:** marshal JSON-format string(s) to map(s) in one function

- **Merge:** merge maps together

- **Sequence:** extract map keys **in-order**, with corresponding values

Support int32, int64, uint32, uint64, string key maps

## slice

- **Apply:** apply a function to a slice (like **lambda** function in other language)

- **Contain:** check if a slice contains given element(s)

- **Index:** find indexes of given element(s) in a slice

- **Merge**: merge slices together (with unique merge option)

- **Reverse**: reverse a slice in-place

- **Sort**: stable (reverse) sort slices, by length or numerical sequence

- **Unique**: convert a slice to a unique slice (i.e. elements are unique)

Support int32, int64, uint32, uint64, string, bytes slices

## condition

- **Any, All:** judge multiple boolean conditions in one function

- **Ternary:** ternary expression

- **Quaternary:** quaternary expression

- **HasZeroValue:** check if any value is the zero value of its type (e.g. nil, numeric zero)

- **HasNumber, HasLetter:** check if a string contains number/letter


## License

This project is released under MIT license
