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
	taskController task.TaskController
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
	// r.HandleFunc("/api/add-task", acceptCors(addTaskHandler_STRING(spiteService)))

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

func addTaskHandler_STRING(spiteService *SpiteService) http.HandlerFunc {
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

func addTaskHandler(spiteService *SpiteService) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		decoder := json.NewDecoder(req.Body)
		var taskAndData task.TaskAndData
		err := decoder.Decode(&taskAndData)
		if err != nil {
			log.Fatalf("Problem reading content of body:%v\n", err)
		}
		// log.Printf("Task and data read:%v\n", taskAndData)
		// log.Printf("Just data:%v\n", taskAndData.Data)

		// b, err := json.Marshal(taskAndData)
		// if err != nil {
		// 	log.Fatalf("Errored out with:%v\n", err)
		// }
		// fmt.Fprintf(response, string(b))
		fmt.Fprintf(response, "Success reply!")

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
}
