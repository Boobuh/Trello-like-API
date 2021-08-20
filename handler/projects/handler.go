package projects

//go:generate   $GOPATH/bin/mockgen -package mocks -destination=mocks/mock_service.go -package=mocks github.com/Boobuh/golang-school-project/handler/projects Service

import (
	"encoding/json"
	"fmt"
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

const (
	contentTypeHeader = "Content-Type"
	jsonContentType   = "application/json; charset=utf-8"
)

type Service interface {
	//--------------------------------------------------------------//
	GetProjects() ([]dal.Project, error)
	GetProject(id int) (*dal.ExtendedProjectEntities, error)
	CreateProject(project *dal.Project) error
	DeleteProject(id int) error
	UpdateProject(updatedProject *dal.Project) error
	//--------------------------------------------------------------//

}

func NewHandler(service Service, logger *log.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

//===========================================================================//

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new get request")

	getProjects, err := h.service.GetProjects()
	if err != nil {
		h.logger.Printf("error in GET getProjects call:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	payload, err := json.Marshal(getProjects)
	if err != nil {
		h.logger.Printf("error in GET getProjects call - can't marshal object from db:%s", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}

//---------------------------------------------------------------------------//

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new get request")

	vars := mux.Vars(r)
	idRaw, ok := vars["id"]
	if !ok {
		http.Error(w, "id is missing in parameters", http.StatusBadRequest)
		h.logger.Println("id is missing in parameters")
	}
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting id to int:%s", err.Error())
		return
	}
	project, err := h.service.GetProject(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in receiving project by id:%s", err.Error())
		return
	}
	payload, err := json.Marshal(project)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in GET projects call - can't marshal object from db:%s", err.Error())
		return
	}
	//w.Header().Set("Content-Type", "application/json")

	w.Header().Set(contentTypeHeader, jsonContentType)
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}

//---------------------------------------------------------------------------//

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new create request")

	var newProject dal.Project
	err := json.NewDecoder(r.Body).Decode(&newProject)
	if err != nil {
		h.logger.Printf("error in POST project call - can't decode object from request:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.CreateProject(&newProject)
	if err != nil {
		h.logger.Printf("error in CREATE projects call:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	w.WriteHeader(http.StatusCreated)
}

//---------------------------------------------------------------------------//

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new create request")
	vars := mux.Vars(r)
	idRaw, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	id, err := strconv.Atoi(idRaw)
	if err != nil {
		h.logger.Printf("error in converting id to int:%s", err.Error())
		return
	}

	err = h.service.DeleteProject(id)
	if err != nil {
		h.logger.Printf("error in DELETE projects call:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//---------------------------------------------------------------------------//

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new update request")

	vars := mux.Vars(r)
	idRaw, ok := vars["id"]
	if !ok {
		fmt.Println("id is missing in parameters")
	}
	id, errConv := strconv.Atoi(idRaw)
	if errConv != nil {
		h.logger.Printf("error in converting id to int:%s", errConv.Error())
		return
	}

	var updatedProject dal.Project
	err := json.NewDecoder(r.Body).Decode(&updatedProject)
	if err != nil {
		h.logger.Printf("error in POST project call - can't decode object from request:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)

		return
	}
	updatedProject.ID = id
	err = h.service.UpdateProject(&updatedProject)

	if err != nil {
		h.logger.Printf("error in UPDATE projects call - can't marshal object from db:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message":"success"}`))
	w.WriteHeader(http.StatusOK)
}
