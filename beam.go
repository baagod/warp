package beam

import (
	"database/sql/driver"
	"math"
	"time"
)

type Time struct {
	time time.Time
}

func New(t time.Time) Time {
	return Time{time: t}
}

func Now() Time {
	return Time{time: time.Now()}
}

var (
	days = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
)

// AddYear 添加 年[, 月[, 日]]。若不传 [天数] 则月份不会溢出。
func (t Time) AddYear(y int, md ...int) Time {
	md = append(md, []int{0, 0}...)
	m, d := md[0], md[1]

	if d == 0 { // 不传天数
		month := t.Month()
		y += (m) / 12                                     // 总添加年数
		m = m % 12                                        // 剩余月数
		mm := (month + m) % 12                            // 目标月份
		if v := t.Day() - DaysIn(t.Year()+y, mm); v > 0 { // 原天数 > 目标月份的最大天数
			d -= v
		}
	}

	return Time{time: t.time.AddDate(y, m, d)}
}

func (t Time) AddMonth(m int, d ...int) Time {
	return t.AddYear(0, m, append(d, 0)[0])
}

func (t Time) AddWeek(n int) Time {
	return t.AddDay(n * 7)
}

func (t Time) AddDay(d int) Time {
	return t.AddYear(0, 0, d)
}

// Add 返回 t + d 的时间
func (t Time) Add(d time.Duration) Time {
	return Time{time: t.time.Add(d)}
}

// Go 去 年[, 月[, 日]]。若不传 [天数] 则月份不会溢出。
// y 在当前年的基础上偏移 ±y 年，a[0], a[1] 指定到确切的月、日。
// 如果 [月、日] 为负数，则从最后的月、日开始偏移。
func (t Time) Go(y int, a ...int) Time {
	a = append(a, []int{0, 0}...)
	m, d := a[0], a[1]

	if m != 0 {
		mm := int(math.Min(12, math.Abs(float64(m))))
		if month := t.Month(); m > 0 {
			m = -month + mm
		} else {
			m = -month + 13 - mm
		}
	}

	if d != 0 {
		days := DaysIn(y+t.Year(), t.Month()+m)
		dd := int(math.Min(float64(days), math.Abs(float64(d))))
		if d > 0 {
			d = -t.Day() + dd
		} else {
			d = -t.Day() + (days + 1) - dd
		}
	}

	return t.AddYear(y, m, d)
}

func (t Time) GoMonth(m int, d ...int) Time {
	return t.Go(0, m, append(d, 0)[0])
}

func (t Time) GoDay(d int) Time {
	return t.Go(0, 0, d)
}

// StartYear +y 年后第 m 月 d 日的开始时间
func (t Time) StartYear(ymd ...int) Time {
	y, m, d := t.Year(), 1, 1
	if ymd != nil {
		y += ymd[0]
	}

	if len(ymd) > 1 && ymd[1] > 0 {
		m = ymd[1]
	}

	if len(ymd) > 2 && ymd[2] > 0 {
		d = ymd[2]
	}

	_t := time.Date(y, time.Month(m), d, 0, 0, 0, 0, time.Local)
	return Time{time: _t}
}

// StartMonth +m 月后第 d 日的开始时间
func (t Time) StartMonth(md ...int) Time {
	md = append(md, 0, 0)
	return t.StartYear(0, t.Month()+md[0], md[1])
}

// StartDay +d 日的开始时间
func (t Time) StartDay(d ...int) Time {
	dd := t.Day() + append(d, 0)[0]
	return t.StartYear(0, t.Month(), dd)
}

// StartWeek 本周开始时间
func (t Time) StartWeek(w ...int) Time {
	n := t.Weekday()
	if n == 0 {
		n = 7
	}

	day := t.Day() - n + 1
	day += append(w, 0)[0] * 7

	_t := time.Date(t.Year(), time.Month(t.Month()), day, 0, 0, 0, 0, time.Local)
	return Time{time: _t}
}

// InYears 返回 t 与 u 的年差
func (t Time) InYears(u Time) int {
	return int(t.Sub(u).Abs().Hours() / 8760)
}

// InDays 返回 t 与 u 的天差
func (t Time) InDays(u Time) int {
	return int(t.Sub(u).Abs().Hours() / 24)
}

// InHours 返回 t 与 u 的时差
func (t Time) InHours(u Time) int {
	return int(t.Sub(u).Abs().Hours())
}

// InMinute 返回 t 与 u 的分差
func (t Time) InMinute(u Time) int {
	return int(t.Sub(u).Abs().Minutes())
}

// InSeconds 返回 t 与 u 的秒差
func (t Time) InSeconds(u Time) int {
	return int(t.Sub(u).Abs().Seconds())
}

