package notebiz

import (
	"FirstProject/modules/note/notemodel"
	"context"
)

type UpdateNoteStore interface {
	Update(ctx context.Context, data *notemodel.Note, conditions map[string]interface{}, morekeys ...string) error
}
type updateNoteBiz struct {
	store UpdateNoteStore
}

func NewUpdateNoteBiz(store UpdateNoteStore) *updateNoteBiz {
	return &updateNoteBiz{store: store}
}

func (biz *updateNoteBiz) UpdateNote(ctx context.Context, data *notemodel.Note, id int) error {
	conditions := map[string]interface{}{"id": id}
	err := biz.store.Update(ctx, data, conditions)

	return err
}
