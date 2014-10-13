package task

import (
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

func TestControllerCreation(t *testing.T) {
	// taskDao := &MapStoreDao{
	// 	Store: make(map[string]*TaskDataRuns, 0),
	// }
	// controller := NewController(taskDao)
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
