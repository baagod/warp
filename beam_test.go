package beam_test

import (
	"fmt"
	"github.com/baa-god/beam"
	"testing"
)

func TestFunc(_ *testing.T) {
	now := beam.Parse("2023-03-19 15:05:27.792100231")

	fmt.Println(now.Unix())
	fmt.Println(beam.Parse(now.Unix()))
	// beam.New(time.Now())            // 使用 time.Time 创建: 2023-03-19 21:47:36.533
	// beam.Unix(now.Unix(), 0)        // 使用秒时间戳: 2023-03-19 15:05:27.000
	// beam.UnixMilli(now.UnixMilli()) // 使用毫秒时间戳: 2023-03-19 15:05:27.792
	// beam.UnixMicro(now.UnixMicro()) // 使用微秒时间戳: 2023-03-19 15:05:27.792
	// beam.UnixNano(now.UnixNano())   // 使用纳秒时间戳: 2023-03-19 15:05:27.792
}
