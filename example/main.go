package main

import (
	"fmt"
	"time"

	//import parent package/directory
	ft ".."
)

func main() {
	ret := ft.New(&ft.TypeFT{
		Time:   time.Now(),
		Locale: ft.GetLocal(ft.TZAmericaSaoPaulo),
		Format: ft.DTFullDateTime,
	})

	fmt.Println(ret.GetTimeNowString())
	//ret.Locale = ft.GetLocal(ft.TZPacificFiji)
	//fmt.Println(ret.GetTimeNowStruct())
	//ret.Format = ft.DTTime
	//fmt.Println(ret.GetTimeNowString())
	fmt.Println(ft.FormatDateTime(time.Now(), []string{ft.DTHour, ft.DTColon, ft.DTMinute, ft.DTColon, ft.DTSecond}))
}
