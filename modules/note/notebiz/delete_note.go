package notebiz

import (
	"FirstProject/modules/note/notemodel"
	"context"
)

type DeleteNoteStore interface {
	Delete(ctx context.Context, conditions map[string]interface{}, morekeys ...string) ([]notemodel.Note, error)
}
type deleteNoteBiz struct {
	store DeleteNoteStore
}

func NewDeleteNoteBiz(store DeleteNoteStore) *deleteNoteBiz {
	return &deleteNoteBiz{store: store}
}

func (biz *deleteNoteBiz) DeleteNote(ctx context.Context, id int) ([]notemodel.Note, error) {
	conditions := map[string]interface{}{"id": id}
	result, err := biz.store.Delete(ctx, conditions)

	return result, err
}
