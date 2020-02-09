package ding

import (
	"encoding/json"
	"fmt"
)

func Json2map(jsonStr string) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	return jsonMap, err
}

func MapToJson(obj map[string]interface{}) (*[]byte, error) {
	jsonObj, err := json.Marshal(obj)
	return &jsonObj, err
}

func ToString(s interface{}) string {
	return fmt.Sprintf("%v", s)
}
