package task

import (
	"net/http"
	"net/url"

	"github.com/aikizoku/go-gae-template/src/middleware"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"google.golang.org/appengine/taskqueue"
)

type SampleQueueing struct {
}

func (s *SampleQueueing) HogeQueueing(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "cal sample task handler")

	values := url.Values{}
	task := taskqueue.NewPOSTTask("/task/sample", values)

	_, err := taskqueue.Add(ctx, task, "hoge-queue")
	if err != nil {
		log.Errorf(ctx, "add task error: %s", err.Error())
	}

	middleware.RenderSuccess(w)
}
