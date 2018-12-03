package util

import (
	"time"
)

// TimeNow ... 現在時刻をJSTのTimeで取得する
func TimeNow() time.Time {
	return time.Now().In(timeZoneJST())
}

// TimeNowUnix ... 現在時刻をJSTのUnixTimeStamp(ミリ秒)で取得する
func TimeNowUnix() int64 {
	return time.Now().In(timeZoneJST()).UnixNano() / int64(time.Millisecond)
}

// TimeUnix ... UnixTimestampからJSTのTimeを取得する
func TimeUnix(u int64) time.Time {
	return time.Unix(u, 0).In(timeZoneJST())
}

func timeZoneJST() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}
