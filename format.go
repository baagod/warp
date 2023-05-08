package beam

import (
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
func ParseE(value string, loc ...*time.Location) (t Time, err error) {
	value = strings.Trim(value, `"`)

	var layout string
	for _, x := range patterns {
		if ok, _ := regexp.MatchString(x[1], value); ok {
			layout = x[0]
			break
		}
	}
	
	if loc == nil {
		loc = append(loc, time.UTC)
	}

	_t, err := time.ParseInLocation(layout, value, loc[0])
	return Time{_t}, err
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

func (t Time) String() string {
	return t.time.Format(time.DateTime + ".000")
}

func (t Time) StringOr(v string) string {
	if t.IsZero() {
		return v
	}
	return t.String()
}

func (t Time) DateOnly() string {
	return t.time.Format(time.DateOnly)
}

func (t Time) DateTime() string {
	return t.time.Format(time.DateTime)
}

func (t Time) Format(layout string) string {
	return t.time.Format(layout)
}
