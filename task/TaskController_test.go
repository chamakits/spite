package task

import (
	"fmt"
	"log"
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
	err := newTask.SetTaskProcess(`C:\cygwin64\bin\echo.exe`, []string{`some string`})
	if err != nil {
		t.Errorf("Could not create new task process.  Error:%v\n", err)
		log.Fatalf("Test failed with error:%v\n", err)
	}
	log.Printf("New task proc%v\n", newTask.proc)
	buff := newTask.Run(Data{})
	fmt.Printf("Ran command.  Buff out is:%v\n", buff.String())
}
