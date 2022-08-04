package jsonutil

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// TimeLayout a common layout time format
const (
	TimeLayout     = "2006-01-02 15:04:05"
	jsonTimeLayout = "\"2006-01-02 15:04:05\""
)

// JSONTime is a common layout time format, fully compatible with JSON and GORM
type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(time.Time(t).Format(jsonTimeLayout)), nil
}

func (t *JSONTime) UnmarshalJSON(data []byte) error {
	tm, err := time.ParseInLocation(jsonTimeLayout, string(data), time.Local)
	if err != nil {
		return err
	}
	*t = JSONTime(tm)
	return nil
}

func (t JSONTime) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
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
