package task

type Dao interface {
	GetTasksView() []View
	GetTaskDetail(View) Task
	GetTaskHistory(View) []RunInstance
}
