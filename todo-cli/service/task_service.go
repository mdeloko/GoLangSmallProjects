package service

import (
	"os"
	"fmt"
	"github.com/mdeloko/GoLangSmallProjects/todo-cli/model"
	"github.com/mdeloko/GoLangSmallProjects/todo-cli/utils"
)

var taskList []model.Task
var jsonFile *os.File

func init(){
	var err error
	jsonFile, err = utils.GetJSONDatabase()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = utils.ReadTasks(jsonFile,&taskList)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func CreateTask(t model.Task)(err error){
	t.ID = len(taskList) + 1

	taskList = append(taskList, t)
	err = utils.UpdateTasks(jsonFile,taskList)
	if err != nil {
		return err
	}
	return nil
}

func ReadTasks(){
	if len(taskList)==0 {
		fmt.Println("\033[31mNão existem tarefas cadastradas!\nUtilize \033[0m'task help'\033[31m para cadastrar uma.\033[0m")
		return
	}
	for _,task:= range(taskList){
		var doneStr string
		if task.Done {
			doneStr = "Sim"
		}else{
			doneStr = "A fazer"
		}
		fmt.Println("ID:",task.ID,"\nTítulo:",task.Title,"\nFeita?",doneStr)
		fmt.Println("---------------------------------")
	}
}

func UpdateTaskTitle(id int, title string){
	if len(taskList)==0 {
		fmt.Println("\033[31mNão existem tarefas cadastradas!\nUtilize \033[0m'task help'\033[31m para cadastrar uma.\033[0m")
		return
	}
	for idx,task := range taskList{
		if task.ID == id {
			taskList[idx].Title = title
			fmt.Println("Task de ID:",task.ID,"Atualizada!\nTítulo:",taskList[idx].Title)
		}
	}
	
	utils.UpdateTasks(jsonFile,taskList)
}
func UpdateTaskState(id int, state string){
	if len(taskList)==0 {
		fmt.Println("\033[31mNão existem tarefas cadastradas!\nUtilize \033[0m'task help'\033[31m para cadastrar uma.\033[0m")
		return
	}
	for idx,task := range taskList{
		if task.ID == id {
			var doneStr string
			switch state{
			case "s","sim":
				taskList[idx].Done = true
				doneStr = "Sim"
			case "n","nao":
				taskList[idx].Done = false
				doneStr = "A fazer"
			}
			fmt.Println("Task de ID:",task.ID,"Atualizada!\nFeita?",doneStr)
		}
	}
	utils.UpdateTasks(jsonFile,taskList)
}

func DeleteTask(id int){
	if len(taskList)==0 {
		fmt.Println("\033[31mNão existem tarefas cadastradas!\nUtilize \033[0m'task help'\033[31m para cadastrar uma.\033[0m")
		return
	}
	if id > len(taskList){
		fmt.Println("\033[31mNão existe tarefa com este ID!\nUtilize \033[0m'task r'\033[31m para verificar as existentes.\033[0m")
		return
	}
	for idx, task := range taskList {
		if task.ID == id {
			taskList = append(taskList[:idx],taskList[idx+1:]...)
		}
	}
	updateIds()
	utils.UpdateTasks(jsonFile,taskList)
}

func updateIds(){
	for i := range taskList {
		taskList[i].ID = i+1
	}
	utils.UpdateTasks(jsonFile,taskList)
}