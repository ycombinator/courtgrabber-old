package scheduler

import (
	"fmt"
	"time"

	"github.com/ycombinator/courtgrabber/chromedriver"
)

func SelectDay(day time.Time) error {
	session, err := chromedriver.Session()
	if err != nil {
		return ErrDay{err}
	}

	err = session.Url(makeDayURL(dayURL, day))
	if err != nil {
		return ErrDay{err}
	}

	return nil
}

func makeDayURL(dayURL string, day time.Time) string {
	return fmt.Sprintf(dayURL, day.Day(), day.Month(), day.Year())
}
