package cloudfirestore

import (
	"context"
	"reflect"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"

	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/log"
	"github.com/aikizoku/rabbitgo/appengine/default/src/lib/util"
)

// GenerateDocumentRef ... ドキュメント参照を作成する
func GenerateDocumentRef(fCli *firestore.Client, docRefs []*DocRef) *firestore.DocumentRef {
	var dst *firestore.DocumentRef
	for i, docRef := range docRefs {
		if i == 0 {
			dst = fCli.Collection(docRef.CollectionName).Doc(docRef.DocID)
		} else {
			dst = dst.Collection(docRef.CollectionName).Doc(docRef.DocID)
		}
	}
	return dst
}

// Get ... １つ取得する
func Get(ctx context.Context, docRef *firestore.DocumentRef, dst interface{}) (bool, error) {
	dsnp, err := docRef.Get(ctx)
	if dsnp != nil && !dsnp.Exists() {
		return false, nil
	}
	if err != nil {
		log.Errorm(ctx, "docRef.Get", err)
		return false, err
	}
	err = dsnp.DataTo(dst)
	if err != nil {
		log.Errorm(ctx, "dsnp.DataTo", err)
		return false, err
	}
	setDocByDst(dst, dsnp.Ref)
	return true, nil
}

// GetMulti ... 複数取得する
func GetMulti(ctx context.Context, fCli *firestore.Client, docRefs []*firestore.DocumentRef, dsts interface{}) error {
	dsnps, err := fCli.GetAll(ctx, docRefs)
	if err != nil {
		log.Errorm(ctx, "fCli.GetAll", err)
		return err
	}
	rv := reflect.Indirect(reflect.ValueOf(dsts))
	rrt := rv.Type().Elem().Elem()
	for _, dsnp := range dsnps {
		if !dsnp.Exists() {
			continue
		}
		v := reflect.New(rrt).Interface()
		err = dsnp.DataTo(&v)
		if err != nil {
			log.Errorm(ctx, "dsnp.DataTo", err)
			return err
		}
		rrv := reflect.ValueOf(v)
		setDocByDsts(rrv, rrt, dsnp.Ref)
		rv.Set(reflect.Append(rv, rrv))
	}
	return nil
}

// GetByQuery ... クエリで１つ取得する
func GetByQuery(ctx context.Context, query firestore.Query, dst interface{}) (bool, error) {
	it := query.Documents(ctx)
	defer it.Stop()
	dsnp, err := it.Next()
	if err == iterator.Done {
		return false, nil
	}
	err = dsnp.DataTo(dst)
	if err != nil {
		log.Errorm(ctx, "dsnp.DataTo", err)
		return false, err
	}
	setDocByDst(dst, dsnp.Ref)
	return true, nil
}

// GetMultiByQuery ... クエリで複数取得する
func GetMultiByQuery(ctx context.Context, query firestore.Query, dsts interface{}) error {
	it := query.Documents(ctx)
	defer it.Stop()
	rv := reflect.Indirect(reflect.ValueOf(dsts))
	rrt := rv.Type().Elem().Elem()
	for {
		dsnp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Errorm(ctx, "it.Next", err)
			return err
		}
		v := reflect.New(rrt).Interface()
		err = dsnp.DataTo(&v)
		if err != nil {
			log.Errorm(ctx, "dsnp.DataTo", err)
			return err
		}
		rrv := reflect.ValueOf(v)
		setDocByDsts(rrv, rrt, dsnp.Ref)
		rv.Set(reflect.Append(rv, rrv))
	}
	return nil
}

