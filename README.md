# beam

golang 时间包

## 安装

```shell
go get github.com/baa-god/beam
```

## 使用

```go
now := beam.Now() // 现在: 2023-03-19 15:05:27.792
```

### 开始时间

```go
now.StartYear(-1) // 去年开始时间: 2022-01-01 00:00:00.000
now.StartYear())  // 今年开始时间: 2023-01-01 00:00:00.000
now.StartYear(1)  // 明年开始时间: 2024-01-01 00:00:00.000

now.StartMonth(-1) // 上月开始时间: 2023-02-01 00:00:00.000
now.StartMonth()   // 本月开始时间: 2023-03-01 00:00:00.000
now.StartMonth(1)  // 下月开始时间: 2023-04-01 00:00:00.000

now.StartWeek(-1) // 上周开始时间: 2023-03-06 00:00:00.000
now.StartWeek()   // 本周开始时间: 2023-03-13 00:00:00.000
now.StartWeek(1)  // 下周开始时间: 2023-03-20 00:00:00.000

now.StartDay(-1) // 昨天开始时间: 2023-03-18 00:00:00.000
now.StartDay()   // 今天开始时间: 2023-03-19 00:00:00.000
now.StartDay(1)  // 明天开始时间: 2023-03-20 00:00:00.000
```

### 偏移时间

```go
// 添加年月日，若不传天数则月份不会溢出。
now.Add(1, -2, 3) // 此时的明年前两个月+3天: 2024-01-22 15:05:27.792
now.AddMonth(-2, 3) // 只偏移 ±月日，前两个月+3天: 2023-01-22 15:05:27.792
now.AddDay(3) // 只偏移 ±日，三日后: 2023-03-22 15:05:27.792
now.AddWeek(-1) // 偏移 ±周，上周: 2023-03-12 15:05:27.792
now.AddDuration(time.Hour) // 1小时后: 2023-03-19 16:05:27.792

// 去指定的年月日，若不传天数则月份不会溢出。
now.Go(-1. 11, -3) // 去年11月倒数第三天: 2022-11-28 15:05:27.792
now.GoMonth(-3, 2) // 今年倒数第3月(10月)2号: 2023-10-02 15:05:27.792
now.GoDay(10) // 本月10号: 2023-03-10 15:05:27.792
```

### 时间相差

```go
now.InYears(beam.MustParse("2024-03-19"))   // 相差几年: 1
now.InDays(beam.MustParse("2024-03-19"))    // 相差几天: 365
now.InHours(beam.MustParse("2024-03-19"))   // 相差几时: 8768
now.InMinute(beam.MustParse("2024-03-19"))  // 相差几分: 526134
now.InSeconds(beam.MustParse("2024-03-19")) // 相差几秒: 31568072
```

### 字符串值

```go
now.String()   // 毫秒字符串: 2023-03-19 15:05:27.792
now.DateTime() // 日期和时间: 2023-03-19 15:05:27
now.DateOnly() // 日期字符串: 2023-03-19
```
