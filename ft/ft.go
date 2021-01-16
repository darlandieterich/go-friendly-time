package ft

import "time"

type TypeFT struct {
	t time.Time
	l *time.Location
}

func (t TypeFT) SubtractDate(dateEnd time.Time) time.Time {
	return t.t.AddDate(dateEnd.Year(), int(dateEnd.Month()), dateEnd.Day())
}

func (t TypeFT) SubtractTwoDates(dateBegin time.Time, dateEnd time.Time) time.Time {
	return dateBegin.AddDate(dateEnd.Year(), int(dateEnd.Month()), dateEnd.Day())
}
