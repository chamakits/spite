package task

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"
)

// TODO Should be configurable per process, but hardcoding it for now.
const LIMIT_PATH string = "c:/go_tmp"

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

type process struct {
	ExecutablePath string
	Flags          []string
	// Put this in here so that I can specify a path limit for the actual executable
	// to not be found outside a certain bound.
	// THIS IS NOT A HUGE SECURITY BOOST. But it Should mitigate SOME
	pathLimit string
	// init      bool
}

func newProcess(path string, flags []string) (*process, error) {
	// Make path absolute
	absoluteFilePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	// Check if the absolute path is still within the limited file path.
	re := regexp.MustCompile("^" + LIMIT_PATH)
	matchesPathLimit := re.FindStringSubmatch(absoluteFilePath)
	if matchesPathLimit == nil {
		return nil, errors.New("Executable given is off limited scope.")
	}

	// Check if executable file exists
	fileInfo, err := os.Stat(absoluteFilePath)
	if os.IsNotExist(err) {
		return nil,
			errors.New(fmt.Sprintf("No such file or directory: %s", absoluteFilePath))
	}

	// Check if executable file is actually executable
	// fmt.Print(fileInfo)
	mode := fileInfo.Mode().Perm()
	if mode|0111 == 0 {
		return nil,
			errors.New(fmt.Sprintf("File '%s' is not an executable file", absoluteFilePath))
	}

	return &process{
		ExecutablePath: absoluteFilePath,
		Flags:          flags,
		pathLimit:      LIMIT_PATH,
	}, nil
}

func (proc *process) runProcess(data Data) {

}

// Task is used to represent a task, including schema and data.
type Task struct {
	View
	Schema
	process
}

// Run method runs the task as specified with the data given.
func (taskSelf *Task) Run(data Data) {
	taskSelf.runProcess(data)
}

// NewTask creates a new task with name and description
func NewTask(id int, name, description string) *Task {
	newTask := &Task{
		View: View{
			ID:          id,
			Name:        name,
			Description: description,
		},
		Schema: Schema{
			FieldNameToDataType: make(map[string]string, 0),
		},
	}
	return newTask
}

// AddSchemaField adds detail to the schema.
func (taskSelf *Task) AddSchemaField(name, dataType string) {
	taskSelf.FieldNameToDataType[name] = dataType
}

// CopyView copies the view from the task.
func (taskSelf Task) CopyView() View {
	return View{
		ID:          taskSelf.ID,
		Name:        taskSelf.Name,
		Description: taskSelf.Description,
	}
}

// TaskHTTP is a 'task' just wrapped as one field to be used by http to send it
// as json over the wire.
type TaskHTTP struct {
	Task Task `json:"task"`
}

//TaskAndData is used mostly to retrieve the data from the Post request made.
type TaskAndData struct {
	Task Task `json:"task"`
	Data Data `json:"data"`
}

// DataHTTP used to wrap Data to be sent as json over the wire.
type DataHTTP struct {
	Data Data `json:"data"`
}

// RunInstance is used to represent a single instance of a running of a task.
type RunInstance struct {
	StartTime time.Time
	EndTime   time.Time
	Data      Data
}
