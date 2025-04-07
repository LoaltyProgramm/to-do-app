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

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if len(repeat) < 2 {
		return "", fmt.Errorf("\n%v Example of the format repeat: d 7. You format: %s", ERROR001, repeat)
	}

	parseRepeat := strings.Split(repeat, " ")
	if parseRepeat[0] != "d" && parseRepeat[0] != "y" {
		return "", fmt.Errorf("\n%v Example of the format repeat: d 7. You format: %s", ERROR002, repeat)
	}

	parseTimeDstart, err := time.Parse(layout, dstart)
	if err != nil {
		return "", fmt.Errorf("Error parse dstart: %v", err)
	}
	
	var next time.Time
	var isNotValidate = true

	switch parseRepeat[0] {
	case "d":
		days, err := strconv.Atoi(parseRepeat[1])
		if err != nil {
			return "", fmt.Errorf("Error convertation: %v", err)
		}

		valueDuration := time.Duration(24 * days) * time.Hour

		for isNotValidate {
			next = parseTimeDstart.Add(valueDuration)
			nextDateStr := next.Format(layout)
			nowDateStr := now.Format(layout)
			if nextDateStr >= nowDateStr{
				isNotValidate = false
			} else {
				parseTimeDstart = next
			}
		}

	}
	return fmt.Sprint(next.Format(layout)), nil
}