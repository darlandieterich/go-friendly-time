package main

import (
	"log"
	"time"

	ft "./model"
)

func main() {
	log.Println(time.Now().Format(ft.FullDateTime))
}
