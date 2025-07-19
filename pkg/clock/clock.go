package clock

import (
	"context"
	"time"
)

const (
	DateTimeHyphenLayout = "2006/01/02 15:04:05"
)

var (
	UTC = time.UTC
)

type typeCtxClockKey struct{}

var ctxClockKey typeCtxClockKey = struct{}{}

func WithTime(ctx context.Context, t time.Time) context.Context {
	return context.WithValue(ctx, ctxClockKey, t)
}

func Now(ctx context.Context) time.Time {
	v := ctx.Value(ctxClockKey)
	if v == nil {
		return time.Now()
	}
	res, _ := v.(time.Time)
	return res
}

func Date(year, month, day int, loc *time.Location) time.Time {
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
}
