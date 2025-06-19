package dto

import "fmt"

type PositionCloseAllResponseDto struct {
	Symbol string              `json:"symbol"`
	Status int32               `json:"status"`
	Body   OrderGetResponseDto `json:"body"`
}

func (p PositionCloseAllResponseDto) String() string {
	return fmt.Sprintf(
		`PositionCloseAllResponseDto:
  Symbol: %s
  Status: %d
  Body:
    %s
`, p.Symbol, p.Status, indent(p.Body.String(), "    "))
}

// Helper function to indent multi-line nested output
func indent(s, prefix string) string {
	lines := []rune(s)
	result := ""
	for i, r := range lines {
		if i == 0 {
			result += prefix
		}
		result += string(r)
		if r == '\n' && i < len(lines)-1 {
			result += prefix
		}
	}
	return result
}
