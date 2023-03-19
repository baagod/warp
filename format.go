package beam

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
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
		{"2006", `\d{4}$`},
	}
)

// Parse 解析 value 并返回它所表示的时间
func Parse[T int | int64 | string](value T) (t Time, err error) {
	v := fmt.Sprintf("%v", value)
	v = strings.Trim(v, `"`)

	_, err = strconv.ParseInt(v, 10, 64)
	if err == nil { // 解析时间戳
		if l := len(v); l >= 10 && l < 20 {
			sec, _ := strconv.ParseInt(v[:10], 10, 64)
			nsec, _ := strconv.ParseInt(v[10:], 10, 64)
			return Time{time: time.Unix(sec, nsec*1000000)}, nil
		}
		return t, errors.New("error: timestamp range is 10-19 digits")
	}

	var layout string
	for _, x := range patterns {
		if ok, _ := regexp.MatchString(x[1], v); ok {
			layout = x[0]
			break
		}
	}

	_t, err := time.ParseInLocation(layout, v, time.Local)
	return Time{_t}, err
}

// MustParse 返回忽略错误的 Parse()
func MustParse[T int | int64 | string](value T) Time {
	t, _ := Parse(value)
	return t
}

func (t Time) String() string {
	return t.time.Format(time.DateTime + ".000")
}

func (t Time) DateOnly() string {
	return t.time.Format(time.DateOnly)
}

func (t Time) DateTime() string {
	return t.time.Format(time.DateTime)
}
