package task

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"time"
)

// TODO Should be configurable per process, but hardcoding it for now.

const LIMIT_PATH = "C:/"

var ABS_LIMIT_PATH string

func makePathWindowsFriendly(pathIn string) string {
	return strings.Replace(pathIn, `\`, `\\`, -1)
}

func init() {
	var err error
	ABS_LIMIT_PATH, err = filepath.Abs(LIMIT_PATH)
	// ABS_LIMIT_PATH = strings.Replace(ABS_LIMIT_PATH, `\`, `\\`, -1)
	//TODO Make OS check here.
	ABS_LIMIT_PATH = makePathWindowsFriendly(ABS_LIMIT_PATH)
	if err != nil {
		log.Fatalf("Failed ")
	}
}

// const LIMIT_PATH, ERROR_LIMIT_PATH = filepath.Abs("c:/go_tmp")

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

func equalsMapsStringString(leftMap, rightMap map[string]string) bool {
	if len(leftMap) != len(rightMap) {
		return false
	}

	for counter, _ := range leftMap {
		if leftMap[counter] != rightMap[counter] {
			return false
		}
	}

	return true
}

// Data contains the data of a Task.  Pretty much the HTTP post info.
type Data struct {
	FieldNameToValue map[string]string `json:"nameToData"`
}

type process struct {
	ExecutablePath string
	Arguments      []string
	// Put this in here so that I can specify a path limit for the actual executable
	// to not be found outside a certain bound.
	// THIS IS NOT A HUGE SECURITY BOOST. But it Should mitigate SOME
	pathLimit string
	// init      bool
}

func (proc *process) Equals(procOther *process) bool {
	return proc.ExecutablePath == procOther.ExecutablePath &&
		compareProcArguments(proc.Arguments, procOther.Arguments) &&
		proc.pathLimit == proc.pathLimit
}

func compareProcArguments(procLeft, procRight []string) bool {
	if len(procLeft) != len(procRight) {
		return false
	}
	for counter, _ := range procLeft {
		if procLeft[counter] != procRight[counter] {
			return false
		}
	}
	return true
}

//TODO This will probably need to change.  Realistically, its best to just
//not include this in the scope of the program for now.
func checkIfExecutableIsInAllowedScope(absoluteFilePath string) bool {
	re := regexp.MustCompile("^" + ABS_LIMIT_PATH)
	matchesPathLimit := re.FindStringSubmatch(absoluteFilePath)
	return matchesPathLimit == nil
}

func newProcess(path string, arguments []string) (*process, error) {
	// Make path absolute
	absoluteFilePath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	log.Println("Found file absolute path.")

	// Check if the absolute path is still within the limited file path.
	if runtime.GOOS == "windows" {
		absoluteFilePath = makePathWindowsFriendly(absoluteFilePath)
	}
	matchesPathLimit := checkIfExecutableIsInAllowedScope(absoluteFilePath)
	if matchesPathLimit {
		return nil, errors.New(fmt.Sprintf(
			"Executable given is off limited scope. Scope '%v', Given '%v'\n",
			ABS_LIMIT_PATH, absoluteFilePath))
	}
	log.Println("File for task within limited file path.")

	// Check if executable file exists
	fileInfo, err := os.Stat(absoluteFilePath)
	if os.IsNotExist(err) {
		return nil,
			errors.New(fmt.Sprintf("No such file or directory: %s", absoluteFilePath))
	}
	log.Println("File exists.")

	// Check if executable file is actually executable
	// fmt.Print(fileInfo)
	mode := fileInfo.Mode().Perm()
	if mode|0111 == 0 {
		return nil,
			errors.New(fmt.Sprintf("File '%s' is not an executable file", absoluteFilePath))
	}
	log.Println("File is executable.")

	return &process{
		ExecutablePath: absoluteFilePath,
		Arguments:      arguments,
		pathLimit:      ABS_LIMIT_PATH,
	}, nil
}

func (proc *process) runProcess(data Data) bytes.Buffer {
	log.Printf("Proc:%v\n", proc)
	command := exec.Command(proc.ExecutablePath, proc.Arguments...)

	var buff bytes.Buffer
	command.Stdout = &buff

	err := command.Run()
	if err != nil {
		//TODO Do actual handling of this.
		log.Fatalf("Error running process.  Error:%v\n", err)
	}
	return buff
}

// Task is used to represent a task, including schema and data.
type Task struct {
	View
	Schema
	proc *process
}

func (taskSelf *Task) Equals(taskOther *Task) bool {
	return taskSelf.ID == taskOther.ID &&
		taskSelf.Name == taskOther.Name &&
		taskSelf.Description == taskOther.Description &&
		equalsMapsStringString(taskSelf.FieldNameToDataType, taskOther.FieldNameToDataType) &&
		//TODO add equality for maps in schema.
		taskSelf.proc.Equals(taskOther.proc)
}

// Run method runs the task as specified with the data given.
func (taskSelf *Task) Run(data Data) bytes.Buffer {
	return taskSelf.proc.runProcess(data)
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

func (selfTask *Task) SetTaskProcess(path string, arguments []string) error {
	// selfTask.proc = newProcess(path)
	proc, err := newProcess(path, arguments)
	log.Printf("New proc created:%v\n", proc)
	selfTask.proc = proc
	return err
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
type ViewHTTP struct {
	View View `json:"view"`
}

type ViewsHTTP struct {
	Views []View `json:"taskViews"`
}

type ViewAndDataHTTP struct {
	View View `json:"view"`
	Data Data `json:"data"`
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
