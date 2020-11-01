package web

import (
	"github.com/sfeir-open-source/sfeir-school-go/dao"
	"github.com/sfeir-open-source/sfeir-school-go/model"
	logger "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

const (
	prefix = "/tasks"
)

// TaskController is a controller for tasks resource
type TaskController struct {
	taskDao dao.TaskDAO
	Routes  []Route
	Prefix  string
}

// NewTaskController creates a new task controller to manage tasks
func NewTaskController(taskDAO dao.TaskDAO) *TaskController {
	controller := TaskController{
		taskDao: taskDAO,
		Prefix:  prefix,
	}

	// build routes
	routes := []Route{}
	// GetAll
	routes = append(routes, Route{
		Name:        "Get all tasks",
		Method:      http.MethodGet,
		Pattern:     "",
		HandlerFunc: controller.GetAll,
	})
	// Get
	routes = append(routes, Route{
		Name:        "Get one task",
		Method:      http.MethodGet,
		Pattern:     "/{id}",
		HandlerFunc: controller.Get,
	})
	// TODO add the create route as a Post with no ID, calling create method
	// Create

	// Update
	routes = append(routes, Route{
		Name:        "Update a task",
		Method:      http.MethodPut,
		Pattern:     "/{id}",
		HandlerFunc: controller.Update,
	})
	// Delete
	routes = append(routes, Route{
		Name:        "Delete a task",
		Method:      http.MethodDelete,
		Pattern:     "/{id}",
		HandlerFunc: controller.Delete,
	})

	controller.Routes = routes

	return &controller
}

// GetAll retrieve all entities with optional paging of items (start / end are item counts 50 to 100 for example)
func (tc *TaskController) GetAll(w http.ResponseWriter, r *http.Request) {

	startStr := QueryParamAsString("start", r)
	endStr := QueryParamAsString("end", r)

	start := dao.NoPaging
	end := dao.NoPaging
	var err error
	if startStr != "" && endStr != "" {
		start, err = strconv.Atoi(startStr)
		if err != nil {
			start = dao.NoPaging
		}
		end, err = strconv.Atoi(endStr)
		if err != nil {
			end = dao.NoPaging
		}
	}

	// find all tasks
	tasks, err := tc.taskDao.GetAll(start, end)
	if err != nil {
		if err == dao.ErrNotFound {
			logger.WithField("error", err).
				WithField("start", start).
				WithField("end", end).Warn("no tasks found")
			SendJSONWithHTTPCode(w, nil, http.StatusNoContent)
			return
		}
		logger.WithField("error", err).Warn("unable to retrieve tasks")
		SendJSONError(w, "Error while retrieving tasks", http.StatusInternalServerError)
		return
	}

	logger.WithField("tasks", tasks).Debug("tasks found")
	SendJSONOk(w, tasks)
}

// Get retrieve an entity by id
func (tc *TaskController) Get(w http.ResponseWriter, r *http.Request) {
	// get the task's ID from the URL
	taskID := URLParamAsString("id", r)

	// find the task
	task, err := tc.taskDao.GetByID(taskID)
	if err != nil {
		if err == dao.ErrNotFound {
			logger.WithField("error", err).WithField("task ID", taskID).Warn("unable to retrieve task by ID")
			SendJSONNotFound(w)
			return
		}

		logger.WithField("error", err).WithField("task ID", taskID).Warn("unable to retrieve task by ID")
		SendJSONError(w, "Error while retrieving task by ID", http.StatusInternalServerError)
		return
	}

	logger.WithField("tasks", task).Debug("task found")
	SendJSONOk(w, task)
}

// Create create an entity
func (tc *TaskController) Create(w http.ResponseWriter, r *http.Request) {
	// TODO declare a pointer to a task to be decoded
	// task to be created

	// TODO retrieve the content of the task in the request using GetJSONContent
	// get the content body

	// TODO if an error occurs, log it, send an error as a Bad Request and return

	// TODO if ok, save the task with the DAO
	// save task

	// TODO if error occurs, log it, send an error as an Internal Server Error and return

	// TODO if ok send the created task (with ID) and Status Created to the client with SendJSONWithHTTPCode
	// TODO send response
}

// Update update an entity by id
func (tc *TaskController) Update(w http.ResponseWriter, r *http.Request) {
	// get the task ID from the URL
	taskID := URLParamAsString("id", r)

	// task to be created
	task := model.Task{}
	// get the content body
	err := GetJSONContent(&task, r)

	if err != nil {
		logger.WithField("error", err).Warn("unable to decode task to update")
		SendJSONError(w, "Error while decoding task to update", http.StatusBadRequest)
		return
	}

	// save task
	task.ID = taskID
	task, err = tc.taskDao.Update(task)
	if err != nil {
		logger.WithField("error", err).WithField("task", task).Warn("unable to create task")
		SendJSONError(w, "Error while updating task", http.StatusInternalServerError)
		return
	}

	// send response
	SendJSONOk(w, task)
}

// Delete delete an entity by id
func (tc *TaskController) Delete(w http.ResponseWriter, r *http.Request) {
	// TODO get the task ID from the URL using URLParamAsString function

	// TODO delete the task by its ID

	// TODO if error, log it, send an error Internal Server Error and return

	// TODO log the successful deletion
	// TODO send an empty response to the client with No Content status
}
