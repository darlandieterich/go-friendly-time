package main

import (
	"fmt"
	"time"

	//import parent package/directory
	ft ".."
)

func main() {
	//New declaration
	ret := ft.New(&ft.TypeFT{
		Time:   time.Now(),
		Locale: ft.GetLocal(ft.TZAmericaSaoPaulo),
		Format: ft.DTFormatFullDateTime,
	})
	//Get Time now in string
	fmt.Println(ret.GetTimeNowString())
	//Show the time in personalized format with array of string normalized constants
	fmt.Println(ret.FormatDateTime(time.Now(), []string{ft.DTFormatHour, ft.DTFormatColon, ft.DTFormatMinute, ft.DTFormatColon, ft.DTFormatSecond}))
	//Up to 10 years
	fmt.Println("Years: ", ret.CalcTime(10, ft.DTCycleYear))
	//Change the Location
	ret.Locale = ft.GetLocal(ft.TZPacificFiji)
	//Show the location in Time struct
	fmt.Println(ret.GetTimeNowStruct())
	//Change the format to only time
	ret.Format = ft.DTFormatTime
	//Collect the new time
	fmt.Println(ret.GetTimeNowString())
}
