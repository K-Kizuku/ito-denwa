package uuid

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		ctx      context.Context
		validate func(t *testing.T, got string)
	}{
		{
			name: "コンテキストにUUIDが設定されていない場合、新しいUUIDが生成される",
			ctx:  context.Background(),
			validate: func(t *testing.T, got string) {
				_, err := uuid.Parse(got)
				if err != nil {
					t.Errorf("生成された文字列が有効なUUIDではありません: %v", err)
				}
			},
		},
		{
			name: "コンテキストにUUIDが設定されている場合、そのUUIDが返される",
			ctx:  context.WithValue(context.Background(), CtxUUIDKey, "test-uuid"),
			validate: func(t *testing.T, got string) {
				if got != "test-uuid" {
					t.Errorf("期待値: test-uuid, 実際の値: %s", got)
				}
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := New(tt.ctx)
			tt.validate(t, got)
		})
	}
}
