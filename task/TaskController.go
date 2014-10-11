package task

type Controller struct {
	Dao Dao
}

func NewController(dao Dao) *Controller {
	return &Controller{dao}
}

func (taskController *Controller) GetTasksView() []View {
	return taskController.Dao.GetTasksView()
}
func (taskController *Controller) GetTaskDetail(taskView View) Task {
	return taskController.Dao.GetTaskDetail(taskView)
}
func (taskController *Controller) GetTaskHistory(taskView View) []RunInstance {
	return taskController.Dao.GetTaskHistory(taskView)

}
