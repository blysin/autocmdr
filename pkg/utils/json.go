// Package utils provides common utility functions including JSON manipulation and path handling.
package utils

import (
	"encoding/json"
	"fmt"
)

// ExtractFirstJSON extracts the first valid JSON object from a string
func ExtractFirstJSON(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("input string is empty")
	}

	var stack []rune
	var start, end int
	found := false

	for i, ch := range input {
		switch ch {
		case '{':
			if len(stack) == 0 {
				start = i // Record JSON start position
			}
			stack = append(stack, ch)
		case '}':
			if len(stack) > 0 {
				stack = stack[:len(stack)-1] // Pop the top '{'
				if len(stack) == 0 {
					end = i + 1 // Record JSON end position
					found = true
					break
				}
			}
		}
	}

	if !found || end <= start {
		return "", fmt.Errorf("no valid JSON found in input")
	}

	jsonStr := input[start:end]

	// Validate if it's valid JSON
	var js map[string]interface{}
	if err := json.Unmarshal([]byte(jsonStr), &js); err != nil {
		return "", fmt.Errorf("extracted string is not valid JSON: %w", err)
	}

	return jsonStr, nil
}

// PrettyPrintJSON formats JSON string with indentation
func PrettyPrintJSON(jsonStr string) (string, error) {
	var obj interface{}
	if err := json.Unmarshal([]byte(jsonStr), &obj); err != nil {
		return "", fmt.Errorf("invalid JSON: %w", err)
	}

	pretty, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to format JSON: %w", err)
	}

	return string(pretty), nil
}

// ValidateJSON checks if a string is valid JSON
func ValidateJSON(jsonStr string) error {
	var js interface{}
	return json.Unmarshal([]byte(jsonStr), &js)
}

// ParseJSONToMap parses JSON string to map[string]interface{}
func ParseJSONToMap(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	return result, err
}
