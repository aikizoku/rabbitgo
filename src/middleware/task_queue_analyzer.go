package middleware

import (
	"context"
	"net/http"
	"strconv"

	"google.golang.org/appengine/log"
)

// TaskQueueHeaders ... GAEからのタスクキューリクエストのヘッダー情報
type TaskQueueHeaders struct {
	// キューの名前
	// デフォルトの push キューの場合は「default」
	QueueName string
	// タスクの名前
	// 名前が指定されていない場合は、システムが生成した一意のID
	TaskName string
	// このタスクが再試行された回数
	// 最初の試行の場合はこの値は 0
	// この試行回数にはインスタンス数不足が原因でタスクが異常終了したため実行フェーズに到達できなかった試行も含まれています
	TaskRetryCount int
	// このタスクがこれまでに実行フェーズ中に異常終了した回数
	// この回数にはインスタンス数不足が原因の失敗は含まれていません
	TaskExecutionCount int
	// タスクの目標実行時間
	// 1970 年 1 月 1 日からの秒数で表します
	TaskETA string
}

// GetTaskQueueHeaders ... GAEからのタスクキューリクエストのヘッダー情報を取得する
func GetTaskQueueHeaders(ctx context.Context, r *http.Request) *TaskQueueHeaders {
	qn := r.Header.Get("X-AppEngine-QueueName")

	tn := r.Header.Get("X-AppEngine-TaskName")

	trc := r.Header.Get("X-AppEngine-TaskRetryCount")
	itrc, err := strconv.Atoi(trc)
	if err != nil {
		log.Warningf(ctx, "TaskRetryCount parse int error: %s value=%s", err.Error(), trc)
	}

	tec := r.Header.Get("X-AppEngine-TaskExecutionCount")
	itec, err := strconv.Atoi(tec)
	if err != nil {
		log.Warningf(ctx, "TaskExecutionCount parse int error: %s value=%s", err.Error(), tec)
	}

	teta := r.Header.Get("X-AppEngine-TaskETA")

	log.Debugf(ctx, `TaskQueueHeaders
		QueueName: %s
		TaskName: %s
		TaskRetryCount: %d
		TaskExecutionCount: %d
		TaskETA: %s`, qn, tn, itrc, itec, teta)

	return &TaskQueueHeaders{
		QueueName:          qn,
		TaskName:           tn,
		TaskRetryCount:     itrc,
		TaskExecutionCount: itec,
		TaskETA:            teta,
	}
}
