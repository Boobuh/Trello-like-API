package comments

//go:generate   $GOPATH/bin/mockgen -package mocks -destination=mocks/mock_service.go -package=mocks github.com/Boobuh/golang-school-project/handler/comments Service

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Boobuh/golang-school-project/dal"
	"github.com/gorilla/mux"
)

type Handler struct {
	logger  *log.Logger
	service Service
}

type Service interface {
	GetComments() ([]dal.Comment, error)
	GetComment(projectID, columnID, taskID, commentID int) (*dal.Comment, error)
	CreateComment(task *dal.Comment) error
	DeleteComment(projectID, columnID, taskID, commentID int) error
	UpdateComment(task *dal.Comment) error
	GetAllByTaskID(taskID int) ([]dal.Comment, error)
}

func NewHandler(service Service, logger *log.Logger) *Handler {
	return &Handler{logger: logger, service: service}
}

//===========================================================================//

func (h *Handler) GetAllComments(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new GetAllComments request")

	getComments, err := h.service.GetComments()
	if err != nil {
		h.logger.Printf("error in GET getComments call:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	payload, err := json.Marshal(getComments)
	if err != nil {
		h.logger.Printf("error in GET getComments call - can't marshal object from db:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusOK)

}

//---------------------------------------------------------------------------//

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new GetComment request")

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
	commentIdRaw, ok := vars["commentID"]
	if !ok {
		http.Error(w, "commentID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("commentID is missing in parameters")
	}
	commentID, err := strconv.Atoi(commentIdRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting taskID to int:%s", err.Error())
		return
	}

	task, err := h.service.GetComment(projectID, columnID, taskID, commentID)
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

func (h *Handler) CreateComment(w http.ResponseWriter, r *http.Request) {

	h.logger.Print("new CreateComment request")
	vars := mux.Vars(r)
	taskIDRaw, ok := vars["taskID"]
	if !ok {
		http.Error(w, "taskID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("taskID is missing in parameters")
	}
	taskID, err := strconv.Atoi(taskIDRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting id to int:%s", err.Error())
		return
	}

	var newComment dal.Comment
	err = json.NewDecoder(r.Body).Decode(&newComment)
	if err != nil {
		h.logger.Printf("error in POST Comment call - can't decode object from request:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if newComment.TaskID != taskID {
		h.logger.Printf("error in POST task call - columnID mismatched")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.CreateComment(&newComment)
	if err != nil {
		h.logger.Printf("error in CREATE task call - %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

//---------------------------------------------------------------------------//

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new delete comment request")
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

	commentIdRaw, ok := vars["commentID"]
	if !ok {
		http.Error(w, "commentID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("commentID is missing in parameters")
	}
	commentID, err := strconv.Atoi(commentIdRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting commentID to int:%s", err.Error())
		return
	}

	err = h.service.DeleteComment(projectID, columnID, taskID, commentID)
	if err != nil {
		h.logger.Printf("error in DELETE comment call:%s", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//---------------------------------------------------------------------------//

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new UpdateComment request")

	vars := mux.Vars(r)
	columnIDRaw, ok := vars["columnID"]
	if !ok {
		http.Error(w, "columnID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("columnID is missing in parameters")
	}
	_, err := strconv.Atoi(columnIDRaw)
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
	commentIdRaw, ok := vars["commentID"]
	if !ok {
		http.Error(w, "commentID is missing in parameters", http.StatusBadRequest)
		h.logger.Println("commentID is missing in parameters")
	}
	commentID, err := strconv.Atoi(commentIdRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting commentID to int:%s", err.Error())
		return
	}
	var updatedComment dal.Comment
	err = json.NewDecoder(r.Body).Decode(&updatedComment)
	if err != nil {
		h.logger.Printf("error in PUT comment call - can't decode object from request:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if updatedComment.TaskID != taskID || updatedComment.ID != commentID {
		h.logger.Printf("error in PUT call taskID or commentID mismatched")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.UpdateComment(&updatedComment)

	if err != nil {
		h.logger.Printf("error in UPDATE comment call - can't marshal object from db:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

//---------------------------------------------------------------------------//

func (h *Handler) GetAllByTaskID(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new GetAllByTaskID request")

	vars := mux.Vars(r)
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
	task, err := h.service.GetAllByTaskID(taskID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in receiving tasks by columnID:%s", err.Error())
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

//===========================================================================//
