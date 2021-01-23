package ft

const (
	DTFormatHour          = "15"
	DTFormatMinute        = "04"
	DTFormatSecond        = "05"
	DTFormatDay           = "02"
	DTFormatMonth         = "01"
	DTFormatYear          = "2006"
	DTFormatTime          = "15:04:05"
	DTFormatFullDate      = "2006-01-02"
	DTFormatFullDateTime  = "2006-01-02 15:04:05"
	DTFormatDateFormatYMD = "2006-01-02"
	DTFormatDateFormatDMY = "02-01-2006"
	DTFormatDateFormatMDY = "01-02-2006"
	DTFormatBlank         = " "
	DTFormatDash          = "-"
	DTFormatColon         = ":"
	DTFormatDot           = "."
	DTFormatComma         = ","
	DTFormatSlash         = "/"
)

const (
	DTCycleNanosecond int = iota
	DTCycleMicrosecond
	DTCycleMillisecond
	DTCycleSecond
	DTCycleMinute
	DTCycleHour
	DTCycleDay
	DTCycleMonth
	DTCycleYear
)
