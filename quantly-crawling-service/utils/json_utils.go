package utils

import (
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
