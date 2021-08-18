package handler

import (
	"log"
	"net/http"

	"github.com/Boobuh/golang-school-project/dal"

	"github.com/Boobuh/golang-school-project/handler/columns"
	"github.com/Boobuh/golang-school-project/handler/comments"
	"github.com/Boobuh/golang-school-project/handler/projects"
	"github.com/Boobuh/golang-school-project/handler/tasks"

	columnsUseCase "github.com/Boobuh/golang-school-project/service/columns"
	commentUseCase "github.com/Boobuh/golang-school-project/service/comments"
	projectUseCase "github.com/Boobuh/golang-school-project/service/projects"
	taskUseCase "github.com/Boobuh/golang-school-project/service/tasks"

	"github.com/gorilla/mux"
)

func NewRouter(repo dal.Repository, logger *log.Logger) *mux.Router {
	router := mux.NewRouter()
	projectService := projectUseCase.NewUseCase(repo, logger)
	projectHandler := projects.NewHandler(projectService, logger)

	router.HandleFunc("/projects/", projectHandler.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/projects/{id}", projectHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/projects/", projectHandler.Create).Methods(http.MethodPost)
	router.HandleFunc("/projects/{id}", projectHandler.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/projects/{id}", projectHandler.Update).Methods(http.MethodPut)

	columnService := columnsUseCase.NewUseCase(repo, logger)
	columnHandler := columns.NewHandler(columnService, logger)

	router.HandleFunc("/columns/", columnHandler.GetAllColumns).Methods(http.MethodGet)
	router.HandleFunc("/projects/{projectID}/columns/", columnHandler.GetAllByProjectID).Methods(http.MethodGet)
	router.HandleFunc("/projects/{projectID}/columns/{columnID}", columnHandler.GetColumn).Methods(http.MethodGet)
	router.HandleFunc("/projects/{projectID}/columns/", columnHandler.CreateColumn).Methods(http.MethodPost)
	router.HandleFunc("/projects/{projectID}/columns/{columnID}", columnHandler.DeleteColumn).Methods(http.MethodDelete)
	router.HandleFunc("/projects/{projectID}/columns/{columnID}", columnHandler.UpdateColumn).Methods(http.MethodPut)

	taskService := taskUseCase.NewUseCase(repo, logger)
	taskHandler := tasks.NewHandler(taskService, logger)

	router.HandleFunc("/tasks/", taskHandler.GetAllTasks).Methods(http.MethodGet)
	//router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/", taskHandler.GetTask).Methods(http.MethodGet)
	router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}", taskHandler.GetTask).Methods(http.MethodGet)
	router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/", taskHandler.CreateTask).Methods(http.MethodPost)
	router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}", taskHandler.DeleteTask).Methods(http.MethodDelete)
	router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}", taskHandler.UpdateTask).Methods(http.MethodPut)

	commentService := commentUseCase.NewUseCase(repo, logger)
	commentHandler := comments.NewHandler(commentService, logger)

	router.HandleFunc("/comments/", commentHandler.GetAllComments).Methods(http.MethodGet)
	router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/{commentID}", commentHandler.GetComment).Methods(http.MethodGet)
	//router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/{commentID}", commentHandler.Get).Methods(http.MethodGet)
	router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/", commentHandler.CreateComment).Methods(http.MethodPost)
	router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/{commentID}", commentHandler.DeleteComment).Methods(http.MethodDelete)
	router.HandleFunc("/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/{commentID}", commentHandler.UpdateComment).Methods(http.MethodPut)

	return router
}
