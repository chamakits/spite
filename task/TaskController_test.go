package task

import (
	"log"
	"strings"
	"testing"
)

func init() {
	log.SetFlags(log.Lshortfile)
}

func getNewMapStoreDaoController() *Controller {
	taskDao := &MapStoreDao{
		Store: make(map[string]*TaskDataRuns, 0),
	}
	return NewController(taskDao)
}

func TestControllerCreateTask(t *testing.T) {
	controller := getNewMapStoreDaoController()

	taskName := "NewTaskName"
	newTask := NewTask(0, taskName, "NewTaskDescription")
	controller.AddTask(*newTask)

	views := controller.GetTasksViews()
	previouslyInsertedView := views[0]
	if previouslyInsertedView.Name != taskName {
		t.Error("Failed to retrieve task with correct name.")
	} else {
		t.Logf("Found the correct task! Task view found:%v\n", previouslyInsertedView)
	}
}

func TestSetAndRunTask(t *testing.T) {
	taskName := "NewTaskName"
	newTask := NewTask(0, taskName, "NewTaskDescription")
	// err := newTask.SetTaskProcess(`C:\Windows\System32\cmd.exe`, []string{`/C`, `"ls"`})
	echoString := `some string`
	err := newTask.SetTaskProcess(`C:\cygwin64\bin\echo.exe`, []string{echoString})
	if err != nil {
		t.Errorf("Could not create new task process.  Error:%v\n", err)
		log.Fatalf("Test failed with error:%v\n", err)
	}
	log.Printf("New task proc%v\n", newTask.proc)
	buff := newTask.Run(Data{})
	buffString := buff.String()

	if !strings.Contains(buffString, echoString) {
		t.Errorf("Command ran, but buffer output not as expected. Found '%v' expected '%v'",
			buffString, echoString)
	}
	//TODO Check result of running and compare
	t.Logf("Able to set and run task, and get string result of process 'stdout'.\n")
}

func TestGetTaskDetails(t *testing.T) {
	controller := getNewMapStoreDaoController()

	taskName := "NewTaskName"
	newTask := NewTask(0, taskName, "NewTaskDescription")
	err := newTask.SetTaskProcess(`C:\cygwin64\bin\echo.exe`, []string{`some string`})
	if err != nil {
		t.Errorf("Could not create new task process.  Error:%v\n", err)
		log.Fatalf("Test failed with error:%v\n", err)
	}
	t.Logf("Created task\n")

	controller.AddTask(*newTask)
	taskView := (*newTask).CopyView()
	taskWithDetails := controller.GetTaskDetail(taskView)

	if !newTask.Equals(&taskWithDetails) {
		t.Errorf("Task retrieved is not the expected one. As created '%v' as retrieved '%v'\n",
			newTask, taskWithDetails)
	}

	t.Logf("Created and retrieved task succesfully by view.")
}
