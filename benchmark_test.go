package strftime_test

import (
	"testing"
	"time"

	fastly "github.com/fastly/go-utils/strftime"
	imperfectgo "github.com/imperfectgo/go-strftime"
	itchyny "github.com/itchyny/timefmt-go"
	jehiah "github.com/jehiah/go-strftime"
	lestrrat "github.com/lestrrat-go/strftime"
	tebeka "github.com/tebeka/strftime"
)

var now = time.Now().UTC()

const (
	layout       = "%Y-%m-%d"
	layoutString = "2006-01-02"
)

func BenchmarkStdTimeFormat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = now.Format(layoutString)
	}
}

func BenchmarkImperfectgo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = imperfectgo.Format(now, layout)
	}
}

func BenchmarkItchyny(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = itchyny.Format(now, layout)
	}
}

func BenchmarkFastly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = fastly.Strftime(layout, now)
	}
}

func BenchmarkJehiah(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = jehiah.Format(layout, now)
	}
}

func BenchmarkLestrrat(b *testing.B) {
	f, _ := lestrrat.New(layout)
	for i := 0; i < b.N; i++ {
		_ = f.FormatString(now)
	}
}

func BenchmarkTebeka(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = tebeka.Format(layout, now)
	}
}
