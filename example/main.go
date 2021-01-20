package main

import (
	"log"
	"time"

	//import parent package/directory
	ft ".."
)

func main() {
	ret := ft.TypeFT{
		Time:   time.Now(),
		Locale: ft.GetLocal(ft.AmericaSaoPaulo),
	}
	log.Println(ret.GetTimeNowString(ft.FullDateTime))
	ret.Locale = ft.GetLocal(ft.PacificFiji)
	log.Println(ret.GetTimeNowStruct())
}
