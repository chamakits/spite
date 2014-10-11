package task

type Dao interface {
	GetTasksView() []View
	GetTaskDetail(View) Task
	GetTaskHistory(View) []RunInstance
}

type TaskDataRuns struct {
	Task Task
	Data Data
	Runs []RunInstance
}

type MapStoreDao struct {
	Store map[int]TaskDataRuns
}

func (dao *MapStoreDao) GetTasksView() []View {
	views := make([]View, len(dao.Store))
	counter := 0
	for _, value := range dao.Store {
		views[counter] = value.Task.CopyView()
		counter++
	}
	return views
}
