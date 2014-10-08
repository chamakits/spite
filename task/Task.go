package task

import "time"

// View is used to represent a simple high level view of a Task.
type View struct {
	ID          int
	Name        string
	Description string
}

// Schema is used to represent what data a Task contains.
type Schema struct {
	// TODO probably need to rethink this.
	fieldNameToDataType map[string]string
}

// Data contains the data of a Task.  Pretty much the HTTP post info.
type Data struct {
	fieldNameToValue map[string][]byte
}

// Task is used to represent a task, including schema and data.
type Task struct {
	View
	Schema
}

// RunInstance is used to represent a single instance of a running of a task.
type RunInstance struct {
	StartTime time.Time
	EndTime   time.Time
}
