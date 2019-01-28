package main

import "time"

const dateFormat = "2006-01-02"

// TimeFrame holds information between which dates data can be crawled
type TimeFrame struct {
	since string
	until string
}

// IsValid checks that:
// 		since is mot empty
//		until is not empty
//		since is in the given time format
//		until is in the given time format
//		since <= until
func (t TimeFrame) IsValid() bool {
	if t.since == "" || t.until == "" {
		return false
	}
	sinceTime, err := time.Parse(dateFormat, t.since)
	if err != nil {
		return false
	}
	untilTime, err := time.Parse(dateFormat, t.until)
	if err != nil {
		return false
	}

	return sinceTime.Before(untilTime)
}

// TimeFrameFromDays returns a TimeFrame by counting back the days in time beginning from today
func TimeFrameFromDays(days int) TimeFrame {
	if days <= 0 {
		return TimeFrame{}
	}
	sinceTime := time.Now().AddDate(0, 0, -(days - 1)).Format(dateFormat)
	untilTime := time.Now().AddDate(0, 0, 1).Format(dateFormat) // in order to include tweets from today, we have to crawl from "tomorrow"

	return TimeFrame{since: sinceTime, until: untilTime}
}

// TimeFrameFromSince returns a TimeFrame from today until the given since date
func TimeFrameFromSince(since string) TimeFrame {
	if since == "" {
		return TimeFrame{}
	}
	sinceTime := since
	untilTime := time.Now().Format(dateFormat)

	return TimeFrame{since: sinceTime, until: untilTime}
}
