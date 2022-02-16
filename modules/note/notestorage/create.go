package notestorage

import (
	"FirstProject/modules/note/notemodel"
	"context"
)

func (s *sqlStore) Create(ctx context.Context, data *notemodel.NoteCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