// Year 返回 t 的年份
func (t Time) Year() int {
	return t.time.Year()
}

// Month 返回 t 的月份
func (t Time) Month() int {
	return int(t.time.Month())
}

// Day 返回 t 的天数
func (t Time) Day() int {
	return t.time.Day()
}

// Days 返回本月份的最大天数
func (t Time) Days() int {
	// 公历一年中 4、6、9、11 是小月，都有 30 天；
	// 1、3、5、7、8、10、12 是大月，有 31 天；
	// 2 月份平年有 28 天，闰年有 29 天。
	return DaysIn(t.Year(), t.Month())
}

// YearDay 返回 t 年份中的第几天，非闰年为 1-365，闰年为 1-366。
func (t Time) YearDay() int {
	return t.time.YearDay()
}

// Weekday 返回 t 的星期
func (t Time) Weekday() int {
	return int(t.time.Weekday())
}

// Hour 返回 t 的小时
func (t Time) Hour() int {
	return t.time.Hour()
}

// Minute 返回 t 的分钟
func (t Time) Minute() int {
	return t.time.Minute()
}

// Second 返回 t 的秒数
func (t Time) Second() int {
	return t.time.Second()
}

// Milli 返回 t 的毫秒
func (t Time) Milli() int {
	return t.Nano() / 1e6
}

// Micro 返回 t 的微秒
func (t Time) Micro() int {
	return t.Nano() / 1e3
}

// Nano 返回 t 的纳秒
func (t Time) Nano() int {
	return t.time.Nanosecond()
}

// Unix 返回 t 的秒时间戳
func (t Time) Unix() int64 {
	return t.time.Unix()
}

// UnixMilli 返回 t 毫秒时间戳
func (t Time) UnixMilli() int64 {
	return t.time.UnixMilli()
}

// UnixMicro 返回 t 的微秒时间戳
func (t Time) UnixMicro() int64 {
	return t.time.UnixMicro()
}

// UnixNano 返回 t 的纳秒时间戳
func (t Time) UnixNano() int64 {
	return t.time.UnixNano()
}

// Sub 返回 t - u 的时间差
func (t Time) Sub(u Time) time.Duration {
	return t.time.Sub(u.time)
}

// Compare 比较 t 和 u。
// 如果 t 小于 u 则返回 -1；大于返回 1；等于返回 0。
func (t Time) Compare(u Time) int {
	return t.time.Compare(u.time)
}

// UTC 返回 UTC 时区的 t
func (t Time) UTC() Time {
	return Time{time: t.time.UTC()}
}

// Local 返回本地时区的 t
func (t Time) Local() Time {
	return Time{time: t.time.Local()}
}

// In 返回 loc 时区的 t
func (t Time) In(loc *time.Location) Time {
	return Time{time: t.time.In(loc)}
}

// Location 返回 t 关联的时区信息
func (t Time) Location() *time.Location {
	return t.time.Location()
}

// IsZero 返回 t 是否零时刻
func (t Time) IsZero() bool {
	return t.time.IsZero()
}

// Time 返回 t 的 time.Time
func (t Time) Time() time.Time {
	return t.time
}

// ----

// Scan 由 DB 转到 Go 时调用
func (t *Time) Scan(value any) error {
	if v, ok := value.(time.Time); ok {
		*t = Time{v}
	}
	return nil
}

// Value 由 Go 转到 DB 时调用
func (t Time) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.String(), nil
}

// MarshalJSON 将 t 转为 JSON 字符串时调用
func (t Time) MarshalJSON() ([]byte, error) {
	return []byte(`"` + t.String() + `"`), nil
}

// UnmarshalJSON 将 JSON 字符串转为 t 时调用
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	*t, err = ParseE(string(b))
	return
}

// ----

func Unix(sec, nsec int64) Time {
	return Time{time: time.Unix(sec, nsec)}
}

func UnixMilli(msec int64) Time {
	return Time{time: time.UnixMilli(msec)}
}

func UnixMicro(usec int64) Time {
	return Time{time: time.UnixMicro(usec)}
}

func UnixNano(nsec int64) Time {
	return Time{time: time.Unix(0, nsec)}
}

// Since 返回自 t 以来经过时间。它是 Now().Sub(t) 的简写。
func Since(t Time) time.Duration {
	return time.Since(t.time)
}

// IsLeap 返回 year 是否闰年
func IsLeap(year int) bool {
	return year%4 == 0 && (year%100 != 0 || year%400 == 0)
}

// DaysIn 返回 y 年中 m 月的最大天数
func DaysIn(y, m int) int {
	if m == 2 && IsLeap(y) {
		return 29
	}
	return days[m]
}
