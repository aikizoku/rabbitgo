package util

import "fmt"

// RecordStart ... 速度計測開始
func RecordStart() int64 {
	return TimeNowUnix()
}

// RecordEnd ... 速度計測終了
func RecordEnd(start int64, msg string) int64 {
	end := TimeNowUnix()
	df := end - start
	fmt.Printf("%s %dms\n", msg, df)
	return df
}
