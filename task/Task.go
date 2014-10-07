package task

import "time"

// View is used to represent a simple high level view of a Task.
type View struct {
	ID int
}

// Schema is used to represent what data a Task contains.
type Schema struct {
	// TODO probably need to rethink this.
	fieldNameToDataType map[string]string
}

// Data contains the data of a Task
type Data struct {
	fieldNameToValue map[string]string
}

// Task is used to represent a task, including schema and data.
type Task struct {
	View
	Schema
	Data
}

// RunInstance is used to represent a single instance of a running of a task.
type RunInstance struct {
	StartTime time.Time
	EndTime   time.Time
}
