package beam_test

import (
	"testing"
)

<<<<<<< HEAD
func TestFunc(t *testing.T) {}
=======
func TestFunc(_ *testing.T) {
	now := beam.Now()

	fmt.Println(now.GoYear(2023))
}
>>>>>>> dc02c23e9210c5e6163749e96feea96e0b5d85d1
