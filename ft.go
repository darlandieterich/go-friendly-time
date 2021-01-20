package ft

import "time"

// TypeFT struc of time
type TypeFT struct {
	Time   time.Time
	Locale *time.Location
}

// CalcDate calculate the date sum or subtract dates
func (t TypeFT) CalcDate(dateEnd time.Time) time.Time {
	return t.Time.AddDate(dateEnd.Year(), int(dateEnd.Month()), dateEnd.Day())
}

// GetTimeNowString return time now in string format
func (t TypeFT) GetTimeNowString(format string) string {
	return time.Now().In(t.Locale).Format(format)
}

// GetTimeNowStruct return time now in struct format
func (t TypeFT) GetTimeNowStruct(format string) time.Time {
	return time.Now().In(t.Locale)
}
