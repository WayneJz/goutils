package maps

import "encoding/json"

// StringToJSONMap ... Marshal a JSON-format string to map
func StringToJSONMap(str string) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(str), &m); err != nil {
		return nil, err
	}
	return m, nil
}

// MustStringToJSONMap ... Marshal a JSON-format string to map, ignore error
func MustStringToJSONMap(str string) map[string]interface{} {
	m, _ := StringToJSONMap(str)
	return m
}

// StringsToJSONMaps ... Marshal multiple JSON-format strings to a slice of maps
func StringsToJSONMaps(strs ...string) (res []map[string]interface{}, err error) {
	for _, str := range strs {
		m, err := StringToJSONMap(str)
		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return
}

// MustStringsToJSONMaps ... Marshal multiple JSON-format strings to a slice of maps, ignore error
func MustStringsToJSONMaps(strs ...string) (res []map[string]interface{}) {
	for _, str := range strs {
		m, _ := StringToJSONMap(str)
		res = append(res, m)
	}
	return
}
