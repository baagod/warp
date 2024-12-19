# warp

golang 时间包

## 安装

```shell
go get -u github.com/baagod/thru@main
```

## 使用

### 创建时间

使用 `thru` 包中提供的函数创建时间对象。

```go
now := thru.Now()                              // 现在。通过快捷函数创建
thru.New(time.Now())                           // 通过 time.Time 创建
thru.Unix(now.Unix())                          // 通过秒时间戳创建
thru.Unix(now.Unix(3))                         // 通过毫秒时间戳创建
thru.Parse("2023-03-19")                       // 通过解析自动字符串创建
thru.ParseByLayout("2006-01-02", "2023-03-19") // 通过布局格式创建
```

### 偏移时间

主要使用 `AddYear(ymd ...int)` 方法进行操作。该函数跟 `time.AddDate()` 相似，但是在不传入参数 d (要偏移的天数) 时，年、月不会溢出。
`y, m, d` 可以为负数。

现在假设当前时间是：`2023-01-31 13:14:15 +0800 CST` ( 中国标准时间 )

```go
now.AddYear() // 添加一年
now.AddMonth(-2, 3)   // 在当前时间上偏移 -2月3天
now.AddDay(3)         // 三天后
now.Add(time.Hour)    // 1小时后
```

指定 `d` 参数时可能会导致日期溢出，例如：

```go
now.AddYear(1, 1, 0)
```

该函数在指定 `d` 参数时的行为完全和 `time.AddDate()` 相同，因此返回 `2024-03-02 13:14:15`。这导致我们以为结果和预期不符，这就是日期溢出。

让我来尝试解释一下，这是因为 2024 年 2 月一共有 29 天。当我们预期从 2023-01-31 移动到 2024-02-31 时，实际上天数已经溢出 2
天，结果就是它跑到 2024-02-29 时发现后面还有 2 天要继续增加，最终返回 2024-03-02。

### 选择时间

主要使用 `Go(y int, md ...int)` 方法进行操作。该函数偏移 `±y` 年并选择 `m` 月 `d` 日，如果 `m, d`
为负数，则从最后的月、日开始选择，并且不会溢出。例如：

- `Go(1, 13, 1)` 会返回 `2024-12-31 13:14:15`。
- `Go(1, 2, 31)` 返回二月最后一天 `2024-02-29 13:14:15`。

```go
now.Go(-1, 11, -3) // 去年11月最后三天
now.GoMonth(-1, 2) // 年末第二天
now.GoDay(10)      // 本月10号
```

为了扩展，我还添加了一个 `now.GoYear(y int, md ...int)` 方法。跟 `Go()` 一样，但 `y` 指定为确切年份而非偏移。例如：

```go
now.GoYear(2024, 03, 15) // 返回 2024-03-15 13:14:15
```

但是这就跟直接创建差不多了？有用吗？也许吧，我不知道。

### 开始时间、结束时间

我们经常需要定位要时间边界，因此定义开始时间和结束时间的相关函数。

开始时间主要使用 `Start(ymd ...int)` 进行操作，它和 `AddYear()` 类似，但是返回开始时间。

```go
now.Start()         // 今年开始时间 2023-01-01 00:00:00
now.Start(-1)       // 去年开始时间 2022-01-01 00:00:00
now.Start(1)        // 明年开始时间 2024-01-01 00:00:00
now.Start(1, 1, 1)  // 偏移后的开始时间 2024-03-03 00:00:00

now.StartWeek(-1) // 上周开始时间
now.StartWeek()   // 本周开始时间
now.StartWeek(1)  // 下周开始时间
now.StartWeek(2)  // 两周后的星期一开始时间
```

另外还提供方便方法，和 `Start(ymd ...int)` 一样：

- `StartMonth(md ...int)`: 从 `m` 开始。
- `StartDay(d int)`: 从 `d` 开始。

还有与开始时间 `Start()` 系列方法对应的所有结束时间函数: `End`, `EndMoth`, `EndDay`, `EndWeek`。

### 获取时间

获得各种时间的函数，例如要获得年、月、日、时、分、秒等。

```go
Year()    // 今年
YearDay() // 返回年份中的日期，非闰年的范围为 [1, 365]，闰年的范围为 [1, 366]。
Month()   // 本月
Day()     // 返回本月的第几天
Days()    // 本月最大天数
Hour()    // 返回小时 [0, 23]
Minute()  // 返回分钟 [0, 59]
Second()  // 返回秒 [0, 59]
Second(3) // 毫秒 (3位)
Second(6) // 微秒 (6位)
Second(9) // 纳秒 (9位)

Time.Unix()  // 秒时间戳 (10位)
Time.Unix(3) // 毫秒时间戳 (13位)
Time.Unix(6) // 微秒时间戳 (16位)
Time.Unix(9) // 纳秒时间戳 (19位)

UTC() // 返回 UTC 时间
Local() // 返回本地时间
In(loc *time.Location) // 返回指定的 loc 时间
Location() // 返回时区信息

Time()    // 返回 time.Time
ZeroOr(u Time) // 使用 u 代替零时 (isZero() is true) 时间
thru.DaysIn(2024, 2) // 返回指定年的月份最大天数
```

### 比较时间

比较时间差主要由以下两个函数返回，可指定要比较的时间单位。

- `DiffIn(u Time, uint string)`
- `DiffAbsIn(u Time, uint string)`

```go
DiffIn('y') // 返回年差
DiffIn('M') // 返回月差
DiffIn('d') // 返回天差
DiffIn('h') // 返回时差
DiffIn('m') // 返回分差
DiffIn('s') // 返回秒差
```

另外补充具有和 `time` 包中相同的行为的函数：

- `IsZero() bool`
- `Before(u Time) bool`
- `After(u Time) bool`
- `Equal(u Time) bool`
- `Compare(u Time) int`
- `Sub(u Time) time.Duration` 返回 `t - u` 的时间差
- `thru.Since(Time)` 返回自 `t` 以来经过的时间。它是 `Now().Sub(t)` 的简写。
- `thru.Until(Time)` 返回直到 `t` 的持续时间。它是 `t.Sub(Now())` 的简写。

还可以使用 `thru.IsLeap(int)` 函数判断是否闰年。

### 时间字符串

```go
now.String()   // 返回日期时间字符串: 2023-03-19 15:05:27
now.Format("2006年01年02日") // 指定时间布局格式返回
```
