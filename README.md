# beam

golang 时间包

## 安装

```shell
go get github.com/baa-god/beam
```

## 使用

### 创建时间

```go
now := beam.Now() // 现在: 2023-03-19 15:05:27.792

beam.New(time.Now())            // 通过 time.Time 创建
beam.Unix(now.Unix(), 0)        // 通过秒时间戳
beam.UnixMilli(now.UnixMilli()) // 通过毫秒时间戳
beam.UnixMicro(now.UnixMicro()) // 通过微秒时间戳
beam.UnixNano(now.UnixNano())   // 通过纳秒时间戳
beam.Prase("2023-03-19")        // 通过解析时间字符串
beam.Layout("2006-01-02", "2023-03-19") // 通过布局创建
```

### 开始时间

```go
// 加 y 年后第 m 月 d 日的开始时间
now.StartYear(-1, 2, 1) // 去年2月1日: 2022-02-01 00:00:00.000
now.StartYear(-1)       // 去年开始时间: 2022-01-01 00:00:00.000
now.StartYear()         // 今年开始时间: 2023-01-01 00:00:00.000
now.StartYear(1)        // 明年开始时间: 2024-01-01 00:00:00.000

// 加 m 月后第 d 日的开始时间
now.StartMonth(-1, 2)   // 上月2号: 2023-02-02 00:00:00.000
now.StartMonth(-1)      // 上月开始时间: 2023-02-01 00:00:00.000
now.StartMonth()        // 本月开始时间: 2023-03-01 00:00:00.000
now.StartMonth(1)       // 下月开始时间: 2023-04-01 00:00:00.000

now.StartDay(-1)        // 昨天开始时间: 2023-03-18 00:00:00.000
now.StartDay()          // 今天开始时间: 2023-03-19 00:00:00.000
now.StartDay(1)         // 明天开始时间: 2023-03-20 00:00:00.000

now.StartWeek(-1)       // 上周开始时间: 2023-03-06 00:00:00.000
now.StartWeek()         // 本周开始时间: 2023-03-13 00:00:00.000
now.StartWeek(1)        // 下周开始时间: 2023-03-20 00:00:00.000
```

### 偏移时间

```go
// 添加年月日，若不传天数则月份不会溢出。
now.AddYear(1, -2, 3) // 此时的明年-2月+3天: 2024-01-22 15:05:27.792
now.AddMonth(-2, 3)   // 前两个月+3天: 2023-01-22 15:05:27.792
now.AddDay(3)         // 三天后: 2023-03-22 15:05:27.792
now.AddWeek(-1)       // 上周: 2023-03-12 15:05:27.792
now.Add(time.Hour)    // 1小时后: 2023-03-19 16:05:27.792

// 去 t.Year()+y 年 a[0] 月 a[1] 日。月、日均不溢出。
now.Go(-1, 11, -3) // 去年11月倒数第三天: 2022-11-28 15:05:27.792
now.GoMonth(-3, 2) // 今年倒数第3月(10月)2号: 2023-10-02 15:05:27.792
now.GoDay(10)      // 本月10号: 2023-03-10 15:05:27.792
```

### 获取时间

```go
now.Year()    // 今年: 2023
now.YearDay() // 今年的第几天，平年 1-365，闰年 1-366: 78
now.Month()   // 本月: 3
now.Day()     // 今天: 19
now.Days()    // 本月最大天数: 31
now.Hour()    // 此时: 15
now.Minute()  // 此分: 5
now.Second()  // 此秒: 27
now.Milli()   // 毫秒: 792
now.Micro()   // 微秒: 792000
now.Nano()    // 纳秒: 792000000
beam.DaysIn(2024, 2) // 2024 年 2 月的最大天数: 29
now.Time()           // 返回 now 的 time.Time
```

### 时间戳

```go
now.Unix()      // 秒时间戳: 1679209527
now.UnixMilli() // 毫秒时间戳: 1679209527792
now.UnixMicro() // 微秒时间戳: 1679209527792000
now.UnixNano()  // 纳秒时间戳: 1679209527792000000
```

### 时间差

```go
now.InYears(beam.Parse("2024-03-19"))          // 相差几年: 1
now.InDays(beam.Parse("2024-03-19"))           // 相差几天: 365
now.InHours(beam.Parse("2024-03-19"))          // 相差几时: 8768
now.InMinute(beam.Parse("2024-03-19"))         // 相差几分: 526134
now.InSeconds(beam.Parse("2024-03-19"))        // 相差几秒: 31568072
now.Sub(beam.Parse("2023-03-19 16:04:27.792")) // 时间差: 1h0m0s
beam.Since(now) // now 到现在所经过的持续时间: 0s
```

### 判断时间

```go
IsLeap(2023) // 是否闰年: false
now.IsZero() // 是否零时，即第1年1月1日 00:00:00 UTC: false
now.Compare(beam.Parse("2024")) // 小于指定的日期返回 -1，大于返回 1，等于返回 0: -1
now.Before(beam.Parse("2024"))  // 在指定时间之前 (t < u): true 
now.After(beam.Parse("2024"))   // 在指定时间之后 (t > u): false
now.Equal(beam.Parse("2024"))   // 是否相等: false
```

### 时间字符串

```go
now.String()   // 毫秒字符串: 2023-03-19 15:05:27.792
now.DateTime() // 日期和时间: 2023-03-19 15:05:27
now.DateOnly() // 日期字符串: 2023-03-19
now.Format("2006年01年02日") // 格式化字符串: 2023年03月19日
```
