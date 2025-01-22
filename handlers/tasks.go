package handlers

import (
	"6ar8nas/test-app/util"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

var tasks = util.NewCache[int, string]()
var taskIdInc = util.AutoIncrement{}

func GetTask(w http.ResponseWriter, r *http.Request) {
	taskId, err := strconv.Atoi(r.PathValue("taskId"))
	if err != nil {
		http.Error(w, "Expected taskId to be an integer", http.StatusBadRequest)
		return
	}

	task, exists := tasks.Get(taskId)
	if !exists {
		http.Error(w, "Requested task was not found", http.StatusNotFound)
		return
	}

	resp := map[string]any{"taskId": taskId, "task": task}
	value, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(value)
}

func PostTask(w http.ResponseWriter, r *http.Request) {
	var id = taskIdInc.Id()

	var task string
	json.NewDecoder(r.Body).Decode(&task)

	tasks.Set(id, task)
	resp := map[string]any{"taskId": id, "task": task}
	value, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(value)
}
