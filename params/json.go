package paramsUtils

import (
	"encoding/json"
)

func TextToJson(text string) (map[string]interface{}, error) {
	var data map[string]interface{}

	return data, json.Unmarshal([]byte(text), &data)
}
