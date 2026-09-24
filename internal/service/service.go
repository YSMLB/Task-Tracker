package service

import (
	"fmt"
	"os"
	"time"

	"task-tracker/internal/storage"
)

func AddTask(description string) error {
	maxID := 0
	argWithProg := os.Args
	if len(argWithProg) < 3 {
		fmt.Println("описания нет")
		return fmt.Errorf("error")
	}

	read, err := storage.ReadTask()
	if err != nil {
		return err
	} else {
		for _, task := range read {
			if task.ID > maxID {
				maxID = task.ID
			}
		}
	}
	id := maxID + 1

	newTask := storage.Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdateAt:    time.Now(),
	}

	read = append(read, newTask)

	storage.WriteTask(read)

	return nil
}

func ListTask(targetStatus string) error {
	read, err := storage.ReadTask()
	todo := make([]storage.Task, 0)
	inProgress := make([]storage.Task, 0)
	done := make([]storage.Task, 0)

	if err != nil {
		fmt.Println("неудалось прочитать список")
		return err
	}

	if len(read) == 0 {
		fmt.Println("Список задач пуст")
	} else {
		for _, task := range read {
			if targetStatus == "" {
				fmt.Println("ID: ", task.ID,
					"\n Задача: ", task.Description,
					"\n Статус:", task.Status,
					"\n Время создания задачи: ", task.CreatedAt,
					"\n Время обновления: ", task.UpdateAt)
			} else {
				if task.Status == "todo" {
					todo = append(todo, task)

				} else if task.Status == "in-progress" {
					inProgress = append(inProgress, task)

				} else if task.Status == "done" {
					done = append(done, task)
				}
			}

		}
		if targetStatus == "todo" {
			fmt.Println("Статус todo: ")
			for _, statusTodo := range todo {
				fmt.Println(
					"ID: ", statusTodo.ID,
					"\n Задача: ", statusTodo.Description,
					"\n Время создания", statusTodo.CreatedAt,
					"\n Время обновления: ", statusTodo.UpdateAt)
			}
			//fmt.Println(todo)
		} else if targetStatus == "in-progress" {
			fmt.Println("Статус in-progress: ")
			for _, statusInProgress := range inProgress {
				fmt.Println(
					"ID: ", statusInProgress.ID,
					"\n Задача: ", statusInProgress.Description,
					"\n Время создания", statusInProgress.CreatedAt,
					"\n Время обновления: ", statusInProgress.UpdateAt)
			}
			//fmt.Println(inProgress)
		} else if targetStatus == "done" {
			fmt.Println("Статус done: ")
			for _, statusDone := range done {
				fmt.Println(
					"ID: ", statusDone.ID,
					"\n Задача: ", statusDone.Description,
					"\n Время создания", statusDone.CreatedAt,
					"\n Время обновления: ", statusDone.UpdateAt)
			}
			//fmt.Println(done)
		}
	}
	return nil
}

//
//func listStatTask() error{
//	read, err := storage.ReadTask()
//	todo := make([]storage.Task, 0)
//	inProgress := make([]storage.Task, 0)
//	done := make([]storage.Task, 0)
//
//	if err != nil{
//		fmt.Println("неудалось прочитать список")
//		return err
//	}
//	if len(read) == 0{
//		fmt.Println("список пуст")
//	}else{
//		for _,status := range read{
//			if status.Status == "todo"{
//				todo = append(todo, status)
//			}else if status.Status == "in-progress"{
//				inProgress = append(inProgress, status)
//			}else if status.Status == "done"{
//				done = append(done, status)
//			}
//		}
//		//fmt.Println("todo: \n", todo,
//		//"\nin-progress:\n", inProgress,
//		//"\ndone: \n", done)
//	}
//
//	return nil
//}

func MarkInProgressTask(id int) error {
	read, err := storage.ReadTask()
	if err != nil {
		return err
	}
	found := false
	for i := range read {
		if read[i].ID == id {
			read[i].Status = "in-progress"
			read[i].UpdateAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}
	return storage.WriteTask(read)
}


func MarkDoneTask(id int) error {
	read, err := storage.ReadTask()
	if err != nil {
		return err
	}
	found := false
	for i := range read {
		if read[i].ID == id {
			read[i].Status = "done"
			read[i].UpdateAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}
	return storage.WriteTask(read)
}



func UpdateTask(id int, newDescription string) error {
	read, err := storage.ReadTask()
	if err != nil {
		return err
	}
	found := false
	for i := range read {
		if read[i].ID == id {
			read[i].Description = newDescription
			read[i].UpdateAt = time.Now()
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}
	return storage.WriteTask(read)
}


func DeleteTask(id int) error {
	read, err := storage.ReadTask()
	if err != nil {
		return err
	}
	found := false
	for i := range read {
		if read[i].ID == id {
			// Удаление элемента из среза в Go
			read = append(read[:i], read[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}
	return storage.WriteTask(read)
}

//Команда add:
//
//Проверить, передано ли описание задачи. - +
//Прочитать текущий список задач. - +
//Сгенерировать новый ID. - +
//Зафиксировать текущее время для createdAt и updatedAt. -
//Установить статус todo. -
//Добавить задачу в список и сохранить в файл. -
//Вывести подтверждение: Task added successfully (ID: X). -
//
//Команда list (без фильтра):
//
//Прочитать задачи.
//Если задач нет — вывести информативное сообщение (например, «Список задач пуст»).
//Если задачи есть — вывести их на экран в понятном виде (список или таблица с колонками ID, Статус, Описание, Дата).
//Команда list <status>:
//
//Проверить, что переданный статус валиден (todo, in-progress или done).
//Отфильтровать задачи по этому статусу и вывести.
//
//Команды mark-in-progress и mark-done:
//
//Проверить, передан ли ID.
//Найти задачу с таким ID.
//Изменить её статус и обновить поле updatedAt текущим временем.
//Сохранить изменения.
//
//Команда update:
//
//Проверить, переданы ли оба обязательных аргумента: ID и новое описание.
//Найти задачу, обновить её description и updatedAt.
//Сохранить изменения.
//
//Команда delete:
//
//Найти задачу по ID.
//Удалить её из списка и сохранить обновлённый список.
