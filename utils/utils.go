package utils

import (
	"time"

	naturalDate "github.com/tj/go-naturaldate"
)

func ParseDate(date string) (time.Time, error) {
	parsedDate, err := naturalDate.Parse(date, time.Now())
	if err != nil {
		return time.Time{}, err
	}

	return parsedDate, nil
}
