package main

import (
	"fmt"
	"time"

	//import parent package/directory
	ft ".."
)

func main() {
	ret := ft.TypeFT{
		Time:   time.Now(),
		Locale: ft.GetLocal(ft.AmericaSaoPaulo),
		Format: ft.FullDateTime,
	}
	fmt.Println(ret.GetTimeNowString())
	ret.Locale = ft.GetLocal(ft.PacificFiji)
	fmt.Println(ret.GetTimeNowStruct())
	ret.Format = ft.Time
	fmt.Println(ret.GetTimeNowString())
}
