package lib

import (
	"encoding/json"
)

func ToJsonBytes(entity interface{}) ([]byte, error) {
	return json.Marshal(entity)
}

func ToJsonString(entity interface{}) string {
	result, err := json.MarshalIndent(entity, "", "    ")
	if err != nil {
		return ""
	}
	return string(result)
}

func FromJsonBytes(data []byte, entity interface{}) error {
	return json.Unmarshal(data, entity)
}
