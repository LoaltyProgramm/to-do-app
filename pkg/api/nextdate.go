package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	layout = "20060102"
)

const (
	ERROR001 = "ERROR001:"
	ERROR002 = "ERROR002:"
)

func comparingDate (nowDate, nextDate time.Time) bool {
	nextDateStr := nextDate.Format(layout)
	nowDateStr := nowDate.Format(layout)

	if nextDateStr >= nowDateStr {
		return true
	} else {
		return false
	}
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	parseRepeat := strings.Split(repeat, " ")

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

		for {
			next = parseTimeDstart.AddDate(0, 0, days)

			if comparingDate(now, next) {
				break
			} else {
				parseTimeDstart = next
			}
		}

	case "y":
		next = parseTimeDstart.AddDate(1, 0, 0)
	default:
		return "", fmt.Errorf("%v Unsupported format %v", ERROR002, parseRepeat[0])
	}

	return fmt.Sprint(next.Format(layout)), nil
}

func main() {
	a, err := NextDate(time.Now(), "20250329", "d 7")
	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(a)
}