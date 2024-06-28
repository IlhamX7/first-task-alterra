package helper

import (
	"fmt"
	"time"
)

func StringToDate(birthDate string) (time.Time, error) {
	if birthDate == "" {
		return time.Time{}, fmt.Errorf("birthDate string is empty")
	}
	layout := "2006-01-02"
	parsedTime, err := time.Parse(layout, birthDate)
	if err != nil {
		return time.Time{}, fmt.Errorf("error parsing date: %w", err)
	}
	return parsedTime, nil
}
