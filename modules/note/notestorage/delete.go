package notestorage

import (
	"FirstProject/modules/note/notemodel"
	"context"
)

func (s *sqlStore) Delete(ctx context.Context, conditions map[string]interface{}, morekeys ...string) ([]notemodel.Note, error) {
	var result []notemodel.Note
	db := s.db

	db = db.Table(notemodel.Note{}.TableName()).Where(conditions)

	if err := db.Delete(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
