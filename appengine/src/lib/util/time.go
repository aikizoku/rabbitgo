package util

import (
	"time"
)

// TimeNow ... 現在時刻をJSTのTimeで取得する
func TimeNow() time.Time {
	return time.Now().In(TimeZoneJST())
}

// TimeNowUnix ... 現在時刻をJSTのUnixtimestamp(ミリ秒)で取得する
func TimeNowUnix() int64 {
	return time.Now().In(TimeZoneJST()).UnixNano() / int64(time.Millisecond)
}

// TimeByUnix ... Unixtimestamp(ミリ秒)からJSTのTimeを取得する
func TimeByUnix(u int64) time.Time {
	uNano := u * 1000 * 1000
	uSec := u / 1000
	return time.Unix(uSec, uNano-(uSec*1000*1000*1000)).In(TimeZoneJST())
}

// TimeToUnix ... TimeからUnixtimestamp(ミリ秒)に変換する
func TimeToUnix(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

// TimeZoneJST ... 日本のタイムゾーンを取得する
func TimeZoneJST() *time.Location {
	return time.FixedZone("Asia/Tokyo", 9*60*60)
}

// TimeSecondsToMiliseconds ... 指定秒数をミリ秒に変換する
func TimeSecondsToMiliseconds(seconds int) int64 {
	return int64(seconds * 1000)
}

// TimeMinutesToMiliseconds ... 指定分数をミリ秒に変換する
func TimeMinutesToMiliseconds(minutes int) int64 {
	return int64(minutes * 60 * 1000)
}

// TimeHoursToMiliseconds ... 指定時数をミリ秒に変換する
func TimeHoursToMiliseconds(hours int) int64 {
	return int64(hours * 60 * 60 * 1000)
}

// TimeDaysToMiliseconds ... 指定日数をミリ秒に変換する
func TimeDaysToMiliseconds(days int) int64 {
	return int64(days * 24 * 60 * 60 * 1000)
}
