package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	todo_app "study/todo_app"
	"time"

	"github.com/gorilla/mux"
)

type HTTPHandlers struct {
	todoList *todo_app.List
}

func NewHTTPHandlers(todoList *todo_app.List) *HTTPHandlers {
	return &HTTPHandlers{
		todoList: todoList,
	}

}

/*
pattern: /tasks
method: POST
info: JSON in HTTP request body

succeed:
-status code: 201 Created
-response body: JSON represent created task

failed:
-status code: 400, 409, 500..
-response body: JSON with error + time


*/

func (h *HTTPHandlers) HandlerCreateTask(w http.ResponseWriter, r *http.Request) {
	var taskDTO TaskDTO
	if err := json.NewDecoder(r.Body).Decode(&taskDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return

	}
	if err := taskDTO.ValidateForCreate(); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
	}
	todoTask := todo_app.NewTask(taskDTO.Title, taskDTO.Description)

	if err := h.todoList.AddTask(todoTask); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo_app.ErrTaskAlreadyExists) {
			http.Error(w, errDTO.ToString(), http.StatusConflict)

		} else {

			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(todoTask, "", "    ")
	if err != nil {
		panic(err)

	}
	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write to http response:", err)
		return
	}
}

/*
pattern: /tasks/{title}
method: GET        ^
info: pattern-------

succeed:
-status code: 200 Ok
-response body: JSON represented found task

failed:
-status code: 400, 404, 500..
-response body: JSON with error + time


*/

func (h *HTTPHandlers) HandlerGetTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	task, err := h.todoList.GetTask(title)
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		if errors.Is(err, todo_app.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)
		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)
		}
		return
	}
	b, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}
}

/*
pattern: /tasks
method: GET
info:

succeed:
-status code: 200 Ok
-response body: JSON represented found task

failed:
-status code: 400, 500..
-response body: JSON with error + time


*/

func (h *HTTPHandlers) HandlerGetAllTasks(w http.ResponseWriter, r *http.Request) {
	tasks := h.todoList.ListTasks()
	b, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response")
		return
	}
}

/*
pattern: /tasks?completed=true
method: GET        ^
info: query params--

succeed:
-status code:
-response body:

failed:
-status code:
-response body:


*/

func (h *HTTPHandlers) HandleGetAllUncompletedRTasks(w http.ResponseWriter, r *http.Request) {
	uncompletedTasks := h.todoList.ListUnCompletedTasks()
	b, err := json.MarshalIndent(uncompletedTasks, "", "    ")
	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response")
		return
	}
}

/*
pattern: /tasks/{title}
method: PATCH
info:  pattern + JSON in request body

succeed:
-status code: 200 OK
-response body: JSON represented changed Task

failed:
-status code: 400, 409, 500...
-response body: JSON with error + time
*/

func (h *HTTPHandlers) HandleComleteTask(w http.ResponseWriter, r *http.Request) {
	var completedDTO CompleteTaskDto

	if err := json.NewDecoder(r.Body).Decode(&completedDTO); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}
		http.Error(w, errDTO.ToString(), http.StatusBadRequest)
		return

	}
	title := mux.Vars(r)["title"]

	var (
		changedTask todo_app.Task
		err         error
	)
	if completedDTO.Complete {
		changedTask, err = h.todoList.CompleteTask(title)

	} else {
		changedTask, err = h.todoList.UnCompleteTask(title)

	}
	if err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, todo_app.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)

		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)

		}
		return

	}
	b, err := json.MarshalIndent(changedTask, "", "    ")
	if err != nil {
		panic(err)
	}
	if _, err := w.Write(b); err != nil {
		fmt.Println("failed to write http response", err)
		return
	}

}

/*
pattern: /tasks/{title}
method: DELETE
info:  pattern

succeed:
-status code: 204 No Content
-response body:

failed:
-status code: 400, 404, 500...
-response body: JSON with error + time
*/

func (h *HTTPHandlers) HandleDeletedTask(w http.ResponseWriter, r *http.Request) {
	title := mux.Vars(r)["title"]
	if err := h.todoList.DeleteTask(title); err != nil {
		errDTO := ErrorDTO{
			Message: err.Error(),
			Time:    time.Now(),
		}

		if errors.Is(err, todo_app.ErrTaskNotFound) {
			http.Error(w, errDTO.ToString(), http.StatusNotFound)

		} else {
			http.Error(w, errDTO.ToString(), http.StatusInternalServerError)

		}
		return

	}
	w.WriteHeader(http.StatusNoContent)

}
