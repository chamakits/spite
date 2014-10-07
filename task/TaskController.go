package task

type TaskController interface {
	GetTasksView() []View
	GetTaskDetail(taskView View) Task
	GetTaskHistory(taskView View) []RunInstance
	// AddT
}
