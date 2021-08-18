package columns

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Boobuh/golang-school-project/service/columns"

	"github.com/Boobuh/golang-school-project/dal"
	"github.com/gorilla/mux"
)

type Handler struct {
	logger  *log.Logger
	service columns.Service
}

func NewHandler(service columns.Service, logger *log.Logger) *Handler {
	return &Handler{logger: logger, service: service}
}

//===========================================================================//

func (h *Handler) GetAllColumns(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new get request")

	getColumns, err := h.service.GetColumns()
	if err != nil {
		h.logger.Printf("error in GET getColumns call in service.GetColumns call:%s", err.Error())
		return
	}
	payload, err := json.Marshal(getColumns)
	if err != nil {
		h.logger.Printf("error in GET getColumns call - can't marshal object from db:%s", err.Error())
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
	columnIDRaw, ok := vars["columnID"]
	if !ok {
		http.Error(w, "id is missing in parameters", http.StatusBadRequest)
		h.logger.Println("id is missing in parameters")
	}
	columnID, err := strconv.Atoi(columnIDRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting id to int:%s", err.Error())
		return
	}
	projectIdRaw, ok := vars["projectID"]
	if !ok {
		http.Error(w, "id is missing in parameters", http.StatusBadRequest)
		h.logger.Println("id is missing in parameters")
	}
	projectID, err := strconv.Atoi(projectIdRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting id to int:%s", err.Error())
		return
	}
	column, err := h.service.GetProjectColumn(projectID, columnID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in receiving project by id:%s", err.Error())
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

func (h *Handler) CreateColumn(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new create request")
	vars := mux.Vars(r)
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
	var newColumn dal.Column
	err = json.NewDecoder(r.Body).Decode(&newColumn)
	if err != nil {
		h.logger.Printf("error in POST column call - can't decode object from request:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if newColumn.ProjectID != projectID {
		h.logger.Printf("error in POST column call - projectID mismatched")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.service.CreateColumn(&newColumn)
	if err != nil {
		h.logger.Printf("error in CREATE column call - %s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

//---------------------------------------------------------------------------//

func (h *Handler) DeleteColumn(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new delete column request")
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

	err = h.service.DeleteColumn(projectID, columnID)
	if err != nil {
		h.logger.Printf("error in DELETE column call:%s", err.Error())
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

//---------------------------------------------------------------------------//

func (h *Handler) UpdateColumn(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new UpdateColumn request")

	vars := mux.Vars(r)
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
	var updatedColumn dal.Column
	err = json.NewDecoder(r.Body).Decode(&updatedColumn)
	if err != nil {
		h.logger.Printf("error in POST column call - can't decode object from request:%s", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if updatedColumn.ID != columnID || updatedColumn.ProjectID != projectID {
		h.logger.Printf("error in PUT call columnID or projectID mismatched")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.service.UpdateColumn(&updatedColumn)

	if err != nil {
		h.logger.Printf("error in UPDATE column call - can't marshal object from db:%s", err.Error())
		return
	}
	w.WriteHeader(http.StatusOK)
}

//===========================================================================//

func (h *Handler) GetAllByProjectID(w http.ResponseWriter, r *http.Request) {
	h.logger.Print("new GetAllByProjectID request")

	vars := mux.Vars(r)
	projectIdRaw, ok := vars["projectID"]
	if !ok {
		http.Error(w, "id is missing in parameters", http.StatusBadRequest)
		h.logger.Println("id is missing in parameters")
	}
	projectID, err := strconv.Atoi(projectIdRaw)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in converting id to int:%s", err.Error())
		return
	}
	column, err := h.service.GetAllByProjectID(projectID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		h.logger.Printf("error in receiving project by id:%s", err.Error())
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
