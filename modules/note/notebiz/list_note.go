package notebiz

import (
	"FirstProject/common"
	"FirstProject/modules/note/notemodel"
	"context"
)

type ListNoteStore interface {
	ListDataByCondition(ctx context.Context, conditions map[string]interface{}, paging *common.Paging, morekeys ...string) ([]notemodel.Note, error)
}

type listNoteBiz struct {
	store ListNoteStore
}

func NewListNoteBiz(store ListNoteStore) *listNoteBiz {
	return &listNoteBiz{store: store}
}

func (biz *listNoteBiz) ListNote(ctx context.Context, paging *common.Paging) ([]notemodel.Note, error) {
	result, err := biz.store.ListDataByCondition(ctx, nil, paging)
	return result, err
}

func (biz *listNoteBiz) ListNoteById(ctx context.Context, id int, paging *common.Paging) ([]notemodel.Note, error) {
	conditions := map[string]interface{}{"id": id}
	result, err := biz.store.ListDataByCondition(ctx, conditions, paging)
	return result, err
}
