package pg

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type JSON json.RawMessage

// Returns the *j as the JSON encoding of j.
func (j *JSON) MarshalJSON() ([]byte, error) {
	if len([]byte(*j)) == 0 {
		return []byte("null"), nil
	}
	return *j, nil
}

// UnmarshalJSON sets *j to a copy of data
func (j *JSON) UnmarshalJSON(data []byte) error {
	if j == nil {
		return errors.New("JsonText: UnmarshalJSON on nil pointer")
	}
	*j = append((*j)[0:0], data...)
	return nil

}

// Value returns j as a value.  This does a validating unmarshal into another
// RawMessage.  If j is invalid json, it returns an error.
func (j JSON) Value() (driver.Value, error) {
	var m json.RawMessage
	var err = j.Unmarshal(&m)
	if err != nil {
		return []byte{}, err
	}
	return []byte(j), nil
}

// Scan stores the src in *j.  No validation is done.
func (j *JSON) Scan(src interface{}) error {
	var source []byte
	switch src.(type) {
	case string:
		source = []byte(src.(string))
	case []byte:
		source = src.([]byte)
	default:
		source = []byte("null")
	}
	*j = JSON(append((*j)[0:0], source...))
	return nil
}

// Unmarshal unmarshal's the json in j to v, as in json.Unmarshal.
func (j *JSON) Unmarshal(v interface{}) error {
	return json.Unmarshal([]byte(*j), v)
}
