package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mdeloko/GoLangSmallProjects/todo-cli/model"
	"github.com/mdeloko/GoLangSmallProjects/todo-cli/service"
	"github.com/mdeloko/GoLangSmallProjects/todo-cli/utils"
)

func main() {
	var operation string

	if len(os.Args) > 1 {
		operation = os.Args[1]
	}

	switch operation {
	case "help", "h":
		utils.PrintHelpMessage()
	case "create", "c":
		if len(os.Args) == 4 {
			title:= os.Args[2] 
			var isDone bool
			switch os.Args[3]{
			case "s","sim":
				isDone = true
			case "n","nao":
				isDone = false
			default:
				utils.PrintWrongCommand(os.Args)
				return
			}
			err := service.CreateTask(model.Task{Title: title, Done: isDone})
			if err != nil{
				fmt.Println(err.Error())
			}else{
				fmt.Printf("Sucesso ao Inserir [%s] às tarefas!",title)
			}
		}else{

		}
	case "read","r":
		if len(os.Args) == 2 {
			service.ReadTasks()
		}else{
			utils.PrintWrongCommand(os.Args)
		}
	case "update","u":
		if len(os.Args) == 5{
			id,err := strconv.Atoi(os.Args[2])
			if err != nil{
				utils.PrintWrongCommand(os.Args)
				return
			}
			var prop string
			switch(os.Args[3]){
			case "titulo", "t":
				prop = os.Args[4]
			case "estado", "e":
				prop = os.Args[4]
				switch prop {
				case "s","sim","n","nao":
				default:
					utils.PrintWrongCommand(os.Args)
					return
				}
			default:
				utils.PrintWrongCommand(os.Args)
				return
			}
			switch prop {
			case "s","sim","n","nao":
				service.UpdateTaskState(id,prop)
			default:
				service.UpdateTaskTitle(id,prop)
			}
		}else{
			utils.PrintWrongCommand(os.Args)
			return
		}
	case "delete","d":
		if len(os.Args) == 3{
			id,err := strconv.Atoi(os.Args[2])
			if err != nil{
				utils.PrintWrongCommand(os.Args)
				return
			}
			service.DeleteTask(id)
		}else{
			utils.PrintWrongCommand(os.Args)
			return
		}
	default:
		utils.PrintHelpMessage()
		return
	}
}
