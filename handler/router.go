package handler

import (
	"log"
	"net/http"

	"github.com/Boobuh/golang-school-project/dal"

	"github.com/Boobuh/golang-school-project/handler/projects"

	projectUseCase "github.com/Boobuh/golang-school-project/service/projects"

	"github.com/gorilla/mux"
)

func NewRouter(repo dal.Repository, logger *log.Logger) *mux.Router {
	router := mux.NewRouter()
	projectService := projectUseCase.NewUseCase(repo, logger)
	projectHandler := projects.NewHandler(projectService, logger)
	//----------------------------------------------------------------------------------------------//
	router.HandleFunc("/projects/", projectHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/projects/{id}", projectHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/projects/", projectHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/projects/{id}", projectHandler.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/projects/{id}", projectHandler.Update).Methods(http.MethodPut)
	//----------------------------------------------------------------------------------------------//
	return router
}
