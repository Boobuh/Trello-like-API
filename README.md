# Task Management Application

## Brief explanation

Golang School Project is the **back-end** (REST API) part of the task management application (like [Trello](https://www.youtube.com/watch?v=noguPYxyv6g)) with no authentication and authorization required.

## How to start

go run main.go

## How to test

Use Postman at http://127.0.0.1:4040/

List of all endpoints:

/projects/ GET
/projects/{id} GET 
/projects/ POST
/projects/ DELETE
/projects/{id} PUT

/columns/ GET
/projects/{projectID}/columns/ GET
/projects/{projectID}/columns/{columnID} GET
/projects/{projectID}/columns/ POST
/projects/{projectID}/columns/{columnID} DELETE
/projects/{projectID}/columns/{columnID} PUT

/tasks/ GET
/projects/{projectID}/columns/{columnID}/tasks/ GET
/projects/{projectID}/columns/{columnID}/tasks/{taskID} GET
/projects/{projectID}/columns/{columnID}/tasks/ POST
/projects/{projectID}/columns/{columnID}/tasks/{taskID} DELETE
/projects/{projectID}/columns/{columnID}/tasks/{taskID} PUT

/comments/ GET
/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/{commentID} GET
/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/ GET
/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/ POST
/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/{commentID} DELETE
/projects/{projectID}/columns/{columnID}/tasks/{taskID}/comments/{commentID} PUT

Or use swagger.yaml directly

## Link to cloud service
