package scheduler

import (
	"github.com/ycombinator/courtgrabber/chromedriver"
)

const (
	loginURL = "https://www.avac.playtennisconnect.com/pages/memberpage_login.cfm"
	dayURL   = "https://www.avac.playtennisconnect.com/pages/scheduler/index.cfm?cday=%d&cmonth=%d&cyear=%d&ReservationTypeID=1&layout=timeline&courtgroupid="
	slotURL  = "https://www.avac.playtennisconnect.com/pages/scheduler/index.cfm?ReservationID=0&slot=%s&ReservationTypeID=1&SortOrder=%d&courtgroupid="
)

func Close() error {
	return chromedriver.Close()
}
