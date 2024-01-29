package services

import (
	"notes_go/models"
	"notes_go/resposity"
)

// 获取所有笔记
func GetAllNotes() ([]models.Note, error) {
	return resposity.SelectAllNote()
}

// 新建笔记
func CreateNote(note models.Note) (bool, error) {
	return resposity.CreateNote(note)
}

// 根据ID查询笔记
func GetNoteByID(id int) (*models.Note, error) {
	return resposity.SelectNoteByID(id)
}

// 根据ID修改笔记
func ModifyNote(note models.Note) (bool, error) {
	return resposity.UpdateNote(note)
}

// 根据id删除笔记
func RemoveNoteByID(id int) (bool, error) {
	return resposity.DeleteNoteByID(id)
}
