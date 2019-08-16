package content

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/k0kubun/pp"

	"github.com/aikizoku/rabbitgo/appengine/src/lib/cloudfirestore"
	"github.com/aikizoku/rabbitgo/appengine/src/lib/util"
	"github.com/aikizoku/rabbitgo/appengine/src/model"
)

// Sample ... サンプルのシードデータ
type Sample struct {
	FCli *firestore.Client
}

// Generate ... シードデータを作成する
func (m *Sample) Generate(ctx context.Context) {
	// 大量に作成
	m.batchCreate(ctx)

	// Get
	//m.get(ctx)

	// GetMulti
	// m.getMulti(ctx)

	// GetByQuery
	//m.getByQuery(ctx)

	// GetMultiByQuery
	//m.getMultiByQuery(ctx)

	// GetMultiByQueryAndCursor
	// m.getMultiByQueryCursor(ctx)
}

func (m *Sample) batchCreate(ctx context.Context) {
	colRef := m.FCli.Collection("sample")
	bt := m.FCli.Batch()
	for i := 0; i < 50; i++ {
		src := &model.Sample{
			Category:  "a",
			Name:      fmt.Sprintf("%d", i),
			Disabled:  false,
			CreatedAt: util.TimeNowUnix(),
		}
		cloudfirestore.BtCreate(ctx, bt, colRef, src)
	}
	_, err := bt.Commit(ctx)
	if err != nil {
		panic(err)
	}
}

func (m *Sample) get(ctx context.Context) {
	ref := m.FCli.Collection("sample").Doc("0JJ91ryPgeBVawkmkt0W")
	dst := &model.Sample{}
	exist, err := cloudfirestore.Get(ctx, ref, dst)
	if err != nil {
		panic(err)
	}
	if !exist {
		// NotFound
		return
	}
	pp.Println(dst)
}

func (m *Sample) getMulti(ctx context.Context) {
	dsts := []*model.Sample{}
	err := cloudfirestore.GetMulti(ctx, m.FCli, []*firestore.DocumentRef{
		m.FCli.Collection("sample").Doc("0JJ91ryPgeBVawkmkt0W"),
		m.FCli.Collection("sample").Doc("0P6Bii8KMmzcuEexkhnb"),
		m.FCli.Collection("sample").Doc("33NfZSbzFKdJn4o8bruQ"),
		m.FCli.Collection("sample").Doc("3gftXUHb7l72Htn2C0lx"),
		m.FCli.Collection("sample").Doc("4hNoG2Mhy1xIGXrp86HD"),
		m.FCli.Collection("sample").Doc("5HNMdSFZ3tqa7arOGWr7"),
		m.FCli.Collection("sample").Doc("6WSCssPzcQqjMCDDuyFV"),
		m.FCli.Collection("sample").Doc("7SScTJVwPod4iPuINLv0"),
		m.FCli.Collection("sample").Doc("8nKCJpYpOGJWzXyMqDlf"),
		m.FCli.Collection("sample").Doc("8ym2pSmJ2ZAHajWjEbKS"),
	}, &dsts)
	if err != nil {
		panic(err)
	}
	if len(dsts) == 0 {
		// NotFound
		return
	}
	pp.Println(dsts)
}

func (m *Sample) getByQuery(ctx context.Context) {
	query := m.FCli.Collection("sample").Where("name", "==", "22")
	dst := &model.Sample{}
	exist, err := cloudfirestore.GetByQuery(ctx, query, dst)
	if err != nil {
		panic(err)
	}
	if !exist {
		// NotFound
		return
	}
	pp.Println(dst)
}

func (m *Sample) getMultiByQuery(ctx context.Context) {
	query := m.FCli.Collection("sample").Where("disabled", "==", false)
	dsts := []*model.Sample{}
	err := cloudfirestore.GetMultiByQuery(ctx, query, &dsts)
	if err != nil {
		panic(err)
	}
	if len(dsts) == 0 {
		// NotFound
		return
	}
	pp.Println(dsts)
}

func (m *Sample) getMultiByQueryCursor(ctx context.Context) {
	// １回目のリクエスト
	query := m.FCli.Collection("sample").Where("disabled", "==", false)
	dsts := []*model.Sample{}
	nextCursor, err := cloudfirestore.GetMultiByQueryCursor(ctx, query, 30, nil, &dsts)
	if err != nil {
		panic(err)
	}
	if len(dsts) == 0 {
		// NotFound
		return
	}
	fmt.Printf("1: %d, %v\n", len(dsts), nextCursor)

	// ２回目のリクエスト
	dsts = []*model.Sample{}
	nextCursor, err = cloudfirestore.GetMultiByQueryCursor(ctx, query, 30, nextCursor, &dsts)
	if err != nil {
		panic(err)
	}
	if len(dsts) == 0 {
		// NotFound
		return
	}
	fmt.Printf("2: %d, %s\n", len(dsts), nextCursor)
}
