package clock

import (
	"context"
	"testing"
	"time"
)

func TestWithTime(t *testing.T) {
	ctx := context.Background()
	now := time.Date(2024, 3, 15, 10, 0, 0, 0, time.UTC)

	ctxWithTime := WithTime(ctx, now)

	// コンテキストから時間を取得して検証
	if got := Now(ctxWithTime); !got.Equal(now) {
		t.Errorf("WithTime() = %v, want %v", got, now)
	}
}

func TestNow(t *testing.T) {
	tests := []struct {
		name string
		ctx  context.Context
		want time.Time
	}{
		{
			name: "コンテキストに時間が設定されていない場合",
			ctx:  context.Background(),
			want: time.Now(), // 現在時刻との比較は難しいため、実行時の時刻を使用
		},
		{
			name: "コンテキストに時間が設定されている場合",
			ctx:  WithTime(context.Background(), time.Date(2024, 3, 15, 10, 0, 0, 0, time.UTC)),
			want: time.Date(2024, 3, 15, 10, 0, 0, 0, time.UTC),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Now(tt.ctx)
			if tt.name == "コンテキストに時間が設定されていない場合" {
				// 現在時刻との比較は厳密な一致を期待できないため、スキップ
				return
			}
			if !got.Equal(tt.want) {
				t.Errorf("Now() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDate(t *testing.T) {
	tests := []struct {
		name  string
		year  int
		month int
		day   int
		loc   *time.Location
		want  time.Time
	}{
		{
			name:  "UTCでの日付作成",
			year:  2024,
			month: 3,
			day:   15,
			loc:   time.UTC,
			want:  time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC),
		},
		{
			name:  "JSTでの日付作成",
			year:  2024,
			month: 3,
			day:   15,
			loc:   time.FixedZone("JST", 9*60*60),
			want:  time.Date(2024, 3, 15, 0, 0, 0, 0, time.FixedZone("JST", 9*60*60)),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Date(tt.year, tt.month, tt.day, tt.loc)
			if !got.Equal(tt.want) {
				t.Errorf("Date() = %v, want %v", got, tt.want)
			}
		})
	}
}
