package notestorage

import (
	"FirstProject/modules/note/notemodel"
	"context"
)

func (s *sqlStore) Update(ctx context.Context, data *notemodel.Note, conditions map[string]interface{}, morekeys ...string) error {
	db := s.db
	if err := db.Table(notemodel.Note{}.TableName()).Where(conditions).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
