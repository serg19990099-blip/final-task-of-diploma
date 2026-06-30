package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "20060102"

func afterNow(date, now time.Time) bool {
	y1, m1, d1 := date.Date()
	y2, m2, d2 := now.Date()

	dateOnly := time.Date(y1, m1, d1, 0, 0, 0, 0, time.UTC)
	nowOnly := time.Date(y2, m2, d2, 0, 0, 0, 0, time.UTC)

	return dateOnly.After(nowOnly)
}

func NextDate(now time.Time, dstart string, repeat string) (string, error) {
	if repeat == "" {
		return "", errors.New("empty repeat")
	}

	date, err := time.Parse(dateFormat, dstart)
	if err != nil {
		return "", err
	}

	parts := strings.Fields(repeat)
	if len(parts) == 0 {
		return "", errors.New("invalid repeat")
	}

	switch parts[0] {
	case "d":
		if len(parts) != 2 {
			return "", errors.New("invalid daily repeat")
		}

		interval, err := strconv.Atoi(parts[1])
		if err != nil {
			return "", err
		}

		if interval <= 0 || interval > 400 {
			return "", errors.New("invalid daily interval")
		}

		for {
			date = date.AddDate(0, 0, interval)
			if afterNow(date, now) {
				return date.Format(dateFormat), nil
			}
		}

	case "y":
		if len(parts) != 1 {
			return "", errors.New("invalid yearly repeat")
		}

		for {
			date = date.AddDate(1, 0, 0)
			if afterNow(date, now) {
				return date.Format(dateFormat), nil
			}
		}

	default:
		return "", errors.New("unsupported repeat format")
	}
}

func NextDateHandler(w http.ResponseWriter, r *http.Request) {
	nowStr := r.URL.Query().Get("now")
	dateStr := r.URL.Query().Get("date")
	repeat := r.URL.Query().Get("repeat")

	now, err := time.Parse(dateFormat, nowStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	next, err := NextDate(now, dateStr, repeat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte(next))
}
