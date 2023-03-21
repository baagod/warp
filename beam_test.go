package beam_test

import (
	"fmt"
	"github.com/baa-god/beam"
	"testing"
)

func TestFunc(_ *testing.T) {
	now := beam.Now()

	fmt.Println(now.GoYear(2023))
}
