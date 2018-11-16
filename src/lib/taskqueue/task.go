package taskqueue

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/aikizoku/beego/src/lib/log"
	"google.golang.org/appengine/taskqueue"
)

// AddTask ... 通常リクエストをEnqueueする
func AddTask(ctx context.Context, queue string, path string, params url.Values) error {
	task := taskqueue.NewPOSTTask(queue, params)
	return Add(ctx, queue, task)
}

// AddTaskToJSON ... JSONのリクエストをEnqueueする
func AddTaskToJSON(ctx context.Context, queue string, path string, src interface{}) error {
	h := make(http.Header)
	h.Set("Content-Type", "application/json")
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
