package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExtractFirstJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "valid JSON object",
			input:    `Some text {"key": "value", "number": 42} more text`,
			expected: `{"key": "value", "number": 42}`,
			wantErr:  false,
		},
		{
			name:     "nested JSON object",
			input:    `{"outer": {"inner": "value"}, "array": [1, 2, 3]}`,
			expected: `{"outer": {"inner": "value"}, "array": [1, 2, 3]}`,
			wantErr:  false,
		},
		{
			name:    "no JSON in string",
			input:   `This is just plain text without JSON`,
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   ``,
			wantErr: true,
		},
		{
			name:    "malformed JSON",
			input:   `{"key": "value"`,
			wantErr: true,
		},
		{
			name:     "JSON with special characters",
			input:    `{"message": "Hello \"world\"!", "emoji": "🚀"}`,
			expected: `{"message": "Hello \"world\"!", "emoji": "🚀"}`,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ExtractFirstJSON(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestPrettyPrintJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:  "compact JSON",
			input: `{"name":"John","age":30,"city":"New York"}`,
			expected: `{
  "age": 30,
  "city": "New York",
  "name": "John"
}`,
			wantErr: false,
		},
		{
			name:    "invalid JSON",
			input:   `{"name":"John","age":}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := PrettyPrintJSON(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestValidateJSON(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
	}{
		{
			name:    "valid JSON object",
			input:   `{"key": "value"}`,
			wantErr: false,
		},
		{
			name:    "valid JSON array",
			input:   `[1, 2, 3]`,
			wantErr: false,
		},
		{
			name:    "invalid JSON",
			input:   `{"key": "value"`,
			wantErr: true,
		},
		{
			name:    "empty string",
			input:   ``,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateJSON(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseJSONToMap(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected map[string]interface{}
		wantErr  bool
	}{
		{
			name:  "valid JSON object",
			input: `{"name": "John", "age": 30}`,
			expected: map[string]interface{}{
				"name": "John",
				"age":  float64(30), // JSON numbers are parsed as float64
			},
			wantErr: false,
		},
		{
			name:    "invalid JSON",
			input:   `{"name": "John", "age":}`,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := ParseJSONToMap(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				assert.Nil(t, result)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
