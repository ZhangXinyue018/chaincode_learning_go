package lib

import (
	"encoding/json"
	"strconv"
	"strings"
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

func StringToBytes(str string) []byte {
	inputs := strings.Split(str, ",")
	requestBytes := make([]byte, 0)
	for _, input := range inputs {
		inputv, _ := strconv.Atoi(input)
		requestBytes = append(requestBytes, byte(inputv))
	}
	return requestBytes
}

func ByteToString(bytes []byte) string {
	if len(bytes) == 0 {
		return ""
	}
	resultStr := strconv.Itoa(int(bytes[0]))
	for index, value := range bytes {
		if index != 0 {
			temp := strconv.Itoa(int(value))
			resultStr = resultStr + "," + temp
		}
	}
	return resultStr
}