// GetMultiByQueryCursor ... クエリで複数取得する（ページング）
func GetMultiByQueryCursor(ctx context.Context, query firestore.Query, limit int, cursor *firestore.DocumentSnapshot, dsts interface{}) (*firestore.DocumentSnapshot, error) {
	if cursor != nil {
		query = query.StartAfter(cursor)
	}
	it := query.Limit(limit).Documents(ctx)
	defer it.Stop()
	rv := reflect.Indirect(reflect.ValueOf(dsts))
	rrt := rv.Type().Elem().Elem()
	var lastDsnp *firestore.DocumentSnapshot
	for {
		dsnp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Errorm(ctx, "it.Next", err)
			return nil, err
		}
		v := reflect.New(rrt).Interface()
		err = dsnp.DataTo(v)
		if err != nil {
			log.Errorm(ctx, "doc.DataTo", err)
			return nil, err
		}
		rrv := reflect.ValueOf(v)
		setDocByDsts(rrv, rrt, dsnp.Ref)
		rv.Set(reflect.Append(rv, rrv))
		lastDsnp = dsnp
	}
	if rv.Len() == limit {
		return lastDsnp, nil
	}
	return nil, nil
}

// TxGet ... １つ取得する（トランザクション）
func TxGet(ctx context.Context, tx *firestore.Transaction, docRef *firestore.DocumentRef, dst interface{}) (bool, error) {
	dsnp, err := tx.Get(docRef)
	if dsnp != nil && !dsnp.Exists() {
		return false, nil
	}
	if err != nil {
		log.Errorm(ctx, "docRef.Get", err)
		return false, err
	}
	err = dsnp.DataTo(dst)
	if err != nil {
		log.Errorm(ctx, "dsnp.DataTo", err)
		return false, err
	}
	setDocByDst(dst, dsnp.Ref)
	return true, nil
}

// TxGetMulti ... 複数取得する（トランザクション）
func TxGetMulti(ctx context.Context, tx *firestore.Transaction, docRefs []*firestore.DocumentRef, dsts interface{}) error {
	dsnps, err := tx.GetAll(docRefs)
	if err != nil {
		log.Errorm(ctx, "tx.GetAll", err)
		return err
	}
	rv := reflect.Indirect(reflect.ValueOf(dsts))
	rrt := rv.Type().Elem().Elem()
	for _, dsnp := range dsnps {
		if !dsnp.Exists() {
			continue
		}
		v := reflect.New(rrt).Interface()
		err = dsnp.DataTo(&v)
		if err != nil {
			log.Errorm(ctx, "dsnp.DataTo", err)
			return err
		}
		rrv := reflect.ValueOf(v)
		setDocByDsts(rrv, rrt, dsnp.Ref)
		rv.Set(reflect.Append(rv, rrv))
	}
	return nil
}

// TxGetByQuery ... クエリで１つ取得する（トランザクション）
func TxGetByQuery(ctx context.Context, tx *firestore.Transaction, query firestore.Query, dst interface{}) (bool, error) {
	it := tx.Documents(query)
	defer it.Stop()
	dsnp, err := it.Next()
	if err == iterator.Done {
		return false, nil
	}
	err = dsnp.DataTo(dst)
	if err != nil {
		log.Errorm(ctx, "dsnp.DataTo", err)
		return false, err
	}
	setDocByDst(dst, dsnp.Ref)
	return true, nil
}

// TxGetMultiByQuery ... クエリで複数取得する（トランザクション）
func TxGetMultiByQuery(ctx context.Context, tx *firestore.Transaction, query firestore.Query, dsts interface{}) error {
	it := tx.Documents(query)
	defer it.Stop()
	rv := reflect.Indirect(reflect.ValueOf(dsts))
	rrt := rv.Type().Elem().Elem()
	for {
		dsnp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Errorm(ctx, "it.Next", err)
			return err
		}
		v := reflect.New(rrt).Interface()
		err = dsnp.DataTo(&v)
		if err != nil {
			log.Errorm(ctx, "dsnp.DataTo", err)
			return err
		}
		rrv := reflect.ValueOf(v)
		setDocByDsts(rrv, rrt, dsnp.Ref)
		rv.Set(reflect.Append(rv, rrv))
	}
	return nil
}

// Create ... 作成する
func Create(ctx context.Context, colRef *firestore.CollectionRef, src interface{}) error {
	ref, _, err := colRef.Add(ctx, src)
	if err != nil {
		log.Errorm(ctx, "colRef.Add", err)
		return err
	}
	setDocByDst(src, ref)
	return nil
}

