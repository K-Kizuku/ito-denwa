package base64

import (
	"context"
	"encoding/base64"
)

type typeCtxBase64Key struct{}

var CtxBase64Key typeCtxBase64Key = struct{}{}

func Encode(ctx context.Context, data []byte) string {
	v := ctx.Value(CtxBase64Key)
	if v == nil {
		return base64.StdEncoding.EncodeToString(data)
	}
	res, _ := v.(string)
	return res
}

func Decode(ctx context.Context, data string) ([]byte, error) {
	v := ctx.Value(CtxBase64Key)
	if v == nil {
		return base64.StdEncoding.DecodeString(data)
	}
	res, _ := v.(string)
	return base64.StdEncoding.DecodeString(res)
}
