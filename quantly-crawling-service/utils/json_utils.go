package utils

import (
	"encoding/json"

	"github.com/trmviet0801/quantly/quantly-crawling-serivce/models"
)

func WrapJSONObjectAsArray(data []byte) (string, bool) {
	var result string
	var stack models.Stack[string]

	if len(data) == 0 {
		return result, false
	}

	result += "[\n"
	for _, character := range string(data) {
		if string(character) == "{" {
			stack.Push("{")
			result += "{"
		} else if string(character) == "}" {
			_, isOk := stack.Pop()
			if !isOk {
				return result, false
			}

			result += "}"

			if stack.IsEmpty() {
				result += ",\n"
			}
		} else {
			result += string(character)
		}
	}

	//Unclose - invalid format
	if !stack.IsEmpty() {
		return result, false
	}

	result = result[:len(result)-3]
	result += "]"
	return result, true
}

// Convert fields (map[string]interface{} to Struct)
func MapToStruct[T any](fields map[string]string) (T, error) {
	var result T

	jsonText, err := json.Marshal(fields)
	if err != nil {
		OnError(err)
		return result, err
	}

	err = json.Unmarshal(jsonText, &result)
	if err != nil {
		OnError(err)
		return result, err
	}

	return result, nil
}