// BtCreate ... 作成する（バッチ書き込み）
func BtCreate(ctx context.Context, bt *firestore.WriteBatch, colRef *firestore.CollectionRef, src interface{}) {
	id := util.StrUniqueID()
	ref := colRef.Doc(id)
	bt.Create(ref, src)
	setDocByDst(src, ref)
}

// TxCreate ... 作成する（トランザクション）
func TxCreate(ctx context.Context, tx *firestore.Transaction, colRef *firestore.CollectionRef, src interface{}) error {
	id := util.StrUniqueID()
	ref := colRef.Doc(id)
	err := tx.Create(ref, src)
	if err != nil {
		log.Errorm(ctx, "tx.Create", err)
		return err
	}
	setDocByDst(src, ref)
	return nil
}

// Update ... 更新する
func Update(ctx context.Context, docRef *firestore.DocumentRef, kv map[string]interface{}) error {
	srcs := []firestore.Update{}
	for k, v := range kv {
		src := firestore.Update{Path: k, Value: v}
		srcs = append(srcs, src)
	}
	_, err := docRef.Update(ctx, srcs)
	if err != nil {
		log.Errorm(ctx, "docRef.Update", err)
		return err
	}
	return nil
}

// BtUpdate ... 更新する（バッチ書き込み）
func BtUpdate(ctx context.Context, bt *firestore.WriteBatch, docRef *firestore.DocumentRef, kv map[string]interface{}) {
	srcs := []firestore.Update{}
	for k, v := range kv {
		src := firestore.Update{Path: k, Value: v}
		srcs = append(srcs, src)
	}
	_ = bt.Update(docRef, srcs)
}

// TxUpdate ... 更新する（トランザクション）
func TxUpdate(ctx context.Context, tx *firestore.Transaction, docRef *firestore.DocumentRef, kv map[string]interface{}) error {
	srcs := []firestore.Update{}
	for k, v := range kv {
		src := firestore.Update{Path: k, Value: v}
		srcs = append(srcs, src)
	}
	err := tx.Update(docRef, srcs)
	if err != nil {
		log.Errorm(ctx, "tx.Update", err)
		return err
	}
	return nil
}

// Set ... 上書きする
func Set(ctx context.Context, docRef *firestore.DocumentRef, src interface{}) error {
	_, err := docRef.Set(ctx, src)
	if err != nil {
		log.Errorm(ctx, "docRef.Set", err)
		return err
	}
	setDocByDst(src, docRef)
	return nil
}

// BtSet ... 上書きする（バッチ書き込み）
func BtSet(ctx context.Context, bt *firestore.WriteBatch, docRef *firestore.DocumentRef, src interface{}) {
	_ = bt.Set(docRef, src)
	setDocByDst(src, docRef)
}

// TxSet ... 上書きする（トランザクション）
func TxSet(ctx context.Context, tx *firestore.Transaction, docRef *firestore.DocumentRef, src interface{}) error {
	err := tx.Set(docRef, src)
	if err != nil {
		log.Errorm(ctx, "tx.Set", err)
		return err
	}
	setDocByDst(src, docRef)
	return nil
}

// Delete ... 削除する
func Delete(ctx context.Context, docRef *firestore.DocumentRef) error {
	_, err := docRef.Delete(ctx)
	if err != nil {
		log.Errorm(ctx, "docRef.Delete", err)
		return err
	}
	return nil
}

// BtDelete ... 削除する（バッチ書き込み）
func BtDelete(ctx context.Context, bt *firestore.WriteBatch, docRef *firestore.DocumentRef) {
	_ = bt.Delete(docRef)
}

// TxDelete ... 削除する（トランザクション）
func TxDelete(ctx context.Context, tx *firestore.Transaction, docRef *firestore.DocumentRef) error {
	err := tx.Delete(docRef)
	if err != nil {
		log.Errorm(ctx, "tx.Delete", err)
		return err
	}
	return nil
}
