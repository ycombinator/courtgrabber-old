package scheduler

import "time"

type Config struct {
	Username string
	Password string

	StartTime time.Time
	Duration  time.Duration
}

func DefaultConfig() Config {
	return Config{
		StartTime: defaultStartTime(),
		Duration:  90 * time.Minute,
	}
}

// today + 3 days @ 18:30 Pacific Time
func defaultStartTime() time.Time {
	t := time.Now()
	t = t.Add(3 * 24 * time.Hour)

	l, _ := time.LoadLocation("America/Los_Angeles")

	return time.Date(
		t.Year(),
		t.Month(),
		t.Day(),
		18,
		30,
		0,
		0,
		l,
	)
}
