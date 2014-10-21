package task

import (
	"log"
	"time"
)

type Controller struct {
	Dao Dao
}

func NewController(dao Dao) *Controller {
	return &Controller{dao}
}

func (taskController *Controller) GetTasksViews() []View {
	return taskController.Dao.GetTasksViews()
}
func (taskController *Controller) GetTaskDetail(taskView View) Task {
	return taskController.Dao.GetTaskDetail(taskView)
}
func (taskController *Controller) GetTaskHistory(taskView View) []RunInstance {
	return taskController.Dao.GetTaskHistory(taskView)
}

func (taskController *Controller) AddTask(taskIn Task) {
	taskController.Dao.AddTask(taskIn)
}

func (taskController *Controller) RunTask(taskView View, data Data) {
	startTime := time.Now()
	log.Printf("taskView:%v\n", taskView)
	taskFound := taskController.GetTaskDetail(taskView)
	taskFound.Run(data)

	endTime := time.Now()

	newRunInstance := RunInstance{
		StartTime: startTime,
		EndTime:   endTime,
		Data:      data,
	}

	taskController.Dao.AddTaskRun(taskFound.Name, newRunInstance)

}
