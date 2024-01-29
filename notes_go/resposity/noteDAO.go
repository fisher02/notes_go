package resposity

import (
	"notes_go/models"
	"notes_go/utils"
)

// 查询所有笔记
func SelectAllNote() ([]models.Note, error) {
	var notes []models.Note
	result := utils.DB.Find(&notes)
	return notes, result.Error
}

// 创建笔记
func CreateNote(note models.Note) (bool, error) {
	result := utils.DB.Create(&note)

	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// 根据Id查询笔记
func SelectNoteByID(id int) (*models.Note, error) {
	var note *models.Note
	result := utils.DB.First(&note, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return note, nil
}

// 根据id修改笔记
func UpdateNote(note models.Note) (bool, error) {
	result := utils.DB.Model(&note).Where("note_id=?", note.NoteId).Updates(models.Note{
		Title:   note.Title,
		Content: note.Content,
	})
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

// 根据id删除笔记
func DeleteNoteByID(id int) (bool, error) {
	result := utils.DB.Where("note_id=?", id).Delete(&models.Note{})
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
