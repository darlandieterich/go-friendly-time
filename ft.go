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

// CalcTime calculate the date sum or subtract dates
func (t TypeFT) CalcTime(duration int, cycle int) time.Time {
	tm := t.Time.In(t.Locale)
	switch cycle {
	case DTCycleNanosecond:
		return tm.Add(time.Duration(duration) * time.Nanosecond)
	case DTCycleMicrosecond:
		return tm.Add(time.Duration(duration) * time.Microsecond)
	case DTCycleMillisecond:
		return tm.Add(time.Duration(duration) * time.Millisecond)
	case DTCycleSecond:
		return tm.Add(time.Duration(duration) * time.Second)
	case DTCycleMinute:
		return tm.Add(time.Duration(duration) * time.Minute)
	case DTCycleHour:
		return tm.Add(time.Duration(duration) * time.Hour)
	case DTCycleDay:
		return tm.AddDate(0, 0, duration)
	case DTCycleMonth:
		return tm.AddDate(0, duration, 0)
	case DTCycleYear:
		return tm.AddDate(duration, 0, 0)
	default:
		return tm
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
