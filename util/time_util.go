package util

import "time"

func TimeStampBefore24h() int64 {
	now := time.Now().UTC()
	d, _ := time.ParseDuration("-24h")
	d24h := now.Add(d).UTC()
	return d24h.Unix()
}

func GetCurrentDatetime(datetime int64, period string) int64 {
	datetimeStr := time.Unix(int64(datetime), 0).UTC().Format("2006-01-02 15:04:05")
	dt, _ := time.ParseInLocation("2006-01-02 15:04:05", datetimeStr, time.UTC)
	newDatetime := int64(0)
	switch period {
	case "1min":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute(), 0, 0, time.UTC).Unix()
	case "5min":
		step := dt.Minute() % 5
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-step, 0, 0, time.UTC).Unix()
	case "15min":
		step := dt.Minute() % 15
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-step, 0, 0, time.UTC).Unix()
	case "30min":
		step := dt.Minute() % 30
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-step, 0, 0, time.UTC).Unix()
	case "1hour":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), 0, 0, 0, time.UTC).Unix()
	case "4hour":
		step := dt.Hour() % 4
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour()-step, 0, 0, 0, time.UTC).Unix()
	case "1day":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), 0, 0, 0, 0, time.UTC).Unix()
	case "1week":
		step := int(dt.Weekday())
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day()-step, 0, 0, 0, 0, time.UTC).Unix()
	case "1mon":
		newDatetime = time.Date(dt.Year(), dt.Month(), 1, 0, 0, 0, 0, time.UTC).Unix()
	}
	return newDatetime
}

func GetBeforeDatetime(datetime int64, period string) int64 {
	datetimeStr := time.Unix(int64(datetime), 0).UTC().Format("2006-01-02 15:04:05")
	dt, _ := time.ParseInLocation("2006-01-02 15:04:05", datetimeStr, time.UTC)
	newDatetime := int64(0)
	switch period {
	case "1min":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-1, 0, 0, time.UTC).Unix()
	case "5min":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-5, 0, 0, time.UTC).Unix()
	case "15min":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-15, 0, 0, time.UTC).Unix()
	case "30min":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()-30, 0, 0, time.UTC).Unix()
	case "1hour":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour()-1, 0, 0, 0, time.UTC).Unix()
	case "4hour":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour()-4, 0, 0, 0, time.UTC).Unix()
	case "1day":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day()-1, 0, 0, 0, 0, time.UTC).Unix()
	case "1week":
		newDatetime = time.Date(dt.Year(), dt.Month(), dt.Day()-7, 0, 0, 0, 0, time.UTC).Unix()
	case "1mon":
		newDatetime = time.Date(dt.Year(), dt.Month()-1, 1, 0, 0, 0, 0, time.UTC).Unix()
	}
	return newDatetime
}

func GetNextDatetime(datetime int64, period string) int64 {
	datetimeStr := time.Unix(int64(datetime), 0).UTC().Format("2006-01-02 15:04:05")
	dt, _ := time.ParseInLocation("2006-01-02 15:04:05", datetimeStr, time.UTC)
	nextDatetime := int64(0)
	switch period {
	case "1min":
		nextDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()+1, 0, 0, time.UTC).Unix()
	case "5min":
		nextDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()+5, 0, 0, time.UTC).Unix()
	case "15min":
		nextDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()+15, 0, 0, time.UTC).Unix()
	case "30min":
		nextDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour(), dt.Minute()+30, 0, 0, time.UTC).Unix()
	case "1hour":
		nextDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour()+1, 0, 0, 0, time.UTC).Unix()
	case "4hour":
		nextDatetime = time.Date(dt.Year(), dt.Month(), dt.Day(), dt.Hour()+4, 0, 0, 0, time.UTC).Unix()
	case "1day":
		nextDatetime = time.Date(dt.Year(), dt.Month(), dt.Day()+1, 0, 0, 0, 0, time.UTC).Unix()
	case "1week":
		nextDatetime = time.Date(dt.Year(), dt.Month(), dt.Day()+7, 0, 0, 0, 0, time.UTC).Unix()
	case "1mon":
		nextDatetime = time.Date(dt.Year(), dt.Month()+1, 1, 0, 0, 0, 0, time.UTC).Unix()
	}
	return nextDatetime
}
