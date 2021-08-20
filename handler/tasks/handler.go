package tasks

//go:generate   $GOPATH/bin/mockgen -package mocks -destination=mocks/mock_service.go -package=mocks github.com/Boobuh/golang-school-project/handler/tasks Service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Boobuh/golang-school-project/dal"
	"github.com/gorilla/mux"
)

type Service interface {
	GetTasks() ([]dal.Task, error)
	GetTask(projectID, columnID, taskID int) (*dal.ExtendedTask, error)
	CreateTask(task *dal.Task) error
	DeleteTask(projectID, columnID, taskID int) error
	UpdateTask(task *dal.Task) error
	GetAllByColumnID(columnID int) ([]dal.ExtendedTask, error)
}

type Handler struct {
	logger  *log.Logger
	service Service
}

func NewHandler(service Service, logger *log.Logger) *Handler {
	return &Handler{logger: logger, service: service}
}

func (h *Handler) GetAllTasks(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new get request")

	getTasks, err := h.service.GetTasks()
	if err != nil {
		h.logger.Printf("error in GET getColumns call:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	payload, err := json.Marshal(getTasks)
	if err != nil {
		h.logger.Printf("error in GET getColumns call - can't marshal object from db:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}

//---------------------------------------------------------------------------//

func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new GetTask request")

	vars := mux.Vars(r)
	columnIDRaw, ok := vars["columnID"]
	if !ok {
		http.Error(w, "columnID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("columnID is missing in parameters")
	}
	columnID, err := strconv.Atoi(columnIDRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting id to int:%s", err.Error())
		return
	}
	projectIdRaw, ok := vars["projectID"]
	if !ok {
		http.Error(w, "projectID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("projectID is missing in parameters")
	}
	projectID, err := strconv.Atoi(projectIdRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting id to int:%s", err.Error())
		return
	}
	taskIdRaw, ok := vars["taskID"]
	if !ok {
		http.Error(w, "taskID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("taskID is missing in parameters")
	}
	taskID, err := strconv.Atoi(taskIdRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting taskID to int:%s", err.Error())
		return
	}

	task, err := h.service.GetTask(projectID, columnID, taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in receiving task by id:%s", err.Error())
		return
	}
	payload, err := json.Marshal(task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in GET task call - can't marshal object from db:%s", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}

//---------------------------------------------------------------------------//

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {

	h.logger.Print("new create task request")
	vars := mux.Vars(r)
	columnIDRaw, ok := vars["columnID"]
	if !ok {
		http.Error(w, "columnID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("columnID is missing in parameters")
	}
	columnID, err := strconv.Atoi(columnIDRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting id to int:%s", err.Error())
		return
	}

	var newTask dal.Task
	err = json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		h.logger.Printf("error in POST column call - can't decode object from request:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if newTask.ColumnID != columnID {
		h.logger.Printf("error in POST task call - columnID mismatched")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.CreateTask(&newTask)
	if err != nil {
		h.logger.Printf("error in CREATE task call - %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

//---------------------------------------------------------------------------//

func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new delete task request")
	vars := mux.Vars(r)
	columnIDRaw, ok := vars["columnID"]
	if !ok {
		http.Error(w, "columnID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("columnID is missing in parameters")
	}
	columnID, err := strconv.Atoi(columnIDRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting id to int:%s", err.Error())
		return
	}
	projectIdRaw, ok := vars["projectID"]
	if !ok {
		http.Error(w, "projectID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("projectID is missing in parameters")
	}
	projectID, err := strconv.Atoi(projectIdRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting projectID to int:%s", err.Error())
		return
	}
	taskIdRaw, ok := vars["taskID"]
	if !ok {
		http.Error(w, "taskID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("taskID is missing in parameters")
	}
	taskID, err := strconv.Atoi(taskIdRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting taskID to int:%s", err.Error())
		return
	}

	err = h.service.DeleteTask(projectID, columnID, taskID)
	if err != nil {
		h.logger.Printf("error in DELETE task call:%s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//---------------------------------------------------------------------------//

func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new UpdateTask request")

	vars := mux.Vars(r)
	columnIDRaw, ok := vars["columnID"]
	if !ok {
		http.Error(w, "columnID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("columnID is missing in parameters")
	}
	columnID, err := strconv.Atoi(columnIDRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting id to int:%s", err.Error())
		return
	}
	taskIdRaw, ok := vars["taskID"]
	if !ok {
		http.Error(w, "taskID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("taskID is missing in parameters")
	}
	taskID, err := strconv.Atoi(taskIdRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting taskID to int:%s", err.Error())
		return
	}
	var updatedTask dal.Task
	err = json.NewDecoder(r.Body).Decode(&updatedTask)
	if err != nil {
		h.logger.Printf("error in PUT Task call - can't decode object from request:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if updatedTask.ColumnID != columnID || updatedTask.ID != taskID {
		h.logger.Printf("error in PUT call columnID or taskID mismatched")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.UpdateTask(&updatedTask)

	if err != nil {
		h.logger.Printf("error in UPDATE task call - can't marshal object from db:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//---------------------------------------------------------------------------//

func (h *Handler) GetAllByColumnID(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new GetAllByProjectID request")

	vars := mux.Vars(r)
	columnIDRaw, ok := vars["columnID"]
	if !ok {
		http.Error(w, "columnID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("columnID is missing in parameters")
	}
	columnID, err := strconv.Atoi(columnIDRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting id to int:%s", err.Error())
		return
	}
	column, err := h.service.GetAllByColumnID(columnID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in receiving tasks by columnID:%s", err.Error())
		return
	}
	payload, err := json.Marshal(column)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in GET projects call - can't marshal object from db:%s", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}

//---------------------------------------------------------------------------//
