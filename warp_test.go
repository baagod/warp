package warp

import (
	"fmt"
	"testing"
	"time"
)

func TestWarp(t *testing.T) {
	pt := Parse("2024-01-31")
	pt = pt.AddYear(0, 1) // 2023-01-31

	// input: 2024-05-01
	fmt.Println("time:", pt.Format(time.DateTime))
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Parse("2024-01-02")
	}
}
