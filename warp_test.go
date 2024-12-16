package warp

import (
	"fmt"
	"testing"
	"time"
)

func TestWarp(t *testing.T) {
	pt := Parse("2024-02-29").Time()
	pt = pt.AddDate(1, 0, 0)
	// pt, err := time.Parse(time.Stamp, "Feb 12 15:04:05.221")
	// // pt, err := ParseE("January 12 15:04:05")
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	//
	fmt.Println("time:", pt.Format(time.DateTime+".999"))
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Parse("2024-01-02")
	}
}
