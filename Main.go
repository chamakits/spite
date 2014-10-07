package main

import "log"

func init() {
	log.SetFlags(log.Lshortfile)
}

func main() {
	log.Println("Running main")

	//In scope for version 0.0.1
	// > Loggability.  Must log all actions.
	// > Backup-ability.  Allows backing up of files.
	// > Provide the 'log' file of the process to the user.
	// > No Root required. Should be runnable as non-priviledged user.
	// > Don't use 'system'.  Use 'execv' family or similar. See SHELLSHOCK.

	//Out of scope for version 0.0.1
	// > Security.  Assume FOR NOW if someone is using the service, they are not malicious.
	// > Users. Assume single user for now.
	// > Permissions.  Because there3 is no concept of user in scope, there is no concept of permissions.
	// >> CAVEAT: Process WILL NOT be ran as root.  SHOULD BE RUNNABLE as non-priviledged user.

	//Nice to haves for version 0.0.1
	// > Will block root from using.  As currently there is no 'permission' concept, it would be dangerous to run as root.  Should try to block that.

	//TODO Will start service here
	//Service methods:
	//1)Service will get list of tasks.
	// ->When clicking on a task, you'll get a list of inputs.
	// ->This will include a file upload possibly.
	// ->It will then bring you to another screen, showing output as it comes.

	//2)Service will let you provide input for specific task
	// ->Input will be defined by a previously defined 'schema' for that task.
	// ->Input can be a list of named text fields and a file upload.

	//3)Service will display output of a specific task
	// ->Based on task given, it will display an 'output' of the task done.
	// ->Maybe client continue polling througout and get a diff.  Maybe do websockets.
	// ->This display SHOULD map out to a log file that exists on the server.

	//4)Service will let you create a new task
	// ->When creating a new task, you will need to specify:
	// --> Schema
	// ---> Schema contains, textfields, prepopulated textfield, calculated textfield, constants, and file uploads(provide way to backup file).
	// --> Executable to run
	// ---> As part of the executable, you can provide inputs to the executable as flags
	// ---> Flags can be based on schema, or constants.
	// ---> When providing executables, on 'save' the service will indicate if the mentioned executable exists on server.

}
