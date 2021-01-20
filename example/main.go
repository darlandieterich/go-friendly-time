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
		Locale: ft.GetLocal(ft.AsiaTokyo),
	}
	log.Println(ret.Time.In(ret.Locale).Format(ft.FullDateTime))
}
