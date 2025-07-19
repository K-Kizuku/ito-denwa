package null

import (
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		input    *int
		expected Null[int]
	}{
		{
			name:     "nilの場合",
			input:    nil,
			expected: Null[int]{Valid: false},
		},
		{
			name:     "値がある場合",
			input:    ptr(42),
			expected: Null[int]{Value: 42, Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := New(tt.input)
			if result.Valid != tt.expected.Valid {
				t.Errorf("New() Valid = %v, want %v", result.Valid, tt.expected.Valid)
			}
			if result.Valid && result.Value != tt.expected.Value {
				t.Errorf("New() Value = %v, want %v", result.Value, tt.expected.Value)
			}
		})
	}
}

func TestIsNull(t *testing.T) {
	tests := []struct {
		name     string
		input    Null[int]
		expected bool
	}{
		{
			name:     "nullの場合",
			input:    Null[int]{Valid: false},
			expected: true,
		},
		{
			name:     "値がある場合",
			input:    Null[int]{Value: 42, Valid: true},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.IsNull()
			if result != tt.expected {
				t.Errorf("IsNull() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestSetValue(t *testing.T) {
	tests := []struct {
		name     string
		input    *int
		expected Null[int]
	}{
		{
			name:     "nilを設定",
			input:    nil,
			expected: Null[int]{Valid: false},
		},
		{
			name:     "値を設定",
			input:    ptr(42),
			expected: Null[int]{Value: 42, Valid: true},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var n Null[int]
			n.SetValue(tt.input)
			if n.Valid != tt.expected.Valid {
				t.Errorf("SetValue() Valid = %v, want %v", n.Valid, tt.expected.Valid)
			}
			if n.Valid && n.Value != tt.expected.Value {
				t.Errorf("SetValue() Value = %v, want %v", n.Value, tt.expected.Value)
			}
		})
	}
}

func TestGetValue(t *testing.T) {
	tests := []struct {
		name        string
		input       Null[int]
		expected    int
		expectError bool
	}{
		{
			name:        "nullの場合",
			input:       Null[int]{Valid: false},
			expected:    0,
			expectError: true,
		},
		{
			name:        "値がある場合",
			input:       Null[int]{Value: 42, Valid: true},
			expected:    42,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := tt.input.GetValue()
			if tt.expectError {
				if err == nil {
					t.Error("GetValue() expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("GetValue() unexpected error: %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("GetValue() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGetValueOptional(t *testing.T) {
	tests := []struct {
		name     string
		input    Null[int]
		expected *int
	}{
		{
			name:     "nullの場合",
			input:    Null[int]{Valid: false},
			expected: nil,
		},
		{
			name:     "値がある場合",
			input:    Null[int]{Value: 42, Valid: true},
			expected: ptr(42),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.GetValueOptional()
			if tt.expected == nil {
				if result != nil {
					t.Error("GetValueOptional() expected nil but got value")
				}
				return
			}
			if result == nil {
				t.Error("GetValueOptional() expected value but got nil")
				return
			}
			if *result != *tt.expected {
				t.Errorf("GetValueOptional() = %v, want %v", *result, *tt.expected)
			}
		})
	}
}

// ヘルパー関数
func ptr[T any](v T) *T {
	return &v
}
