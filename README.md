# Tasks-CLI

For https://roadmap.sh/projects/task-tracker

### Description: Task tracker is a project used to track and manage your tasks.
* Cobra was used for managing CLI work flow with all comands and flags.
* Simple saving in json file with simple saving logic.
* All comands and flags were tested.
### Usage :
* tasks-cli add "Buy goods" // Add task with description
* tasks-cli delete 1 // delete task with id=1
* tasks-cli update --id 1 "Smt another" // update a description and also an updatedAt time
#### tasks-cli list // All tasks :
* tasks-cli list -f todo
* tasks-cli list -f in-progress
* tasks-cli list -f done // All tasks with such status

#### To start working with it, firstly install the repo, `git clone https://github.com/Asylann/Basic-Todo-CLI`
#### Then run `go build -o tasks-cli.exe main.go` in project folder.


#### (Optional) add to Path to run cli from everywhere!
