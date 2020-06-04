# goutils [![Build Status](https://travis-ci.org/WayneJz/goutils.svg?branch=master)](https://travis-ci.org/WayneJz/goutils) [![Coverage Status](https://coveralls.io/repos/github/WayneJz/goutils/badge.svg?branch=master)](https://coveralls.io/github/WayneJz/goutils?branch=master)
Utility Code for Golang, for better error collection, slice manipulation and orms handling

Require Go 1.11 or newer

### slice utility

Support int32, int64, uint32, uint64, string slices

- Apply a function to a slice (like lambda function in other language)

- Check if a slice contains element(s)

- Index element(s) in a slice

- Merge slices together (with unique merge option)

- Stably sort or reversely sort slices

- Convert a slice to a unique slice (e.g. elements are unique)

### errors

- An error receiver, which receives errors and output them in traceback mode

### orms

- Concurrent [gorm](https://github.com/jinzhu/gorm) execution (use go routine), with timeout option

## License

This project is released under MIT license
