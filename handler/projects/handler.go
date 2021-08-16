package projects

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Boobuh/golang-school-project/dal"

	"github.com/gorilla/mux"

	"github.com/Boobuh/golang-school-project/service/projects"
)

type Handler struct {
	logger  *log.Logger
	service projects.Service
}

func NewHandler(service projects.Service, logger *log.Logger) *Handler {
	return &Handler{service: service, logger: logger}
}

//---------------------------------------------------------------------------//

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new get request")

	getProjects, err := h.service.GetProjects()
	payload, err := json.Marshal(getProjects)
	if err != nil {
		h.logger.Printf("error in GET getProjects call - can't marshal object from db:%s", err.Error())
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
	w.Header().Set("Content-Type", "application/json")
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
	project := h.service.CreateProject(&newProject)
	payload, err := json.Marshal(project)
	if err != nil {
		h.logger.Printf("error in CREATE projects call - can't marshal object from db:%s", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
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

	project := h.service.DeleteProject(id)
	payload, err := json.Marshal(project)
	if err != nil {
		h.logger.Printf("error in DELETE projects call - can't marshal object from db:%s", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
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

	db.First(&updatedProject)
	updatedProject.ID = id
	r.db.Save(&updatedProject)

	//updatedProject{ID: id}
	err := h.service.UpdateProject(r, updatedProject)

	//project := h.service.UpdateProject(r, id)
	payload, err := json.Marshal(updatedProject)
	if err != nil {
		h.logger.Printf("error in UPDATE projects call - can't marshal object from db:%s", err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}

//===========================================================================//

func (h *Handler) GetAllColumns(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new get request")

	getProjects, err := h.service.GetProjects()
	payload, err := json.Marshal(getProjects)
	if err != nil {
		h.logger.Printf("error in GET getProjects call - can't marshal object from db:%s", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}

//---------------------------------------------------------------------------//

func (h *Handler) GetColumn(w http.ResponseWriter, r *http.Request) {
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
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}

//---------------------------------------------------------------------------//

func (h *Handler) CreateColumns(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new create request")

	var newProject dal.Project
	err := json.NewDecoder(r.Body).Decode(&newProject)
	if err != nil {
		h.logger.Printf("error in POST project call - can't decode object from request:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	project := h.service.CreateProject(&newProject)
	payload, err := json.Marshal(project)
	if err != nil {
		h.logger.Printf("error in CREATE projects call - can't marshal object from db:%s", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusCreated)
}

//---------------------------------------------------------------------------//

func (h *Handler) DeleteColumns(w http.ResponseWriter, r *http.Request) {
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

	project := h.service.DeleteProject(id)
	payload, err := json.Marshal(project)
	if err != nil {
		h.logger.Printf("error in DELETE projects call - can't marshal object from db:%s", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusNoContent)
}

//---------------------------------------------------------------------------//

func (h *Handler) UpdateColumns(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new update request")
	var newProject dal.Project
	err := json.NewDecoder(r.Body).Decode(&newProject)
	if err != nil {
		h.logger.Printf("error in POST project call - can't decode object from request:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	project := h.service.UpdateProject(r)
	payload, err := json.Marshal(project)
	if err != nil {
		h.logger.Printf("error in UPDATE projects call - can't marshal object from db:%s", err.Error())
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(payload)
	w.WriteHeader(http.StatusOK)
}
