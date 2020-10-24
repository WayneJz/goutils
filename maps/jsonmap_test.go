package maps

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestStringToJSONMap(t *testing.T) {
	Convey("TestStringToJSONMap", t, func() {
		Convey("Embedded JSON", func() {
			testStr := "teststr"
			str := func() string {
				var Test = struct {
					TestInt    int32
					TestStrPtr *string
					TestMap    map[string]interface{}
				}{
					TestInt:    123,
					TestStrPtr: &testStr,
					TestMap:    map[string]interface{}{"testmap": "ok"},
				}
				b, _ := json.Marshal(&Test)
				return string(b)
			}()
			res, err := StringToJSONMap(str)
			So(res, ShouldContainKey, "TestInt")
			So(res, ShouldContainKey, "TestStrPtr")
			So(res, ShouldContainKey, "TestMap")
			So(err, ShouldBeNil)
		})
		Convey("JSON unmarshal failure", func() {
			testStr := "teststr"
			res, err := StringToJSONMap(testStr)
			So(res, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestMustStringToJSONMap(t *testing.T) {
	Convey("TestMustStringToJSONMap", t, func() {
		Convey("Embedded JSON", func() {
			testStr := "teststr"
			str := func() string {
				var Test = struct {
					TestInt    int32
					TestStrPtr *string
					TestMap    map[string]interface{}
				}{
					TestInt:    123,
					TestStrPtr: &testStr,
					TestMap:    map[string]interface{}{"testmap": "ok"},
				}
				b, _ := json.Marshal(&Test)
				return string(b)
			}()
			res := MustStringToJSONMap(str)
			So(res, ShouldContainKey, "TestInt")
			So(res, ShouldContainKey, "TestStrPtr")
			So(res, ShouldContainKey, "TestMap")
		})
		Convey("JSON unmarshal failure", func() {
			testStr := "teststr"
			res := MustStringToJSONMap(testStr)
			So(res, ShouldBeEmpty)
		})
	})
}

func TestStringsToJSONMaps(t *testing.T) {
	Convey("TestStringsToJSONMaps", t, func() {
		testStr := "teststr"
		str := func() string {
			var Test = struct {
				TestInt    int32
				TestStrPtr *string
				TestMap    map[string]interface{}
			}{
				TestInt:    123,
				TestStrPtr: &testStr,
				TestMap:    map[string]interface{}{"testmap": "ok"},
			}
			b, _ := json.Marshal(&Test)
			return string(b)
		}()
		Convey("Embedded JSON", func() {
			res, err := StringsToJSONMaps(str)
			So(res, ShouldHaveLength, 1)
			So(res[0], ShouldContainKey, "TestInt")
			So(res[0], ShouldContainKey, "TestStrPtr")
			So(res[0], ShouldContainKey, "TestMap")
			So(err, ShouldBeNil)
		})
		Convey("JSON unmarshal failure", func() {
			res, err := StringsToJSONMaps(str, testStr)
			So(res, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})
	})
}

func TestMustStringsToJSONMaps(t *testing.T) {
	Convey("TestMustStringsToJSONMaps", t, func() {
		testStr := "teststr"
		str := func() string {
			var Test = struct {
				TestInt    int32
				TestStrPtr *string
				TestMap    map[string]interface{}
			}{
				TestInt:    123,
				TestStrPtr: &testStr,
				TestMap:    map[string]interface{}{"testmap": "ok"},
			}
			b, _ := json.Marshal(&Test)
			return string(b)
		}()
		Convey("Embedded JSON", func() {
			res := MustStringsToJSONMaps(str)
			So(res, ShouldHaveLength, 1)
			So(res[0], ShouldContainKey, "TestInt")
			So(res[0], ShouldContainKey, "TestStrPtr")
			So(res[0], ShouldContainKey, "TestMap")
		})
		Convey("JSON unmarshal failure", func() {
			res := MustStringsToJSONMaps(str, testStr)
			So(res, ShouldHaveLength, 2)
			So(res[0], ShouldContainKey, "TestInt")
			So(res[0], ShouldContainKey, "TestStrPtr")
			So(res[0], ShouldContainKey, "TestMap")
			So(res[1], ShouldBeEmpty)
		})
	})
}
