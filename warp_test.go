package warp

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Order struct {
	Time Time `gorm:"column:fldSC_SetTime" json:"fldSC_SetTime"` // 手术结算月
}

func TestFunc(t *testing.T) {
	data := `{"fldSC_SetTime": "0001-01-01 00:00:00"}`
	var o Order

	err := json.Unmarshal([]byte(data), &o)
	fmt.Println("Unmarshal error:", err)
	fmt.Printf("order: %+v, isZero: %v\n", o, o.Time.IsZero())
}
