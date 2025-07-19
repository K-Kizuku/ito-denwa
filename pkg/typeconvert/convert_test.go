package typeconvert

import "testing"

func TestToPtr(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected any
	}{
		{
			name:     "int",
			input:    42,
			expected: 42,
		},
		{
			name:     "string",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "float",
			input:    3.14,
			expected: 3.14,
		},
		{
			name:     "bool",
			input:    true,
			expected: true,
		},
		{
			name:     "struct",
			input:    struct{ Name string }{Name: "test"},
			expected: struct{ Name string }{Name: "test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ptr := ToPtr(tt.input)
			if *ptr != tt.expected {
				t.Errorf("ToPtr() = %v, want %v", *ptr, tt.expected)
			}
		})
	}
}
