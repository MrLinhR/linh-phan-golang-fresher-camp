package notebiz

import (
	"FirstProject/modules/note/notemodel"
	"context"
	"errors"
)

type CreateNoteStore interface {
	Create(ctx context.Context, data *notemodel.NoteCreate) error
}
type createNoteBiz struct {
	store CreateNoteStore
}

func NewCreateNoteBiz(store CreateNoteStore) *createNoteBiz {
	return &createNoteBiz{store: store}
}

func (biz *createNoteBiz) CreateNote(ctx context.Context, data *notemodel.NoteCreate) error {
	if data.Title == "" {
		return errors.New("Note's Id can not be blank")
	}
	err := biz.store.Create(ctx, data)

	return err
}
