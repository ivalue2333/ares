package ares

import (
	"encoding/json"
)

// Marshal marshals v into json bytes.
func Marshal(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return data
}

// MarshalString marshals v into json string
func MarshalString(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(data)
}

// Unmarshal data bytes into v.
func Unmarshal(data []byte, v interface{}) {
	_ = json.Unmarshal(data, v)
}

// UnmarshalString unmarshal data string into v.
func UnmarshalString(data string, v interface{}) {
	Unmarshal([]byte(data), v)
}
