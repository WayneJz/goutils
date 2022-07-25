package jsonutil

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

// TimeLayout common layout time format in China
const TimeLayout = "2006-01-02 15:04:05"

// JSONTime common layout time struct format in China, support JSON and Gorm
type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteRune('"')
	buf.WriteString(time.Time(t).Format(TimeLayout))
	buf.WriteRune('"')
	return buf.Bytes(), nil
}

func (t JSONTime) Value() (driver.Value, error) {
	if time.Time(t).UnixNano() == (time.Time{}).UnixNano() {
		return nil, nil
	}
	return time.Time(t), nil
}

func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime(value)
		return nil
	}
	return fmt.Errorf("cannot convert %v to jsontime", v)
}

func (t JSONTime) ToTime() time.Time {
	return time.Time(t)
}
