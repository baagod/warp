# beam
golang 时间包

```go
now := beam.Now() // 现在

fmt.Println("---- 开始时间 ----")
fmt.Println("去年开始时间:", now.StartYear(-1))
fmt.Println("今年开始时间:", now.StartYear())
fmt.Println("明年开始时间:", now.StartYear(1))

fmt.Println("上月开始时间:", now.StartMonth(-1))
fmt.Println("本月开始时间:", now.StartMonth())
fmt.Println("下月开始时间:", now.StartMonth(1))

fmt.Println("上周开始时间:", now.StartWeek(-1))
fmt.Println("本周开始时间:", now.StartWeek())
fmt.Println("下周开始时间:", now.StartWeek(1))

fmt.Println("昨天开始时间:", now.StartDay(-1))
fmt.Println("今天开始时间:", now.StartDay())
fmt.Println("明天开始时间:", now.StartDay(1))

fmt.Println("---- 偏移时间 ----")
fmt.Println("今天:", now)
fmt.Println("偏移 ±年月日，若不传天数则月份不会溢出:", now.Add(1, -2, 3))
fmt.Println("只偏移 ±月日:", now.AddMonth(1, -2))
fmt.Println("只偏移 ±日:", now.AddDay(1))
fmt.Println("偏移 ±周:", now.AddWeek(-1))
fmt.Println("去年11月最后第3天，月份不会溢出:", now.Go(-1, 11, -3))
fmt.Println("今年3月11日:", now.GoMonth(3, 11))
fmt.Println("本月10号:", now.GoDay(10))

fmt.Println("---- 时间相差 ----")
fmt.Println("相差几年:", now.InYears(beam.MustParse("2024-03-19")))
fmt.Println("相差几天:", now.InDays(beam.MustParse("2024-03-19")))
fmt.Println("相差几时:", now.InHours(beam.MustParse("2024-03-19")))
fmt.Println("相差几分:", now.InMinute(beam.MustParse("2024-03-19")))
fmt.Println("相差几秒:", now.InSeconds(beam.MustParse("2024-03-19")))

fmt.Println("---- 时间转字符串 ----")
fmt.Println("毫秒字符串:", now.String())
fmt.Println("日期和时间:", now.DateTime())
fmt.Println("日期字符串:", now.DateOnly())
```
