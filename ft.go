package ft

import (
	"time"
)

// TypeFT struc of time
type TypeFT struct {
	Time   time.Time
	Locale *time.Location
	Format string
}

// CalcDate calculate the date sum or subtract dates
func (t TypeFT) CalcDate(dateEnd time.Time) time.Time {
	return t.Time.AddDate(dateEnd.Year(), int(dateEnd.Month()), dateEnd.Day())
}

// GetTimeNowString return time now in string format
func (t TypeFT) GetTimeNowString() string {
	return time.Now().In(t.Locale).Format(t.Format)
}

// GetTimeNowStruct return time now in struct
func (t TypeFT) GetTimeNowStruct() time.Time {
	return time.Now().In(t.Locale)
}
