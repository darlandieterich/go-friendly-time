# go-friendly-time
Friendly helper to work with package time

## example usage

```go
package main

import (
	"fmt"
	"time"

	ft "github.com/darlandieterich/go-friendly-time"
)

func main() {
	//New declaration
	ret := ft.New(&ft.TypeFT{
		Time:   time.Now(),
		Locale: ft.GetLocal(ft.TZAmericaSaoPaulo),
		Format: ft.DTFullDateTime,
	})
	//Get Time now in string
	fmt.Println(ret.GetTimeNowString())
	//Show the time in personalized format with array of string normalized constants
	fmt.Println(ret.FormatDateTime(time.Now(), []string{ft.DTHour, ft.DTColon, ft.DTMinute, ft.DTColon, ft.DTSecond}))
	//Up to 10 years
	fmt.Println("Years: ", ret.CalcTime(10, ft.DTCycleYear))
	//Change the Location
	ret.Locale = ft.GetLocal(ft.TZPacificFiji)
	//Show the location in Time struct
	fmt.Println(ret.GetTimeNowStruct())
	//Change the format to only time
	ret.Format = ft.DTTime
	//Collect the new time
	fmt.Println(ret.GetTimeNowString())
}
```

## Constant declarations
### Format
File: ``ft_datetime.go``

Types:
```go
	DTFormatHour          = "15"
	DTFormatMinute        = "04"
	DTFormatSecond        = "05"
	DTFormatDay           = "02"
	DTFormatMonth         = "01"
	DTFormatYear          = "2006"
	DTFormatTime          = "15:04:05"
	DTFormatFullDate      = "2006-01-02"
	DTFormatFullDateTime  = "2006-01-02 15:04:05"
	DTFormatDateFormatYMD = "2006-01-02"
	DTFormatDateFormatDMY = "02-01-2006"
	DTFormatDateFormatMDY = "01-02-2006"
	DTFormatBlank         = " "
	DTFormatDash          = "-"
	DTFormatColon         = ":"
	DTFormatDot           = "."
	DTFormatComma         = ","
	DTFormatSlash         = "/"
```
### Cycle time
File: ``ft_datetime.go``

Types:
```go
	DTCycleNanosecond int = iota
	DTCycleMicrosecond
	DTCycleMillisecond
	DTCycleSecond
	DTCycleMinute
	DTCycleHour
	DTCycleDay
	DTCycleMonth
	DTCycleYear
```
### Timezone Location
File: ``ft_timezone.go``

Types:
```go
	TZAfricaAbidjan                  = "Africa/Abidjan"
	TZAfricaAccra                    = "Africa/Accra"
	TZAfricaAddisAbaba               = "Africa/Addis_Ababa"
	TZAfricaAlgiers                  = "Africa/Algiers"
	TZAfricaAsmara                   = "Africa/Asmara"
	TZAfricaBamako                   = "Africa/Bamako"
	//.....more in ft_timezone.go
```
