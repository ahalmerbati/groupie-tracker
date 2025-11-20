package utils

import (
	"log"
	"time"
)

func FindLastConcertDate(datesLocations map[string][]string) string {
	var latestTime time.Time
	const format = "02-01-2006"

	for _, dates := range datesLocations {
		for _, dateStr := range dates {
			t, err := time.Parse(format, dateStr)
			if err != nil {
				log.Printf("Error parsing date: %s - %v", dateStr, err)
				continue
			}

			if t.After(latestTime) {
				latestTime = t
			}
		}
	}

	if latestTime.IsZero() {
		return "N/A"
	}

	return latestTime.Format(format)
}
