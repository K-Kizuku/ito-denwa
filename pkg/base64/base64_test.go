package base64

import (
	"context"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		input    []byte
		expected string
	}{
		{
			name:     "標準のエンコード",
			ctx:      context.Background(),
			input:    []byte("Hello, World!"),
			expected: "SGVsbG8sIFdvcmxkIQ==",
		},
		{
			name:     "コンテキストに値がある場合",
			ctx:      context.WithValue(context.Background(), CtxBase64Key, "SGVsbG8sIFdvcmxkIQ=="),
			input:    []byte("Hello, World!"),
			expected: "SGVsbG8sIFdvcmxkIQ==",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Encode(tt.ctx, tt.input)
			if result != tt.expected {
				t.Errorf("Encode() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name        string
		ctx         context.Context
		input       string
		expected    []byte
		expectError bool
	}{
		{
			name:        "標準のデコード",
			ctx:         context.Background(),
			input:       "SGVsbG8sIFdvcmxkIQ==",
			expected:    []byte("Hello, World!"),
			expectError: false,
		},
		{
			name:        "コンテキストに値がある場合",
			ctx:         context.WithValue(context.Background(), CtxBase64Key, "SGVsbG8sIFdvcmxkIQ=="),
			input:       "dummy",
			expected:    []byte("Hello, World!"),
			expectError: false,
		},
		{
			name:        "不正なbase64文字列",
			ctx:         context.Background(),
			input:       "invalid-base64!",
			expected:    nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := Decode(tt.ctx, tt.input)
			if tt.expectError {
				if err == nil {
					t.Error("Decode() expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("Decode() unexpected error: %v", err)
				return
			}
			if string(result) != string(tt.expected) {
				t.Errorf("Decode() = %v, want %v", result, tt.expected)
			}
		})
	}
}
