package warp

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

var (
	patterns = [][2]string{
		{"2006-01-02 15:04:05", `\d{4}(-\d{2}){2} \d{2}(:\d{2}){2}(.\d{1,9})?`},
		{"2006/01/02 15:04:05", `\d{4}(\/\d{2}){2} \d{2}(:\d{2}){2}(.\d{1,9})?`},
		{"2006-01-02 15:04", `\d{4}(-\d{2}){2} \d{2}:\d{2}`},
		{"2006/01/02 15:04", `\d{4}(\/\d{2}){2} \d{2}:\d{2}`},
		{"2006-01-02 15", `\d{4}(-\d{2}){2} \d{2}`},
		{"2006/01/02 15", `\d{4}(\/\d{2}){2} \d{2}`},
		{"2006-01-02", `\d{4}(-\d{2}){2}`},
		{"15:04:05", `\d{2}(:\d{2}){2}(.\d{1,9})?`},
		{"2006-01", `\d{4}-\d{2}`},
		{"15:04", `\d{2}:\d{2}`},
		{"2006", `\d{4}`},
	}
)

// ParseE 解析 value 并返回它所表示的时间
func ParseE(value string, loc ...*time.Location) (Time, error) {
	value = strings.Trim(value, `"`)

	var layout string
	for _, x := range patterns {
		if ok, _ := regexp.MatchString(x[1], value); ok {
			layout = x[0]
			break
		}
	}

	if loc == nil {
		loc = append(loc, time.Local)
	}

	pt, err := time.ParseInLocation(layout, value, loc[0])
	if err == nil && pt.Year() < 1000 {
		return Time{}, errors.New("年份超出有效范围")
	}

	return Time{pt}, err
}

// Parse 返回忽略错误的 ParseE()
func Parse(value string, loc ...*time.Location) Time {
	t, _ := ParseE(value, loc...)
	return t
}

// LayoutE 通过 layout 和 value 解析并返回它所表示的时间
func LayoutE(layout string, value string) (Time, error) {
	t, err := time.ParseInLocation(layout, value, time.Local)
	return Time{time: t}, err
}

// Layout 返回忽略错误的 LayoutE()
func Layout(layout string, value string) (t Time) {
	t, _ = LayoutE(layout, value)
	return
}

func (t Time) DateTime() string {
	return t.time.Format(time.DateTime)
}
