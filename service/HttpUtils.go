package service

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chamakits/spite/task"
)

/**
How the request should look:
{
"task" : {
"name"...
}
"data":{
...
}
}
**/

func GetDataFromRequest(request http.Request) task.TaskAndData {
	var taskAndData task.TaskAndData
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&taskAndData)
	if err != nil {
		log.Fatalf("Errored out with:%v\n", err)
	}
	return taskAndData
}
