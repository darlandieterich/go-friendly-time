package main

import (
	"log"
	ft "github.com/darlandieterich/go-friendly-time"
)

func main() {
	/log.Println("oi")
	log.Println(time.Now().Format(ft.Hour))
}
