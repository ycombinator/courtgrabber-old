package scheduler

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/fedesog/webdriver"
	"github.com/ycombinator/courtgrabber/chromedriver"
)

type Court struct {
	number    int
	startTime time.Time
}

func (c *Court) Book() error {
	return nil
}

func FindFreeCourt(startTime time.Time, duration time.Duration) (*Court, error) {
	availableCourtNumbers, err := findAvailableCourts(startTime)
	if err != nil {
		return nil, ErrFindCourt{err}
	}

	nextSlotCourtNumbers, err := findAvailableCourts(startTime.Add(30 * time.Minute))
	if err != nil {
		return nil, ErrFindCourt{err}
	}
	availableCourtNumbers = intersection(availableCourtNumbers, nextSlotCourtNumbers)

	nextSlotCourtNumbers, err = findAvailableCourts(startTime.Add(30 * time.Minute))
	if err != nil {
		return nil, ErrFindCourt{err}
	}
	availableCourtNumbers = intersection(availableCourtNumbers, nextSlotCourtNumbers)

	if len(availableCourtNumbers) == 0 {
		return nil, nil
	}

	fmt.Println(availableCourtNumbers)

	firstAvailableCourtNumber := availableCourtNumbers[0]
	c := Court{
		number:    firstAvailableCourtNumber,
		startTime: startTime,
	}
	return &c, nil
}

func findAvailableCourts(timeSlot time.Time) ([]int, error) {
	session, err := chromedriver.Session()
	if err != nil {
		return nil, err
	}

	var availableCourtNumbers []int

	timeStr := timeSlot.Format("3:04 PM")
	rows, err := session.FindElements(webdriver.CSS_Selector, "table.table > tbody > tr")
	for _, row := range rows {
		cols, err := row.FindElements(webdriver.TagName, "td")
		if err != nil {
			return nil, err
		}

		firstCol := cols[0]
		txt, err := firstCol.Text()
		if err != nil {
			return nil, err
		}

		if txt != timeStr {
			continue
		}

		slotCols := cols[1:5]
		for _, slotCol := range slotCols {
			slotLink, err := slotCol.FindElement(webdriver.TagName, "a")
			if err != nil {
				continue
			}

			txt, err = slotLink.Text()

			if strings.ToLower(txt) != "available" {
				continue
			}

			link, err := slotLink.GetAttribute("href")
			if err != nil {
				return nil, err
			}

			courtNum, err := parseCourtNum(link)
			if err != nil {
				return nil, err
			}

			if courtNum > 9 {
				continue
			}

			availableCourtNumbers = append(availableCourtNumbers, courtNum)
		}
	}

	return availableCourtNumbers, nil
}

func parseCourtNum(link string) (int, error) {
	u, err := url.Parse(link)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(u.Query().Get("SortOrder"))
}

func intersection(s1, s2 []int) []int {
	var result []int
	for _, e := range s1 {
		if inSlice(s2, e) {
			result = append(result, e)
		}
	}
	return result
}

func inSlice(haystack []int, needle int) bool {
	for _, e := range haystack {
		if needle == e {
			return true
		}
	}
	return false
}
