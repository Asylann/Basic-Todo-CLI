package intern

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"slices"
	"tasks-cli/model"
	"time"
)

func AddTask(task model.Task) (int, error) {

	task.CreatedDate = time.Now()
	task.UpdatedDate = task.CreatedDate

	lt := GetListOfTasks()

	if len(lt) == 0 {
		task.Id = 1
	} else {
		LastTaskId := lt[len(lt)-1].Id
		task.Id = LastTaskId + 1
	}

	lt = append(lt, task)

	jsonBytes, err := json.MarshalIndent(lt, "", "  ")
	if err != nil {
		log.Println("Error during Marshaling lists", err.Error())
		return task.Id, err
	}

	WriteToFile("db.json", jsonBytes)
	return task.Id, nil
}

func GetListOfTasks() []model.Task {
	ListJsonBytes, err := os.ReadFile("C:\\Users\\usena\\GolandProjects\\ToDoCLI\\db.json")
	if err != nil {
		log.Println("Error during receiving jsonBytes", err.Error())
		return nil
	}

	var ListOfTasks []model.Task
	err = json.Unmarshal(ListJsonBytes, &ListOfTasks)
	if len(ListOfTasks) == 0 {
		return ListOfTasks
	}
	if err != nil {
		log.Println("Error during decoding jsonBytes to List", err.Error())
		return nil
	}

	return ListOfTasks
}

func GetListOfTasksInProgress() []model.Task {
	ListJsonBytes, err := os.ReadFile("C:\\Users\\usena\\GolandProjects\\ToDoCLI\\db.json")
	if err != nil {
		log.Println("Error during receiving jsonBytes", err.Error())
		return nil
	}

	var ListOfTasks []model.Task
	err = json.Unmarshal(ListJsonBytes, &ListOfTasks)
	if len(ListOfTasks) == 0 {
		return ListOfTasks
	}
	if err != nil {
		log.Println("Error during decoding jsonBytes to List", err.Error())
		return nil
	}

	var result []model.Task
	for _, v := range ListOfTasks {
		if v.IsProgress {
			result = append(result, v)
		}
	}

	return result
}

func GetListOfTasksDone() []model.Task {
	ListJsonBytes, err := os.ReadFile("C:\\Users\\usena\\GolandProjects\\ToDoCLI\\db.json")
	if err != nil {
		log.Println("Error during receiving jsonBytes", err.Error())
		return nil
	}

	var ListOfTasks []model.Task
	err = json.Unmarshal(ListJsonBytes, &ListOfTasks)
	if len(ListOfTasks) == 0 {
		return ListOfTasks
	}
	if err != nil {
		log.Println("Error during decoding jsonBytes to List", err.Error())
		return nil
	}

	var result []model.Task
	for _, v := range ListOfTasks {
		if v.IsDone {
			result = append(result, v)
		}
	}

	return result
}

func GetListOfTasksTodo() []model.Task {
	ListJsonBytes, err := os.ReadFile("C:\\Users\\usena\\GolandProjects\\ToDoCLI\\db.json")
	if err != nil {
		log.Println("Error during receiving jsonBytes", err.Error())
		return nil
	}

	var ListOfTasks []model.Task
	err = json.Unmarshal(ListJsonBytes, &ListOfTasks)
	if len(ListOfTasks) == 0 {
		return ListOfTasks
	}
	if err != nil {
		log.Println("Error during decoding jsonBytes to List", err.Error())
		return nil
	}

	var result []model.Task
	for _, v := range ListOfTasks {
		if !v.IsProgress && !v.IsDone {
			result = append(result, v)
		}
	}

	return result
}

func WriteToFile(name string, jsonBytes []byte) {
	err := os.WriteFile("C:\\Users\\usena\\GolandProjects\\ToDoCLI\\db.json", jsonBytes, 0644)

	if err != nil {
		log.Println("Error during Writing list to Json", err.Error())
		return
	}
}

func UpdateTask(id int, changedDes string) (model.Task, error) {
	lt := GetListOfTasks()

	var IsChanged bool = false
	var UpdatedTask model.Task
	for i := 0; i < len(lt); i++ {
		if lt[i].Id == id {
			lt[i].Description = changedDes
			lt[i].UpdatedDate = time.Now()
			IsChanged = true
			UpdatedTask = lt[i]
		}
	}

	if !IsChanged {
		log.Println("No such id is found")
		return model.Task{}, errors.New("No such id is found")
	}

	jsonBytes, err := json.MarshalIndent(lt, "", "  ")
	if err != nil {
		log.Println("Error during Marshaling lists", err.Error())
		return model.Task{}, errors.New("Error during Marshaling lists")
	}

	WriteToFile("db.json", jsonBytes)
	return UpdatedTask, nil
}

func DeleteTask(id int) error {
	lt := GetListOfTasks()

	for i := 0; i < len(lt); i++ {
		if lt[i].Id == id {
			lt = slices.Delete(lt, i, i+1)
			break
		}
	}

	jsonBytes, err := json.MarshalIndent(lt, "", "  ")
	if err != nil {
		log.Println("Error during Marshaling lists", err.Error())
		return err
	}

	WriteToFile("db.json", jsonBytes)
	return nil
}

func MarkInProgress(id int) error {
	lt := GetListOfTasks()

	var IsChanged bool = false
	for i := 0; i < len(lt); i++ {
		if lt[i].Id == id {
			lt[i].IsProgress = true
			lt[i].UpdatedDate = time.Now()
			IsChanged = true
		}
	}

	if !IsChanged {
		log.Println("No such id is found")
		return errors.New("No such id is found")
	}

	jsonBytes, err := json.MarshalIndent(lt, "", "  ")
	if err != nil {
		log.Println("Error during Marshaling lists", err.Error())
		return err
	}

	WriteToFile("db.json", jsonBytes)
	return nil
}

func MarkDone(id int) error {
	lt := GetListOfTasks()

	var IsChanged bool = false
	for i := 0; i < len(lt); i++ {
		if lt[i].Id == id {
			lt[i].IsDone = true
			lt[i].UpdatedDate = time.Now()
			lt[i].IsProgress = false
			IsChanged = true
		}
	}

	if !IsChanged {
		log.Println("No such id is found")
		return errors.New("No such id is found")
	}

	jsonBytes, err := json.MarshalIndent(lt, "", "  ")
	if err != nil {
		log.Println("Error during Marshaling lists", err.Error())
		return err
	}

	WriteToFile("db.json", jsonBytes)
	return nil
}
