package main

import "log"

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	log.Println("Running main")

	//TODO Will start server here
	//Server will get list of tasks.
	//When clicking on a task, you'll get a list of inputs.
	//This will include a file upload possibly.
	//It will then bring you to another screen, showing output as it comes.
}
