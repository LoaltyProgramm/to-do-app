package api

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	layout = "20060102"
)

func comparingDate(nowDate, nextDate time.Time) bool {
	nextDateStr := nextDate.Format(layout)
	nowDateStr := nowDate.Format(layout)

	if nextDateStr > nowDateStr {
		return true
	} else {
		return false
	}
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	parseRepeat := strings.Split(repeat, " ")

	if len(parseRepeat) < 2 && parseRepeat[0] == "d" {
		return "", errors.New("invalid repeat format: expected 'd <number>")
	}

	parseTimeDstart, err := time.Parse(layout, dstart)
	if err != nil {
		return "", fmt.Errorf("Error parse dstart: %v", err)
	}
	
	var next time.Time

	switch parseRepeat[0] {
	case "d":
		days, err := strconv.Atoi(parseRepeat[1])
		if err != nil {
			return "", fmt.Errorf("Error convertation: %v", err)
		}

		if days > 400 {
			return "", errors.New("The maximum value of the day can be 400")
		}
		
		for {
			next = parseTimeDstart.AddDate(0, 0, days)

			if comparingDate(now, next) {
				break
			} else {
				parseTimeDstart = next
			}
		}

	case "y":
		for {
			next = parseTimeDstart.AddDate(1, 0, 0)
			if comparingDate(now, next) {
				break
			} else {
				parseTimeDstart = next
			}
		}
	default:
		return "", fmt.Errorf("Unsupported format %v", parseRepeat[0])
	}

	return fmt.Sprint(next.Format(layout)), nil
}