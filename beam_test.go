package beam_test

import (
	"fmt"
	"github.com/baa-god/beam"
	"testing"
)

func TestFunc(_ *testing.T) {
	now := beam.Now()

	fmt.Println(now.Before(beam.Parse("2024"))) // 在指定时间之前 (t < u): true
	fmt.Println(now.After(beam.Parse("2024")))  // 在指定时间之后 (t > u): false
	fmt.Println(now.Equal(beam.Parse("2024")))  // 是否相等: false
}
