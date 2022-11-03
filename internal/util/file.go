package util

import "time"

func GetTimeFileName() string {
	return time.Now().Format("2006-01-02T15:04") + ".png"
}
