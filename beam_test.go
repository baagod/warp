package beam

import (
	"fmt"
	"testing"
)

func TestFunc(t *testing.T) {
	fmt.Println(Time{}.StringOr("default"))
}
