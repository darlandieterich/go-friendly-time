package ft

import (
	"strings"
	"time"
)

// TypeFT struc of time
type TypeFT struct {
	Time   time.Time
	Locale *time.Location
	Format string
}

const (
	ConsumeUp int = iota
	ConsumeDown
)

func New(options *TypeFT) TypeFT {
	if options != nil {
		return *options
	}
	return TypeFT{
		Time:   time.Now(),
		Locale: GetLocal(TZLocal),
		Format: DTFullDateTime,
	}
}

// CalcDate calculate the date sum or subtract dates
func (t TypeFT) CalcDate(duration int, expression int) time.Time {
	switch expression {
	case DTCycleNanosecond:
		return t.Time.Add(time.Duration(duration) * time.Nanosecond)
	case DTCycleMicrosecond:
		return t.Time.Add(time.Duration(duration) * time.Microsecond)
	case DTCycleMillisecond:
		return t.Time.Add(time.Duration(duration) * time.Millisecond)
	case DTCycleSecond:
		return t.Time.Add(time.Duration(duration) * time.Second)
	case DTCycleMinute:
		return t.Time.Add(time.Duration(duration) * time.Minute)
	case DTCycleHour:
		return t.Time.Add(time.Duration(duration) * time.Hour)
	case DTCycleDay:
		return t.Time.AddDate(0, 0, duration)
	case DTCycleMonth:
		return t.Time.AddDate(0, duration, 0)
	case DTCycleYear:
		return t.Time.AddDate(duration, 0, 0)
	default:
		return t.Time
	}

}

func FormatDateTime(datetime time.Time, format []string) string {
	return datetime.Format(strings.Join(format[:], ""))
}

func (t TypeFT) FormatDateTime(datetime time.Time, format []string) string {
	return t.Time.In(t.Locale).Format(strings.Join(format[:], ""))
}

// GetTimeNowString return time now in string format
func (t TypeFT) GetTimeNowString() string {
	return time.Now().In(t.Locale).Format(t.Format)
}

// GetTimeNowStruct return time now in struct
func (t TypeFT) GetTimeNowStruct() time.Time {
	return time.Now().In(t.Locale)
}
