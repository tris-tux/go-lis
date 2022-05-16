package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/tris-tux/go-lis/backend/db"
	"github.com/tris-tux/go-lis/backend/schema"
	"github.com/tris-tux/go-lis/backend/services"
)

type taskHandler struct {
	postgres *db.Postgres
	static   *db.Static
}

func (h *taskHandler) GetStatic(w http.ResponseWriter, r *http.Request) {
	ctx := db.SetRepo(r.Context(), h.static)

	taskList, err := service.GetAll(ctx)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseOK(w, taskList)
}

func (h *taskHandler) getAllTask(w http.ResponseWriter, r *http.Request) {
	if h.postgres == nil {
		responseError(w, http.StatusInternalServerError, "must connect to postgres")
		return
	}
	ctx := db.SetRepo(r.Context(), h.postgres)

	taskList, err := service.GetAll(ctx)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}
	responseOK(w, taskList)
}

func (h *taskHandler) insertTask(w http.ResponseWriter, r *http.Request) {
	if h.postgres == nil {
		responseError(w, http.StatusInternalServerError, "must connect to postgres")
		return
	}
	ctx := db.SetRepo(r.Context(), h.postgres)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var task schema.Task
	if err := json.Unmarshal(b, &task); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	task_id, err := service.Insert(ctx, &task)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOK(w, task_id)
}

func (h *taskHandler) updateTask(w http.ResponseWriter, r *http.Request) {
	if h.postgres == nil {
		responseError(w, http.StatusInternalServerError, "must connect to postgres")
		return
	}
	ctx := db.SetRepo(r.Context(), h.postgres)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var task schema.Task
	if err := json.Unmarshal(b, &task); err != nil {
		responseError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = service.Update(ctx, &task)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	responseOK(w, task.TaskId)
}

func (h *taskHandler) deleteTask(w http.ResponseWriter, r *http.Request) {
	if h.postgres == nil {
		responseError(w, http.StatusInternalServerError, "must connect to postgres")
		return
	}
	ctx := db.SetRepo(r.Context(), h.postgres)

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	var req struct {
		TaskId int `json:"task_id"`
	}
	if err := json.Unmarshal(b, &req); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := service.Delete(ctx, req.TaskId); err != nil {
		responseError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
}

func responseOK(w http.ResponseWriter, body interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func responseError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	body := map[string]string{
		"error": message,
	}
	json.NewEncoder(w).Encode(body)
}
