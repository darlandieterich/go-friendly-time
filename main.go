package main

import (
	"log"
	"time"

	ft "./dt"
)

func main() {
	log.Println(time.Now().Format(ft.Hour))
}
