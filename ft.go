package ft

import "time"

type TypeFT struct {
	Time   time.Time
	Locale *time.Location
}

func (t TypeFT) CalcDate(dateEnd time.Time) time.Time {
	return t.Time.AddDate(dateEnd.Year(), int(dateEnd.Month()), dateEnd.Day())
}

//get time now
