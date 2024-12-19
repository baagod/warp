package warp

import (
	"fmt"
	"testing"
)

func TestWarp(t *testing.T) {
	pt, err := ParseE("01-02 03:04:05PM '06 -0700")
	fmt.Println("err:", err)
	fmt.Println("time:", pt)
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Parse("2024-01-02")
	}
}
