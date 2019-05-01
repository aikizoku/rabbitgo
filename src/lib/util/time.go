package util

import (
	"time"
)

// TimeNow ... 現在時刻をJSTのTimeで取得する
func TimeNow() time.Time {
	return time.Now().In(TimeZoneJST())
}

// TimeNowUnix ... 現在時刻をJSTのUnixTimeStamp(ミリ秒)で取得する
func TimeNowUnix() int64 {
	return time.Now().In(TimeZoneJST()).UnixNano() / int64(time.Millisecond)
}

// TimeByUnix ... UnixTimestampからJSTのTimeを取得する
func TimeByUnix(u int64) time.Time {
	return time.Unix(u, 0).In(TimeZoneJST())
}

// TimeToUnix ... TimeからUnixTimeStamp(ミリ秒)に変換する
func TimeToUnix(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// TimeZoneJST ... 日本のタイムゾーンを取得する
func TimeZoneJST() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}
