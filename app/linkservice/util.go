package linkservice

import "time"

func defaultLastClickedDate() time.Time {
	return time.Date(1993, time.April, 30, 12, 0, 0, 0, time.UTC)
}
