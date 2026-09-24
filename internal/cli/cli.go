package cli

import (
	"fmt"
	"os"
	"strconv"

	"task-tracker/internal/service"
)

func Run() error {
	argsWithProg := os.Args

	if len(argsWithProg) < 2 {
		fmt.Println("нет аргументов")
		os.Exit(1)
	}

	switch argsWithProg[1] {
	case "add":
		if len(argsWithProg) < 3 {
			fmt.Println("Пожалуйста укажите описание задачи")
			return nil
		}
		err := service.AddTask(argsWithProg[2])
		if err != nil {
			fmt.Println("error")
		} else {
			fmt.Println("Task added succesfully")
		}

	case "list":
		statusStr := ""
		if len(argsWithProg) >= 3 {
			statusStr = argsWithProg[2]
		}
		err := service.ListTask(statusStr)
		if err != nil {
			fmt.Println("error")
			return nil
		}
	case "update":
		if len(argsWithProg) < 4 {
			fmt.Println("Использование: task-cli update <id> \"новое описание\"")
			return nil
		}
		id, err := strconv.Atoi(argsWithProg[2])
		if err != nil {
			fmt.Println("ID должен быть числом")
			return nil
		}
		if err := service.UpdateTask(id, argsWithProg[3]); err != nil {
			fmt.Println("Ошибка:", err)
			return nil
		}
		fmt.Println("Task updated successfully")
	case "delete":
		if len(argsWithProg) < 3 {
			fmt.Println("Использование: task-cli delete <id>")
			return nil
		}
		id, err := strconv.Atoi(argsWithProg[2])
		if err != nil {
			fmt.Println("ID должен быть числом")
			return nil
		}
		if err := service.DeleteTask(id); err != nil {
			fmt.Println("Ошибка:", err)
			return nil
		}
		fmt.Println("Task deleted successfully")
	case "mark-in-progress":
		if len(argsWithProg) < 3 {
			fmt.Println("Использование: task-cli mark-in-progress <id>")
			return nil
		}
		id, err := strconv.Atoi(argsWithProg[2])
		if err != nil {
			fmt.Println("ID должен быть числом")
			return nil
		}
		if err := service.MarkInProgressTask(id); err != nil {
			fmt.Println("Ошибка:", err)
			return nil
		}
		fmt.Println("Task marked in-progress successfully")
	case "mark-done":
		if len(argsWithProg) < 3 {
			fmt.Println("Использование: task-cli mark-done <id>")
			return nil
		}
		id, err := strconv.Atoi(argsWithProg[2])
		if err != nil {
			fmt.Println("ID должен быть числом")
			return nil
		}
		if err := service.MarkDoneTask(id); err != nil {
			fmt.Println("Ошибка:", err)
			return nil
		}
	default:
		fmt.Println("такой команды не существет")
	}
	return nil
}
