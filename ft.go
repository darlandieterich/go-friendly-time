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
	CalcDateUp int = iota
	CalcDateDown
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
func (t TypeFT) CalcDate(dateEnd time.Time) time.Time {
	return t.Time.AddDate(dateEnd.Year(), int(dateEnd.Month()), dateEnd.Day())
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
