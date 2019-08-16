package plugin

import (
	"encoding/json"
)

type DataProcessor interface {
	Marshal(object interface{}) ([]byte, error)

	UnMarshal(data []byte, v interface{}) error
}

type JsonProcessor struct{}

func (p JsonProcessor) Marshal(object interface{}) ([]byte, error) {
	return json.Marshal(object)
}

func (p JsonProcessor) UnMarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
