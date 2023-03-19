package beam_test

import (
	"encoding/json"
	"fmt"
	"github.com/baa-god/beam"
	"testing"
)

func TestFunc(_ *testing.T) {
	type MM struct {
		Time beam.Time `json:"time"`
	}

	var m MM
	// m.Time = beam.Now()

	data := `{"time":"2023-03-19 23:17:18.944"}`

	_ = json.Unmarshal([]byte(data), &m)
	fmt.Println(m.Time)
}
