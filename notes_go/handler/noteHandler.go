package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"notes_go/models"
	"notes_go/services"
	"strconv"

	"github.com/gorilla/mux"
)

// 获取所有笔记处理函数
func GetAllNotesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	result, err := services.GetAllNotes()
	if err != nil {
		json.NewEncoder(w).Encode(models.R{
			Success: false,
			Data:    nil,
			Message: fmt.Sprint("Cannot get Note list"),
		})
		return
	}
	json.NewEncoder(w).Encode(models.R{
		Success: true,
		Data:    result,
		Message: "Get All Notes Success!",
	})
}

// 创建笔记
func CreateNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var note models.Note
	jerr := json.NewDecoder(r.Body).Decode(&note)
	if jerr != nil {
		json.NewEncoder(w).Encode(models.R{
			Success: false,
			Data:    nil,
			Message: fmt.Sprint("Json Decode error:%v", jerr),
		})
		return
	}
	result, err := services.CreateNote(note)
	if err != nil {
		json.NewEncoder(w).Encode(models.R{
			Success: false,
			Data:    nil,
			Message: fmt.Sprint("Cannot Create Note"),
		})
		return
	}
	json.NewEncoder(w).Encode(models.R{
		Success: true,
		Data:    result,
		Message: "Create Note Success!",
	})
}

// 根据ID查询笔记
func GetNoteByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var note *models.Note
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	note, err := services.GetNoteByID(id)
	if err != nil {
		json.NewEncoder(w).Encode(models.R{
			Success: false,
			Data:    nil,
			Message: fmt.Sprint("Cannot Get Note By ID"),
		})
		return
	}
	json.NewEncoder(w).Encode(models.R{
		Success: true,
		Data:    note,
		Message: "Get Note Success!",
	})
}

// 根据ID修改笔记
func ModifyNoteHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var note models.Note
	jerr := json.NewDecoder(r.Body).Decode(&note)
	if jerr != nil {
		json.NewEncoder(w).Encode(models.R{
			Success: false,
			Data:    nil,
			Message: fmt.Sprint("Json Decode error:%v", jerr),
		})
		return
	}
	result, err := services.ModifyNote(note)
	if err != nil {
		json.NewEncoder(w).Encode(models.R{
			Success: false,
			Data:    nil,
			Message: fmt.Sprint("Cannot Modify Note!"),
		})
		return
	}
	json.NewEncoder(w).Encode(models.R{
		Success: true,
		Data:    result,
		Message: "Update Note Success!",
	})
}

// 根据ID删除笔记
func RemoveNoteByIDHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	result, err := services.RemoveNoteByID(id)
	if err != nil {
		json.NewEncoder(w).Encode(models.R{
			Success: false,
			Data:    nil,
			Message: fmt.Sprint("Cannot remove Note!"),
		})
		return
	}
	json.NewEncoder(w).Encode(models.R{
		Success: true,
		Data:    result,
		Message: "Delete Note Success!",
	})
}
