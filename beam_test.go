package beam_test

import (
	"fmt"
	"github.com/baa-god/beam"
	"testing"
)

func TestFunc(t *testing.T) {
	now := beam.Now()
	fmt.Println(now.StartMonth())
}
