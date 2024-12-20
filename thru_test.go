package thru

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Model struct {
	Time Time `json:"time"`
}

func (m Model) MarshalJSON() ([]byte, error) {
	type Alias Model
	fmt.Printf("%p\n", &m.Time) // 0xc0000940c0
	a := (*Alias)(&m)
	a.Time = m.Time.Layout(DateTime)
	fmt.Printf("%p\n", &m.Time) // 0xc0000940c0
	return json.Marshal(a)
}

func TestWarp(t *testing.T) {
	pt, _ := ParseE("2024-12-20T09:33:59.9549583+08:00")
	// pt, _ := time.Parse(time.DateTime, "2023-01-31 13:14:15")
	// fmt.Println(pt.AddYear(1, 1))
	fmt.Println(pt)

	var m Model
	m.Time = Now()

	b, _ := json.Marshal(&m)
	fmt.Println(string(b))
}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Parse("2024-01-02")
	}
}
