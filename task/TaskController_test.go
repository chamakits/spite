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
	//TODO Check result of running and compare
	t.Logf("Able to set and run task.\n")
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
