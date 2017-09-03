package utils

import (
	"time"
)

//MillisInMinute for just year and month
const (
	MillisInMinute = int64(time.Minute) / int64(time.Millisecond)
	MillisInHour   = MillisInMinute * 60
	MillisInDay    = MillisInHour * 24
	MillisInWeek   = MillisInDay * 7
	MillisInMonth  = MillisInDay * 30
	MillisInYear   = MillisInDay * 365
)

//GetCurrentTime return current time in millis
func GetCurrentTime() int64 {
	return time.Now().UnixNano() * int64(time.Nanosecond) / int64(time.Millisecond)
}

//GetTimestampFromTime return timestamp in millis from time
func GetTimestampFromTime(data time.Time) int64 {
	return data.UnixNano() * int64(time.Nanosecond) / int64(time.Millisecond)
}

//GetTimefromTimeStamp converts int64 to timestamp
func GetTimefromTimeStamp(data int64) time.Time {
	tm := time.Unix(data, 0)
	return tm
}

//GetXDaysTimeInMiliSeconds return time miliseconds for the given number days
func GetXDaysTimeInMiliSeconds(day int64) int64 {
	return 1000 * 60 * 60 * 24 * day
}
