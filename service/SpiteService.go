package service

import (
	"fmt"
	"net/http"

	"github.com/chamakits/spite/task"
	"github.com/gorilla/mux"
)

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
		if "OPTIONS" == req.Method {
			corsEnable(&response)
		} else {
			handlerFunction(response, req)
		}
	}
}

func helloHandler(spiteService *SpiteService) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(response, "Hello!")
	}
}

func (spiteService *SpiteService) initImplementedController() {
	//TODO Fill this out.
}
