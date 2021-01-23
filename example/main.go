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
		Format: ft.DTFullDateTime,
	})
	//Get Time now in string
	fmt.Println(ret.GetTimeNowString())
	//Show the time in personalized format with array of string normalized constants
	fmt.Println(ret.FormatDateTime(time.Now(), []string{ft.DTHour, ft.DTColon, ft.DTMinute, ft.DTColon, ft.DTSecond}))
	//Up 10 years
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
