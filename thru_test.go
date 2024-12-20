package thru

import (
	"fmt"
	"testing"
)

func TestWarp(t *testing.T) {
	pt, _ := ParseE("2024-03-01")
	// pt, _ := time.Parse(time.DateOnly, "2023-03-01")
	fmt.Println(pt.Start(0, 1))
}

func TestStart(t *testing.T) {
	pt, _ := ParseE("2024-03-02 15:04:05.999999999")
	fmt.Println(pt.Start())
}
