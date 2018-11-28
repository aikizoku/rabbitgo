package taskqueue

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/aikizoku/push/src/lib/internalauth"

	"github.com/aikizoku/push/src/lib/log"
	"google.golang.org/appengine/taskqueue"
)

// AddTask ... 通常リクエストをEnqueueする
func AddTask(ctx context.Context, queue string, path string, params url.Values) error {
	h := make(http.Header)
	h.Set("Content-Type", "application/x-www-form-urlencoded")
	h.Set(internalauth.GetHeader())
	task := &taskqueue.Task{
		Path:    path,
		Payload: []byte(params.Encode()),
		Header:  h,
		Method:  "POST",
	}
	return Add(ctx, queue, task)
}

// AddTaskByJSON ... JSONのリクエストをEnqueueする
func AddTaskByJSON(ctx context.Context, queue string, path string, src interface{}) error {
	h := make(http.Header)
	h.Set("Content-Type", "application/json")
	h.Set(internalauth.GetHeader())
	data, err := json.Marshal(src)
	if err != nil {
		log.Errorf(ctx, "json.Marshal error: %s", err.Error())
		return err
	}
	task := &taskqueue.Task{
		Path:    path,
		Payload: data,
		Header:  h,
		Method:  "POST",
	}
	return Add(ctx, queue, task)
}

// Add ... リクエストをEnqueueする
func Add(ctx context.Context, queue string, task *taskqueue.Task) error {
	_, err := taskqueue.Add(ctx, task, queue)
	if err != nil {
		log.Errorf(ctx, "taskqueue.Add error: %s", err.Error())
		return err
	}
	return nil
}
