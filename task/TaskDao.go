package task

type Dao interface {
	GetTasksView() []View
	GetTaskDetail(View) Task
	GetTaskHistory(View) []RunInstance
	AddTask(Task)
	AddRunInstance(RunInstance)
}

type TaskDataRuns struct {
	Task         Task
	RunInstances []RunInstance
}

//This is EXTREMELY not thread safe right now.
type MapStoreDao struct {
	Store map[string]TaskDataRuns
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

func (dao *MapStoreDao) GetTaskDetail(view View) Task {
	return dao.Store[view.Name].Task
}

func (dao *MapStoreDao) GetTaskHistory(view View) []RunInstance {
	return dao.Store[view.Name].RunInstances
}

func (dao *MapStoreDao) AddTask(taskIn Task) {
	dao.Store[taskIn.Name] = TaskDataRuns{
		Task:         taskIn,
		RunInstances: make([]RunInstance, 0),
	}

}