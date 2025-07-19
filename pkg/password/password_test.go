package password

import (
	"testing"
)

func TestEncrypt(t *testing.T) {
	tests := []struct {
		name     string
		password string
		wantErr  bool
	}{
		{
			name:     "正常系：パスワードの暗号化",
			password: "testpassword123",
			wantErr:  false,
		},
		{
			name:     "正常系：空のパスワード",
			password: "",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Encrypt(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == "" {
				t.Error("Encrypt() returned empty string for valid password")
			}
		})
	}
}

func TestCompare(t *testing.T) {
	password := "testpassword123"
	hashedPassword, err := Encrypt(password)
	if err != nil {
		t.Fatalf("Failed to encrypt password for test: %v", err)
	}

	tests := []struct {
		name            string
		hashedPassword  string
		requestPassword string
		wantErr         bool
	}{
		{
			name:            "正常系：正しいパスワード",
			hashedPassword:  hashedPassword,
			requestPassword: password,
			wantErr:         false,
		},
		{
			name:            "異常系：間違ったパスワード",
			hashedPassword:  hashedPassword,
			requestPassword: "wrongpassword",
			wantErr:         true,
		},
		{
			name:            "異常系：空のパスワード",
			hashedPassword:  hashedPassword,
			requestPassword: "",
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Compare(tt.hashedPassword, tt.requestPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compare() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
