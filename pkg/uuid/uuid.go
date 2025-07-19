package uuid

import (
	"context"

	"github.com/google/uuid"
)

type typeCtxUUIDKey struct{}

var CtxUUIDKey typeCtxUUIDKey = struct{}{}

func New(ctx context.Context) string {
	v := ctx.Value(CtxUUIDKey)
	if v == nil {
		return uuid.Must(uuid.NewV7()).String()
	}
	res, _ := v.(string)
	return res
}

func Parse(value string) (string, error) {
	id, err := uuid.Parse(value)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func IsValid(value string) bool {
	return uuid.Validate(value) == nil
}
