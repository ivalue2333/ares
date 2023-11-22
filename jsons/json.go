package jsons

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

// MarshalStr marshals v into json string
func MarshalStr(v interface{}) string {
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

// UnmarshalStr unmarshal data string into v.
func UnmarshalStr(data string, v interface{}) {
	Unmarshal([]byte(data), v)
}
