package task

import "time"

// View is used to represent a simple high level view of a Task.
type View struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Schema is used to represent what data a Task contains.
type Schema struct {
	// TODO probably need to rethink this.
	FieldNameToDataType map[string]string `json:"nameToType"`
}

// Data contains the data of a Task.  Pretty much the HTTP post info.
type Data struct {
	FieldNameToValue map[string]string `json:"nameToData"`
}

// Task is used to represent a task, including schema and data.
type Task struct {
	View
	Schema
}

func (taskSelf Task) CopyView() View {
	return View{
		ID:          taskSelf.ID,
		Name:        taskSelf.Name,
		Description: taskSelf.Description,
	}
}

type TaskHttp struct {
	Task Task `json:"task"`
}

//TaskAndData is used mostly to retrieve the data from the Post request made.
type TaskAndData struct {
	Task Task `json:"task"`
	Data Data `json:"data"`
}

type DataHttp struct {
	Data Data `json:"data"`
}

// RunInstance is used to represent a single instance of a running of a task.
type RunInstance struct {
	StartTime time.Time
	EndTime   time.Time
	Data      Data
}
