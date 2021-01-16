package ft

import "time"

type TypeFT struct {
	t time.Time
	l *time.Location
}

const (
	Hour         string = "15"
	Minute       string = "04"
	Second       string = "05"
	Day          string = "02"
	Month        string = "01"
	Year         string = "2006"
	FullDate     string = "2006-01-02"
	FullDateTime string = "2006-01-02 15:04:05"
)

func (t TypeFT) SubtractDate(dateBegin time.Time, dateEnd time.Time) {

}
