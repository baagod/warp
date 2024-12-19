package warp

import (
	"fmt"
	"testing"
)

func TestWarp(t *testing.T) {
	pt, _ := ParseE("2023-01-31 13:14:15.123")
	// pt, _ := time.Parse(time.DateTime, "2023-01-31 13:14:15")
	// fmt.Println(pt.AddYear(1, 1))
	fmt.Println(pt.End(1, 1))

	// 2024-03-02 13:14:15 +0000 UTC
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Parse("2024-01-02")
	}
}
