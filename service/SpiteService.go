package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/chamakits/spite/task"
	"github.com/gorilla/mux"
)

// SpiteService is a service that hosts Spite server.
// TODO I'll need to handle multipart upload. See here how:
// http://sanatgersappa.blogspot.com/2013/03/handling-multiple-file-uploads-in-go.html
// https://www.socketloop.com/tutorials/golang-upload-file
type SpiteService struct {
	Port           int
	taskController *task.Controller
}

func (spiteService *SpiteService) Init() {
	if spiteService.Port == 0 {
		spiteService.Port = 9090
	}

	spiteService.initImplementedController()

	// r := mux.NewRouter()
	// r.HandleFunc("/hello/", acceptCors(helloHandler(spiteService)))
	//

	r := mux.NewRouter()
	r.HandleFunc("/hello/", acceptCors(helloHandler(spiteService)))
	r.HandleFunc("/api/add-task", acceptCors(addTaskHandler(spiteService)))
	r.HandleFunc("/api/run-task", acceptCors(runTaskHandler(spiteService)))
	// TODO need to create a new handler function for showing tasks
	r.HandleFunc("/api/show-tasks", acceptCors(showTasksHandler(spiteService)))
	r.HandleFunc("/api/get-task-detail", acceptCors(getTaskDetailHandler(spiteService)))

	spiteService.initHTTP(r)

}

func (spiteService *SpiteService) initHTTP(router *mux.Router) {
	http.Handle("/", router)
	http.ListenAndServe(fmt.Sprintf(":%v", spiteService.Port), nil)
}

func corsEnable(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	//fmt.Println("corsEnable")
}

func acceptCors(handlerFunction http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		log.Printf("Method called:%v\n", req.Method)
		corsEnable(&response)
		if "OPTIONS" == req.Method {
			return
		} else {
			handlerFunction(response, req)
		}
	}
}

func addTaskHandlerSTRING(spiteService *SpiteService) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		bytes, error := ioutil.ReadAll(req.Body)
		if error != nil {
			fmt.Println("Error")
			log.Fatal(error)
		}
		result := string(bytes)
		log.Printf("From req directly:%v\n", result)
	}
}

func runTaskHandler(spiteService *SpiteService) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		log.Printf("Run task handler.")
		decoder := json.NewDecoder(req.Body)
		var viewAndData task.ViewAndDataHTTP
		err := decoder.Decode(&viewAndData)
		if err != nil {
			log.Fatalf("Problem reading content of body:%v\n", err)
		}
		log.Printf("Received viewAndData:%v\n", viewAndData)
		spiteService.taskController.RunTask(viewAndData.View, viewAndData.Data)
		log.Printf("Received data:%v\n", viewAndData)
		fmt.Fprintf(response, "Success reply!")
	}
}

func addTaskHandler(spiteService *SpiteService) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var task task.TaskHTTP
		err := decoder.Decode(&task)
		if err != nil {
			log.Fatalf("Problem reading content of body:%v\n", err)
		}
		task.Task.SetTaskProcess(task.CommandStr, nil)
		log.Printf("Task info:%v\n", spiteService.taskController)
		spiteService.taskController.AddTask(task.Task)
		log.Printf("Received data:%v\n", task)
		fmt.Fprintf(response, "Success reply!")

	}
}

func getTaskDetailHandler(spiteService *SpiteService) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {

		decoder := json.NewDecoder(req.Body)
		var view task.ViewHTTP
		err := decoder.Decode(&view)
		if err != nil {
			log.Fatalf("Problem reading content of body:%v\n", err)
		}

		taskDetail := spiteService.taskController.GetTaskDetail(view.View)
		httpTask := task.TaskHTTP{
			Task:       taskDetail,
			CommandStr: taskDetail.GetTaskCommandString(),
		}
		b, err := json.Marshal(httpTask)

		if err != nil {
			log.Fatalf("Errored out with:%v\n", err)
		}
		jsonString := string(b)
		log.Printf("Json string to send:%v\n", jsonString)
		fmt.Fprintf(response, jsonString)

	}
}
func showTasksHandler(spiteService *SpiteService) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		taskViews := spiteService.taskController.GetTasksViews()
		b, err := json.Marshal(task.ViewsHTTP{
			Views: taskViews,
		})
		if err != nil {
			log.Fatalf("Errored out with:%v\n", err)
		}
		fmt.Fprintf(response, string(b))
	}
}

func helloHandler(spiteService *SpiteService) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(response, "Hello!")
		log.Printf("Hello handler!.")
	}
}

func (spiteService *SpiteService) initImplementedController() {
	//TODO Fill this out.
	mapDao := task.NewMapStoreDao()
	spiteService.taskController = task.NewController(mapDao)
}
