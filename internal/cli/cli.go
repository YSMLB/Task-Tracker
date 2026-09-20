package cli

import (
	"fmt"
	"os"
)

func Run() error{
	argsWithProg := os.Args

	if len(argsWithProg) < 2{
		fmt.Println("нет аргументов")
		os.Exit(1)
	}

	switch argsWithProg[1]{
	case "add": fmt.Println("вызвана функция add")
	case "list": fmt.Println("вызвана функция lsit")
	case "update": fmt.Println("вызвана функция update")
	case "delete": fmt.Println("вызвана функция delete")
	//все команды крч
	default: fmt.Println("такой команды не существет")
	}
	return nil
}