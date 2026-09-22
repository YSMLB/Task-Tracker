package cli

import (
	"fmt"
	"os"

	"task-tracker/internal/service"
)

func Run() error{
	argsWithProg := os.Args

	if len(argsWithProg) < 2{
		fmt.Println("нет аргументов")
		os.Exit(1)
	}

	switch argsWithProg[1]{
	case "add": 
	if len(argsWithProg) < 3{
		fmt.Println("Пожалуйста укажите описание задачи")
		return nil
	}
	err := service.AddTask(argsWithProg[2])
	if err != nil{
		fmt.Println("error")
	}else{
		fmt.Println("Task added succesfully")
	}

	case "list": 
	err := service.ListTask()
	if err != nil{
		fmt.Println("error")
		return nil
	}
	case "update": fmt.Println("вызвана функция update")
	case "delete": fmt.Println("вызвана функция delete")
	//все команды крч
	default: fmt.Println("такой команды не существет")
	}
	return nil
}