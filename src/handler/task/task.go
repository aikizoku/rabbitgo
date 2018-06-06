package task

import (
	"net/http"

	"github.com/aikizoku/go-gae-template/src/middleware"
	"github.com/aikizoku/go-gae-template/src/service"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

type SampleTask struct {
	Svc service.Sample
}

func (s *SampleTask) HogeTask(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	log.Debugf(ctx, "cal sample task handler")
	_ = middleware.GetTaskQueueHeaders(ctx, r)
	s.Svc.Hoge(ctx)
	middleware.RenderSuccess(w)
}
