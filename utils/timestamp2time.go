package utils

import "time"

func ConvertToTime(stamp int64) string {
	format := time.Unix(stamp, 0).Format("2006-01-02 15:04:05")
	return format
}
