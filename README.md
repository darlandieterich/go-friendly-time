# go-friendly-time
Friendly helper to work with package time

# example usage

```go
func main() {
	ret := ft.TypeFT{
		Time:   time.Now(),
		Locale: ft.GetLocal(ft.AmericaSaoPaulo),
	}
	log.Println(ret.GetTimeNowString(ft.FullDateTime))
	ret.Locale = ft.GetLocal(ft.PacificFiji)
  log.Println(ret.GetTimeNowStruct())
  ret.Format = ft.Time
	fmt.Println(ret.GetTimeNowString())
}
```