package task

// Dao is the interface for all TaskDaos.
type Dao interface {
	GetTasksView() []View
	GetTaskDetail(View) Task
	GetTaskHistory(View) []RunInstance
	AddTask(Task)
	// AddRunInstance(RunInstance)
	AddTaskRun(taskName string, newRunInstance RunInstance)
}

// TaskDataRuns is a struct that contains a task along with its run instances.
type TaskDataRuns struct {
	Task         Task
	RunInstances []RunInstance
}

//This is EXTREMELY not thread safe right now.
type MapStoreDao struct {
	Store map[string]*TaskDataRuns
}

// GetTasksView returns a small view of all the tasks.
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
	dao.Store[taskIn.Name] = &TaskDataRuns{
		Task:         taskIn,
		RunInstances: make([]RunInstance, 0),
	}
}

func (dao *MapStoreDao) AddTaskRun(
	taskName string, newRunInstance RunInstance) {
	// startTime := time.Now()
	// dao.Store[view.Name].

	runInstances := append(
		dao.Store[taskName].RunInstances,
		newRunInstance,
	)
	dao.Store[taskName].RunInstances = runInstances
}
